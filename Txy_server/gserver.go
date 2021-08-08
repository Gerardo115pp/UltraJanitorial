package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	pecho "github.com/Gerardo115pp/PatriotLib/PatriotEcho/echo"
	patriotfs "github.com/Gerardo115pp/PatriotLib/PatriotFs"
	prouter "github.com/Gerardo115pp/PatriotLib/PatriotRouter"
)

var CONTENT_DIRECTORY = "content"

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		pecho.EchoDebug(fmt.Sprintf("File exists: %s", err.Error()))
	}
	return !os.IsNotExist(err)
}

func pageDirectoryExists(page_name string) bool {
	return fileExists(CONTENT_DIRECTORY + "/" + page_name)
}

func getContentFromFile(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		pecho.EchoErr(err)
	}
	return content
}

type Server struct {
	pages_content map[string]map[string]string
	router        *prouter.Router
	fs            *patriotfs.PatriotsFs
	port          string
	host          string
}

func (self *Server) createPageDirectory(page_name string) {
	// create a new directory for the page if it doesn't exist
	pecho.EchoDebug(fmt.Sprintf("Page %s exists: %t", CONTENT_DIRECTORY+"/"+page_name, fileExists(CONTENT_DIRECTORY+"/"+page_name)))
	if !fileExists(CONTENT_DIRECTORY + "/" + page_name) {
		pecho.Echo(pecho.OrangeFG, "Creating new page directory")
		err := os.Mkdir(CONTENT_DIRECTORY+"/"+page_name, 0755)
		if err != nil {
			pecho.EchoFatal(err)
		}
		self.pages_content[page_name] = make(map[string]string)
		err = self.fs.AddDirectory(page_name, CONTENT_DIRECTORY+"/"+page_name)
		if err != nil {
			pecho.EchoFatal(err)
		}
		self.saveContentData()
	}
}

func (self *Server) enableCors(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Headers", "X-sk")
		response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE, CONNECT")
		if request.Method == http.MethodOptions {
			response.WriteHeader(200)
			return
		}

		handler(response, request)
	}
}

func (self *Server) getPageContent(page_name string) []byte {
	if _, exists := self.pages_content[page_name]; exists {
		var content_data []byte
		content_data, err := json.Marshal(self.pages_content[page_name])
		if err != nil {
			pecho.EchoFatal(err)
		}
		return []byte(content_data)
	} else {
		pecho.EchoErr(fmt.Errorf("Page %s does not exist", page_name))
		return nil
	}
}

func (self *Server) loadContentData() {
	var content_file string = fmt.Sprintf("%s/page-content.json", CONTENT_DIRECTORY)
	if !fileExists(content_file) {
		pecho.Echo(pecho.RedFG, "No content file found, setting empty content")
		self.pages_content = make(map[string]map[string]string, 0)
		// create empty content file
		err := ioutil.WriteFile(content_file, []byte("{}"), 0644)
		if err != nil {
			pecho.EchoFatal(err)
		}
		return
	}

	pecho.Echo(pecho.CyanFG, fmt.Sprintf("Loading content from %s", content_file))
	content_structure := make(map[string]map[string]string, 0)
	content_file_data, err := ioutil.ReadFile(content_file)
	if err != nil {
		pecho.EchoFatal(err)
	}
	err = json.Unmarshal(content_file_data, &content_structure)
	if err != nil {
		pecho.EchoFatal(err)
	}
	pecho.Echo(pecho.GreenFG, fmt.Sprintf("Loaded %d page", len(content_structure)))
	for page, content_data := range content_structure {
		pecho.Echo(pecho.CyanFG, fmt.Sprintf("Loaded %d content for page '%s'", len(content_data), page))
	}
	self.pages_content = content_structure
}

func (self *Server) newImage(next http.HandlerFunc) http.HandlerFunc {
	//
	return func(response http.ResponseWriter, request *http.Request) {
		var filename, page_name, content_name string
		if request.Method == "POST" {
			content_name = request.FormValue("name")
			page_name = request.FormValue("page")
			file, header, err := request.FormFile("file")
			defer file.Close()
			if err != nil {
				pecho.EchoErr(err)
				response.WriteHeader(400)
				return
			}

			filename = header.Filename //content_name + filepath.Ext(header.Filename)
			self.writeToPageContent(page_name, fmt.Sprintf("%s/%s/%s", self.fs.GetPrefix(), page_name, filename), content_name)
		}
		next(response, request)
		if filename != "" {
			if !fileExists(fmt.Sprintf(CONTENT_DIRECTORY+"/%s/%s", page_name, filename)) {
				pecho.Echo(pecho.OrangeFG, fmt.Sprintf("Image was not created, removing from content page"))
				self.removeContentFromContentPage(page_name, content_name)
			} else {
				pecho.Echo(pecho.GreenFG, fmt.Sprintf("Image was created"))
				fmt.Fprintf(response, "%s/%s/%s", self.fs.GetPrefix(), page_name, filename)
			}

		}
	}
}

func (self *Server) parseFormToMap(request *http.Request) map[string]string {
	form := make(map[string]string)
	request.ParseForm()
	request.ParseMultipartForm(32 << 20)
	pecho.EchoDebug(fmt.Sprintf("Request form: %+v", request.Form))
	for key, value := range request.Form {
		form[key] = value[0]
	}
	return form
}

//greets the user to let him know that the server is up and running
func (self *Server) handleIndex(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	fmt.Fprintf(response, "Welcome to the Txy Server, service is up and running with %d pages", len(self.pages_content))
}

// /pages handler
func (self *Server) handlePages(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var pages_content []byte
		pages_content, err := json.Marshal(self.pages_content)
		if err != nil {
			pecho.EchoFatal(err)
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(200)
		response.Write(pages_content)
	} else if request.Method == "POST" {
		var page_name string = request.FormValue("page")
		self.createPageDirectory(page_name)
		response.WriteHeader(200)
	} else {
		response.WriteHeader(405) // Method not allowed
	}

}

// /content handler
func (self *Server) newContent(response http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		self.handleContentRetrival(response, request)
	case "POST":
		pecho.Echo(pecho.CyanFG, "Received new content")
		self.handleNewContent(response, request)
	default:
		pecho.Echo(pecho.RedFG, "Unsupported method")
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// is callable by the newContent function
func (self *Server) handleContentRetrival(response http.ResponseWriter, request *http.Request) {
	page_name := request.URL.Query().Get("page")
	if page_name != "" {
		self.createPageDirectory(page_name)
		var page_content_json []byte = self.getPageContent(page_name)
		if len(page_content_json) > 0 {
			pecho.Echo(pecho.GreenFG, "Success")
			response.WriteHeader(http.StatusOK)
		} else {
			pecho.Echo(pecho.RedFG, "Page does not exist")
			response.WriteHeader(http.StatusAccepted)
		}
		response.Write(page_content_json)
	} else {
		pecho.Echo(pecho.RedFG, "Invalid request")
		response.WriteHeader(http.StatusBadRequest)
	}
}

// is called by the newContent function
func (self *Server) handleNewContent(response http.ResponseWriter, request *http.Request) {
	var request_form map[string]string = self.parseFormToMap(request)
	pecho.EchoDebug(fmt.Sprintf("Request form: %+v", request_form))

	content_name := request_form["name"]
	page_name := request_form["page"]
	content := request_form["content"]

	if content_name != "" && page_name != "" {
		self.createPageDirectory(page_name)

		//cleaing up the content
		content = strings.TrimSpace(content)
		content = strings.Replace(content, "\n", "", -1)

		self.writeToPageContent(page_name, content, content_name)
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusBadRequest)
		pecho.Echo(pecho.RedFG, "Invalid request")
	}
}

func (self *Server) registerFileSystems() {
	for page := range self.pages_content {
		pecho.Echo(pecho.CyanFG, fmt.Sprintf("Registering page '%s'", page))
		self.fs.AddDirectory(page, fmt.Sprintf("%s/%s", CONTENT_DIRECTORY, page))
	}
}

func (self *Server) saveContentData() error {
	/**
	saves the map 'self.page_content' to the file 'page-content.json'
	*/
	var page_content_file string = fmt.Sprintf("%s/page-content.json", CONTENT_DIRECTORY)
	if !fileExists(page_content_file) {
		pecho.Echo(pecho.RedFG, "No page content file found")
	}
	pecho.Echo(pecho.CyanFG, "Saving page content")
	page_content_json, err := json.MarshalIndent(self.pages_content, "", "  ")
	if err != nil {
		pecho.EchoErr(err)
		return err
	}
	err = ioutil.WriteFile(page_content_file, page_content_json, 0640)
	if err != nil {
		pecho.EchoErr(err)
		return err
	}
	pecho.Echo(pecho.GreenFG, "Saved page content")
	return nil
}

func (self *Server) writeToPageContent(page string, content string, content_name string) {
	if !pageDirectoryExists(page) {
		pecho.EchoWarn(fmt.Sprintf("Page directory does not exist: %s", page))
	}
	pecho.EchoDebug(fmt.Sprintf("Writing to page content: %s/%s", page, content_name))
	if _, exists := self.pages_content[page]; !exists {
		self.pages_content[page] = make(map[string]string)
	}
	self.pages_content[page][content_name] = content
	self.saveContentData()
}

func (self *Server) removeContentFromContentPage(page string, content_name string) {
	pecho.EchoDebug(fmt.Sprintf("Removing content from page: %s/%s", page, content_name))
	if _, exists := self.pages_content[page]; !exists {
		pecho.EchoWarn(fmt.Sprintf("Page directory does not exist: %s", page))
		return
	}
	delete(self.pages_content[page], content_name)
	self.saveContentData()
	return
}

func (self *Server) run() {
	self.loadContentData()

	// initial configurations
	if !fileExists(CONTENT_DIRECTORY) {
		pecho.Echo(pecho.RedFG, "Content directory does not exist")
		err := os.Mkdir(CONTENT_DIRECTORY, 0755)
		if err != nil {
			pecho.EchoFatal(err)
		}
	}
	self.fs.AddMiddleware(self.newImage)
	self.router.SetCorsHandler(self.enableCors)

	// setting filesystem up
	var filesystem_prefix string = "/static"
	self.registerFileSystems()
	self.fs.SetPrefix(filesystem_prefix)
	self.createPageDirectory("general") // this page is used for general purpose content like the logo, footer, etc.

	//setting up prefix handlers
	self.router.RedirectIfPrefix(self.fs.GetPrefix(), self.fs)

	// setting up routes
	self.router.RegisterRoute(prouter.NewRoute("/", true), self.handleIndex)
	self.router.RegisterRoute(prouter.NewRoute("/content", true), self.newContent)
	self.router.RegisterRoute(prouter.NewRoute("/pages", true), self.handlePages)

	pecho.Echo(pecho.CyanFG, fmt.Sprintf("Starting server on: %s:%s", self.host, self.port))
	if err := http.ListenAndServe(self.host+":"+self.port, self.router); err != nil {
		pecho.EchoFatal(err)
	}
}

func setEnviromentIfFirstRun() {
	if !fileExists(CONTENT_DIRECTORY) {
		pecho.Echo(pecho.GreenFG, "Content directory does not exist")
		pecho.Echo(pecho.GreenFG, "Creating content directory")
		err := os.Mkdir(CONTENT_DIRECTORY, 0755)
		if err != nil {
			pecho.EchoFatal(err)
		}
		pecho.Echo(pecho.GreenFG, "Content directory created")
	}
	// log the current work directory
	current_directory, err := os.Getwd()
	if err != nil {
		pecho.EchoFatal(err)
	}
	pecho.Echo(pecho.GreenFG, fmt.Sprintf("Serving content from: %s", current_directory))
}

func CreateServer() *Server {
	/*
		create a new server taking the port and host from environment variables and setting
		them to localhost:5000 in case they are not set
	*/
	s := new(Server)
	s.router = prouter.CreateRouter()
	s.fs = patriotfs.CreateFs(false, int64(patriotfs.MB*20))
	s.port = "5000"
	s.host = "127.0.0.1"
	if port := os.Getenv("PORT"); port != "" {
		s.port = port
	}
	if host := os.Getenv("HOST"); host != "" {
		s.host = host
	}
	return s
}

func main() {
	if os.Getenv("CONTENT_DIRECTORY") != "" {
		CONTENT_DIRECTORY = os.Getenv("CONTENT_DIRECTORY")
	}

	setEnviromentIfFirstRun()

	fmt.Println("Starting server...")
	var server *Server = CreateServer()
	server.run()
}

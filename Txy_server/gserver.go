package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	pecho "github.com/Gerardo115pp/PatriotLib/PatriotEcho/echo"
	prouter "github.com/Gerardo115pp/PatriotLib/PatriotRouter"
)

var CONTENT_DIRECTORY = "./content"

func createPageDirectory(page_name string) {
	// create a new directory for the page if it doesn't exist
	pecho.EchoDebug(fmt.Sprintf("Page %s exists: %t", CONTENT_DIRECTORY+"/"+page_name, fileExists(CONTENT_DIRECTORY+"/"+page_name)))
	if !fileExists(CONTENT_DIRECTORY + "/" + page_name) {
		pecho.Echo(pecho.OrangeFG, "Creating new page directory")
		err := os.Mkdir(CONTENT_DIRECTORY+"/"+page_name, 0755)
		if err != nil {
			pecho.EchoFatal(err)
		}
	}
}

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

func getPageContent(page_name string) []byte {
	if fileExists(CONTENT_DIRECTORY + "/" + page_name) {
		//list all the files in the page directory within the content directory
		files, _ := ioutil.ReadDir(CONTENT_DIRECTORY + "/" + page_name)
		var page_content map[string]string = make(map[string]string)
		for _, file := range files {
			if file.IsDir() {
				continue
			} else if strings.HasPrefix(file.Name(), "image") {
				pecho.Echo(pecho.RedFG, "No support impemented for:"+file.Name())
				continue
			}
			page_content[file.Name()] = string(getContentFromFile(CONTENT_DIRECTORY + "/" + page_name + "/" + file.Name()))
		}
		var page_content_json []byte
		page_content_json, err := json.Marshal(page_content)
		if err != nil {
			pecho.EchoFatal(err)
		}
		return page_content_json
	} else {
		pecho.Echo(pecho.RedFG, "Page does not exist")
	}
	return []byte("")
}

func writeToPageContent(page string, filename string, content_name string) {
	if !pageDirectoryExists(page) {
		createPageDirectory(page)
	}

	content_file := fmt.Sprintf("%s/%s/page-content.json", CONTENT_DIRECTORY, page)
	pecho.Echo(pecho.CyanFG, fmt.Sprintf("Writing content to %s", content_file))
	content_data := make(map[string]string, 0)
	if fileExists(content_file) {
		content_file_data, err := ioutil.ReadFile(content_file)
		if err != nil {
			pecho.EchoFatal(err)
		}
		json.Unmarshal(content_file_data, &content_data)
	}
	content_data[content_name] = filename
	content_file_data, err := json.Marshal(content_data)
	if err != nil {
		pecho.EchoFatal(err)
	}
	err = ioutil.WriteFile(content_file, content_file_data, 0640)
	if err != nil {
		pecho.EchoFatal(err)
	}
	pecho.Echo(pecho.GreenFG, fmt.Sprintf("Wrote content to %s", content_file))
}

type Server struct {
	router *prouter.Router
	port   string
	host   string
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

func (self *Server) newImage(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		content_name := request.FormValue("name")
		page_name := request.FormValue("page")
		file, header, err := request.FormFile("file")
		defer file.Close()
		if err != nil {
			pecho.EchoErr(err)
			response.WriteHeader(400)
			return
		}
		pecho.Echo(pecho.CyanFG, fmt.Sprintf("Received new image: %s", header.Filename))
		file_data, err := ioutil.ReadAll(file)
		if err != nil {
			pecho.EchoErr(err)
			response.WriteHeader(400)
			return
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s", CONTENT_DIRECTORY, page_name, header.Filename), file_data, 0640)
		if err != nil {
			pecho.EchoErr(err)
			response.WriteHeader(500)
			return
		}

		pecho.Echo(pecho.CyanFG, fmt.Sprintf("Wrote new image: %s", header.Filename))
		writeToPageContent(page_name, header.Filename, content_name)
		response.WriteHeader(200)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

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

func (self *Server) handleContentRetrival(response http.ResponseWriter, request *http.Request) {
	page_name := request.URL.Query().Get("page")
	if page_name != "" {
		createPageDirectory(page_name)
		var page_content_json []byte = getPageContent(page_name)
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

func (self *Server) handleNewContent(response http.ResponseWriter, request *http.Request) {
	var request_form map[string]string = self.parseFormToMap(request)
	pecho.EchoDebug(fmt.Sprintf("Request form: %+v", request_form))

	content_name := request_form["name"]
	page_name := request_form["page"]
	content := request_form["content"]

	if content_name != "" && page_name != "" {
		createPageDirectory(page_name)

		//cleaing up the content
		content = strings.TrimSpace(content)
		content = strings.Replace(content, "\n", "", -1)

		content_file := CONTENT_DIRECTORY + "/" + page_name + "/" + content_name

		err := ioutil.WriteFile(content_file, []byte(content), 0644)
		if err != nil {
			pecho.EchoErr(err)
			response.WriteHeader(http.StatusInternalServerError)
		}

		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusBadRequest)
		pecho.Echo(pecho.RedFG, "Invalid request")
	}
}

func (self *Server) run() {

	// initial configurations
	if !fileExists(CONTENT_DIRECTORY) {
		pecho.Echo(pecho.RedFG, "Content directory does not exist")
		err := os.Mkdir(CONTENT_DIRECTORY, 0755)
		if err != nil {
			pecho.EchoFatal(err)
		}
	}

	self.router.RegisterRoute(prouter.NewRoute("/content", true), self.enableCors(self.newContent))
	self.router.RegisterRoute(prouter.NewRoute("/images", true), self.enableCors(self.newImage))

	pecho.Echo(pecho.CyanFG, "Starting server on port "+self.port)
	if err := http.ListenAndServe(self.host+":"+self.port, self.router); err != nil {
		pecho.EchoFatal(err)
	}
}

func CreateServer() *Server {
	/*
		create a new server taking the port and host from environment variables and setting
		them to localhost:5000 in case they are not set
	*/
	s := new(Server)
	s.router = prouter.CreateRouter()
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

	fmt.Println("Starting server...")
	var server *Server = CreateServer()
	server.run()
}

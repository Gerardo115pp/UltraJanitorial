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
	if !fileExists(CONTENT_DIRECTORY + "/" + page_name) {
		err := os.Mkdir(CONTENT_DIRECTORY+"/"+page_name, 0620)
		if err != nil {
			pecho.EchoFatal(err)
		}
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
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

func (self *Server) handleNewContent(response http.ResponseWriter, request *http.Request) {
	var request_form map[string]string = self.parseFormToMap(request)

	content_name := request_form["name"]
	page_name := request_form["page"]
	content := request_form["content"]

	if content_name != "" && page_name != "" {
		createPageDirectory(page_name)
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

func (self *Server) parseFormToMap(request *http.Request) map[string]string {
	form := make(map[string]string)
	request.ParseForm()
	pecho.EchoDebug(fmt.Sprintf("Request form: %+v", request.Form))
	for key, value := range request.Form {
		form[key] = value[0]
	}
	return form
}

func (self *Server) newContent(response http.ResponseWriter, request *http.Request) {
	pecho.Echo(pecho.CyanFG, "Received new content")

	switch request.Method {
	case "GET":
		pecho.Echo(pecho.WhiteFG, "GET request")
		self.handleContentRetrival(response, request)
	case "POST":
		pecho.Echo(pecho.WhiteFG, "POST request")
		self.handleNewContent(response, request)
	default:
		pecho.Echo(pecho.RedFG, "Unsupported method")
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (self *Server) handleContentRetrival(response http.ResponseWriter, request *http.Request) {
	content_name := request.URL.Query().Get("content")
	page_name := request.URL.Query().Get("page")
	if content_name == "*" && page_name != "" {
		createPageDirectory(page_name)
		var page_content_json []byte = getPageContent(page_name)
		if len(page_content_json) > 0 {
			response.WriteHeader(http.StatusOK)
		} else {
			response.WriteHeader(http.StatusNotFound)
		}
		response.Write(page_content_json)
	} else {
		pecho.Echo(pecho.RedFG, "Invalid request")
		response.WriteHeader(http.StatusBadRequest)
	}
}

func (self *Server) run() {

	self.router.RegisterRoute(prouter.NewRoute("/register-new-content", true), self.newContent)

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

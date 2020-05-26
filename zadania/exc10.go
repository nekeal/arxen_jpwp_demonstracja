package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type PersonalInformation struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Server struct {
	usersList []*PersonalInformation
}

const PORT_NUM = ":8000"

func Router() (*mux.Router, *Server) {
	// create new instance of Server
	s := &Server{
		usersList: []*PersonalInformation{},
	}

	// create new router
	r := mux.NewRouter()
	// router paths
	r.HandleFunc("/", s.HomeHandler)
	r.HandleFunc("/api/{key}", s.ApiHandler)
	r.HandleFunc("/path/{name}/{surname}", s.NameHandler)

	// return router and server
	return r, s
}

func serve() {
	// create new router and server
	r, _ := Router()

	// serve http on particular port
	if err := http.ListenAndServe(PORT_NUM, r); err != nil {
		panic(err)
	}
}

func (s *Server) NameHandler(writer http.ResponseWriter, request *http.Request) {
	//TODO dodaj do listy usersList imie i nazwisko usera, nastepnie
	panic("implement me")
}

func (s *Server) ApiHandler(writer http.ResponseWriter, request *http.Request) {
	// get all arguments from path
	vars := mux.Vars(request)
	// write status ok to header (good practise)
	writer.WriteHeader(http.StatusOK)
	// try to write response
	if _, err := writer.Write([]byte("my key is " + vars["key"])); err != nil {
		panic("something went horribly wrong")
	}
}

func (s *Server) HomeHandler(writer http.ResponseWriter, request *http.Request) {
	// try to write response
	if _, err := writer.Write([]byte("Some basic response")); err != nil {
		panic("something went horribly wrong")
	}
}

func main() {
	serve()
}

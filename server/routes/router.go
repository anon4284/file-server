package router

import (
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//Router the main router of the application
type Router struct {
	port   int
	server *negroni.Negroni
	router *mux.Router
}

//New create new Router
func New(port int) *Router {
	return &Router{server: negroni.Classic(), router: mux.NewRouter(), port: port}
}

//ServeHTMLifNotFound serves html file if no other route was specified by the frontend
func (r *Router) ServeHTMLifNotFound(path string) {
	r.router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	})
}

//Start assignes handler to Server and starts the server
func (r *Router) Start() {
	r.server.UseHandler(r.router)
	r.server.Run(":" + strconv.Itoa(r.port))
}

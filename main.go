package main

import (
	"net/http"

	"github.com/dragosslujeru/go-rest-api/controllers"
)

func main() {
	webserver := http.NewServeMux()
	webserver.Handle("/", controllers.Home)
	webserver.Handle("/home", controllers.Home)
	webserver.Handle("/user",
		controllers.HandleAll(
			controllers.GetHandler(controllers.GetUser),
			controllers.PostHandler(controllers.PostUser)))
	webserver.Handle("/users", controllers.GetAllUsers)
	http.ListenAndServe(":8080", webserver)
}

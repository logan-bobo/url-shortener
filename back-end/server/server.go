package server

import (
	"github.com/logan-bobo/url_shortener/controllers"
)

func StartServer(server *controllers.Server) {
	r := NewRouter(server)

	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}

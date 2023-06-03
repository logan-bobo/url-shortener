package server

import (
	"github.com/logan-bobo/url_shortener/controllers"
)

func Init() {
	controllers.Init()

	r := NewRouter()

	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}

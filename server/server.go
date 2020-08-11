package server

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	//"sync"
	"../functions"
)

func RunServer() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/product/:code", getProduct),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func getProduct(w rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")
	product := functions.SearchProduct(code)

	if product == nil {
		rest.NotFound(w, r)
		return
	}

	_ = w.WriteJson(product)
}

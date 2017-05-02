package templates

const controllerFile = `package controllers

import (
  "net/http"
  "fmt"
)

func MainControllerGet(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello world")
}
`

const modelFile = `package models

type MainModel struct{}

func(m * MainModel) create(){}
func(m * MainModel) retrieve(){}
func(m * MainModel) update(){}
func(m * MainModel) delete{}
`

// MuxURL ...
const MuxURL = "go get -u github.com/gorilla/mux"

const routerFile = `package routers

import (
    "{{.AppName}}/controllers"
    "github.com/gorilla/mux"
)

func GetRouter() *mux.Router{
  r := mux.NewRouter()
  r.HandleFunc("/", controllers.MainControllerGet)

  return r
}
`

const mainFile = `package main

import (
  "net/http"
  "log"
  "{{.AppName}}/routers"
)

func main(){
  r := routers.GetRouter()
  // http.Handle("/", routers.GetRouter())

  log.Fatal(http.ListenAndServe(":{{.Port}}", r))
}
`

// GetTemplateContent ...
func GetTemplateContent() map[string]string {
	templateContent := make(map[string]string)

	templateContent["controllers"] = controllerFile
	templateContent["models"] = modelFile
	templateContent["routers"] = routerFile
	templateContent["main"] = mainFile

	return templateContent
}

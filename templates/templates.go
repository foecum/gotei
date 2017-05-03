package templates

// LiveReloadJS ...
//I will add the debug option in the config to remove this for production or if
//the user doesn't want to use it
const LiveReloadJS = `<script>function tryConnectToReload(address) {
  var conn;
  // This is a statically defined port on which the app is hosting the reload service.
  conn = new WebSocket("ws://localhost:12450/reload");

  conn.onclose = function(evt) {
    // The reload endpoint hasn't been started yet, we are retrying in 2 seconds.
    setTimeout(() => tryConnectToReload(), 2000);
  };

  conn.onmessage = function(evt) {
    console.log("Refresh received!");

    // If we uncomment this line, then the page will refresh every time a message is received.
    //location.reload()
  };
}

try {
  if (window["WebSocket"]) {
    tryConnectToReload();
  } else {
    console.log("Your browser does not support WebSocket, cannot connect to the reload service.");
  }
} catch (ex) {
  console.log('Exception during connecting to reload:', ex);
}</script>`

const controllerFile = `package controllers

import (
	"bytes"
	"fmt"
	"net/http"
  "html/template"

	"github.com/foecum/gotei2.0/templates"
)
// MainControllerGet ...
func MainControllerGet(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello world")
	writeResponse(w, "Hello world")
}

func writeResponse(w http.ResponseWriter, content string) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteString(templates.LiveReloadJS)

  t, err := template.ParseFiles("edit.html")

  if err != nil{
    e.
  }
	fmt.Fprintf(w, buffer.String())
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

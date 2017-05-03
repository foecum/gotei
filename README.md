# Gotei2.0
A live reload tool for development web apps written in go.
## installation
Install dependencies for the boilerplate app (I need to automate this part)

    github.com/gorilla/mux
    github.com/gorilla/websocket

#### Install Gotei
    go get github.com/foecum/gotei2.0

## Usage
#### Restart app automatically when you make changes
    cd <project folder>
    gotei run

#### To build the app to the go path bin folder
    cd <project folder>
    gotei install

#### To build the app in to the source folder
    cd <project folder>
    gotei install

#### To create a simple boilerplate web app
The Structure includes a controller, a router and a simple model

    cd <project folder>
    gotei new <appname>

### TODO
4. Add options to specify enviroment variables eg IP, PORT...
5. Add flags like test to run tests before every build e.t.c...

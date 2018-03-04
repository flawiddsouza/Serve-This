package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/goji/httpauth"
    "github.com/BurntSushi/toml"
)

// DEFAULT CONFIG
var port int = 15852
var enableAuth bool = false
var username string = "user"
var password string = "pass"

type Configuration struct {
    Port int
    EnableAuth bool
    Username string
    Password string
}

func main() {
    var conf Configuration
    toml.DecodeFile("config.toml", &conf)
    if conf.Port != 0 { port = conf.Port }
    if enableAuth != false { enableAuth = conf.EnableAuth }
    if username != "" { username = conf.Username }
    if password != "" { password = conf.Password }

    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

        fmt.Fprintf(w, "Append your file path to the url to access that file as a served file")

    })

    r.HandleFunc("/{path:(.*?)}", func(w http.ResponseWriter, req *http.Request) {

        // allow cross domain AJAX requests
        w.Header().Set("Access-Control-Allow-Origin", "*")

        path := mux.Vars(req)["path"]

        if _, err := os.Stat(path); err == nil {
            http.ServeFile(w, req, path)
        } else {
            fmt.Fprintf(w, "Requested file not found!")
        }

    })

    if(enableAuth) {
        http.Handle("/", httpauth.SimpleBasicAuth(username, password)(r))
    } else {
        http.Handle("/", r)
    }

    panic(http.ListenAndServe(":" + strconv.Itoa(port), nil))

}

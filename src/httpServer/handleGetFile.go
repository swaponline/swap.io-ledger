package httpserver

import "net/http"

func InitialiseHandleFile() {
    http.Handle("/", http.FileServer(http.Dir("static")))
}

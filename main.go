package main

import (
	"fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting Project...")
    // ListenAndSever fucntion starts the server and binds it to 8080. WHenever it receives an http request it will hand the request off the http.Handler function. In our case its http.FileServer
    http.ListenAndServe(":8080",http.FileServer(http.Dir(".")))
}

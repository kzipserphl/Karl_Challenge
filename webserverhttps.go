package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "runtime"
    "log"
)

var helloTemplate = `<html>
<head>
<title>Hello World</title>
</head>
<body>
<h1>Hello World!</h1>
</body>
</html>`

func main() {
    mux := http.NewServeMux()
    mux.Handle("/index", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(helloTemplate))
    }))
    exit := make(chan struct{}, 1)
    // start the server
    go func() {
        fmt.Println("Listening on https://localhost")
        // http.ListenAndServe(":443", mux)
        err := http.ListenAndServeTLS(":443", "server.crt", "server.key", mux)
        if err != nil {
            log.Fatal("ListenAndServe: ", err)
        }

        exit <- struct{}{}
    }()

    
    var err error
    switch runtime.GOOS {
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://localhost/index").Start()
    case "darwin":

        err = exec.Command("open", "https://localhost/index").Start()
    case "linux":
        err = exec.Command("xdg-open", "https://localhost/index").Start()
    }
    if err != nil {
        fmt.Println(err)
    }
    <-exit
}
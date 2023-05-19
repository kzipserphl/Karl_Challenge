package main

import (
  "fmt"
  "log"
  "net/http"
  "os/exec"
  "runtime"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "<html><head><title>Hello World</title></head><body><h1>Hello World!</h1></body></html>")
    
}

func openbrowser(url string) {
    var err error
    
    switch runtime.GOOS {
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
    case "darwin":
        err = exec.Command("open", url).Start()
    default:
        err = fmt.Errorf("unsupported platform")
    }
    if err != nil {
        log.Fatal(err)
    }
    //
}

func main() {

    http.HandleFunc("/hello", helloHandler)
    
    var url = "http://localhost:8080/hello"
    fmt.Println("Listening at: http://localhost:8080")
    fmt.Println("Opening browser to ", url)
    
    openbrowser(url)
    
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
    
}

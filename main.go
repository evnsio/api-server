package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "os"
)

var SERVICE_HEALTHY bool = true
var HOST string = ""

func main() {

    HOST = os.Getenv("HOSTNAME")

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Logger(Index))
    router.HandleFunc("/health", Logger(Health))
    router.HandleFunc("/status", Logger(Health))
    router.HandleFunc("/toggle-health", Logger(ToggleHealth))
    router.HandleFunc("/crash", Logger(Crash))

    log.Fatal(http.ListenAndServe(":8080", router))
}


func Index(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "Service Running (%v)\n", HOST)
}


func Health(w http.ResponseWriter, r *http.Request) {

    if SERVICE_HEALTHY {
        fmt.Fprintf(w, "Service Healthy (%v)\n", HOST)
    } else {
        message := "Service Unhealthy (" + HOST + ")"
        http.Error(w, message, http.StatusInternalServerError)
    }
}


func ToggleHealth(w http.ResponseWriter, r *http.Request) {
    SERVICE_HEALTHY = !SERVICE_HEALTHY
    fmt.Fprintf(w, "Service Health set to %v (%v)\n", SERVICE_HEALTHY, HOST)
}


func Crash(w http.ResponseWriter, r *http.Request) {
    os.Exit(1)
}
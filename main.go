package main

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
)

type CitiesResponse struct {
    Cities []string `json:"cities"`
}

func CityHandler(res http.ResponseWriter, req *http.Request) {
    citiesResponse := &CitiesResponse{
        Cities: []string{
            "NewYorkCity ~testPushOcir",
            "LA",
            "Chicago",
            "Philly",
            "Mt Laurel",
        },
    }
    data, _ := json.MarshalIndent(citiesResponse, "", "  ")
    res.Header().Set("Content-Type", "application/json; charset=utf-8")
    res.Write(data)
}

func logHandler(w http.ResponseWriter, res *http.Request) {
    res.Write(w)
    fmt.Fprintf(w, "RemoteAddr: %s", res.RemoteAddr)
    log.Println("Request Body: %s", res.Body)
}


func main() {
    log.Println("Listening on this host: http://localhost:5005")

    http.HandleFunc("/", logHandler)
    http.HandleFunc("/cities.json", CityHandler)
    err := http.ListenAndServe(":5005", nil)
    if err != nil {
        log.Fatal("Unable to listen on :5005: ", err)
    }
}

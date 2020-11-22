package main

import (
	"encoding/json"
	"fmt"
	"log"
    "net/http"
    "io/ioutil"
)

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var db map[string]string = make(map[string]string)

var port int = 9090

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func retriveAllKeys() []string {

    keys := make([]string, 0, len(db))

    for k := range db {
        keys = append(keys, k)
    }

    return keys
}

func valueExist(key string) bool {

    var exist bool = false

    if db[key] != "" {
        exist = true
    }

    return exist
}

func getValue(key string) string {

    var value string = ""

    if valueExist(key) {
        value = db[key]
    }

    return value
}

func setValue(key string, value string) bool {

    var inserted bool = false

    db[key] = value

    if valueExist(key) {
        inserted = true
    }

    return inserted
}


func readOrWriteValue(w http.ResponseWriter, r *http.Request) {

    enableCors(&w)

    var key string = r.URL.Path[len("/articles/"):]

    if r.Method == "GET" {

        if key != "" {
            if valueExist(key) {
                fmt.Fprintf(w, getValue(key))
            } else {
                fmt.Fprint(w, "")
            }
        } else {
            json.NewEncoder(w).Encode(retriveAllKeys())
        }

    } else if r.Method == "PUT" {

        body, err := ioutil.ReadAll(r.Body)

        if err != nil {
            panic(err)
        }

        exists := valueExist(key)

        setValue(key, string(body))

        if exists {
            fmt.Fprintf(w, "HTTP/1.1 201 " + http.StatusText(201))
        } else {
            fmt.Fprintf(w, "HTTP/1.1 200 " + http.StatusText(200))
        }
    }
}

func handleRequests() {
    http.HandleFunc("/articles/", readOrWriteValue)
    log.Fatal(http.ListenAndServe(":" + fmt.Sprint(port), nil))
}

func main() {
    fmt.Println("Listening on Port " + fmt.Sprint(port))
    setValue("rest_api", "REpresentational State Transfer Application Programming Interface")
    handleRequests()
}
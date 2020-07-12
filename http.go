package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I am %s\n", hostname)

        // Cut first 6 chars from hostname to use it as hex color in HTML
        // use hostname
        stringToTrim := hostname

        // Step 1: Convert it to a rune
        a := []rune(stringToTrim)

        // Step 2: Grab the num of chars you need
        colorCode := string(a[0:6])

        // Print out the hostname with background-color generated from it
        fmt.Fprintf(w, "<body style='background-color:#%s\n", colorCode)
        fmt.Fprintf(w, ";'>")
        fmt.Fprintf(w, "<p style='color: #000000; background-color: #ffffff; font-weight: bold'>I am: %s\n", hostname)
        fmt.Fprintf(w, "</p></body>")
    })
    log.Fatal(http.ListenAndServe(":" + port, nil))
}

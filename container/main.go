package main

import (
	"log"
	"math"
	"net/http"
	"fmt"
	"time"
	"os"
	"math/rand"
	"strconv"
)

// from https://stackoverflow.com/questions/26152993/go-logger-to-print-timestamp
func Log(l *log.Logger, msg string) {
    l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " [INFO] ")
    l.Print(msg)
}

func handler(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    id := rand.Intn(100000)
    l := log.New(os.Stdout, "", 0)
    hostname, err := os.Hostname()
    if err != nil {
		fmt.Println(err)
		hostname = "pedrito"
	}
    Log(l, "Someone called the method. id: " + strconv.Itoa(id))
	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}
    Log(l, "Finish calling. ---------- id: " + strconv.Itoa(id))
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"Server\": \"%s\", \"message\": \"OK\"}", hostname)
}

func main() {
    l := log.New(os.Stdout, "", 0)
    Log(l, "Starting...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

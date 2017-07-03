package main

import (
	"fmt"
	"log"
	"net/http"

	"routes"
)

func main() {
	http.HandleFunc("/clouldprt/prtorder", routes.PrtOrder)
	addr := ":8080"
	fmt.Println("run at " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

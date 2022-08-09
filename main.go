package main

import (
	"fmt"
	"net/http"
)

//port Number
const portNumber = ":8989"

func main() {

	http.HandleFunc("/", Home)
	//http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}

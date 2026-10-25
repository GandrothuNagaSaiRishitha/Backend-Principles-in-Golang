// Routing and middleware
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/home",homeHandler);
	http.ListenAndServe(":8080",)
}
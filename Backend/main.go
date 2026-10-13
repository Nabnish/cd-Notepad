package main

import(
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"heelow")
}

func main(){
	http.HandleFunc("/",Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


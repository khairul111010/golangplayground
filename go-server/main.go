package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "Parseform() err: %v",err)
		return
	}
	fmt.Fprintf(w, "Post request successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}


func hellohandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!!!")
}


func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formhandler)
	http.HandleFunc("/hello",hellohandler)
	fmt.Print("server started 8080\n")
	if err := http.ListenAndServe(":8080",nil); err !=nil{
		log.Fatal(err)
}
}
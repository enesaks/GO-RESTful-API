package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", AnaSayfa)
	http.HandleFunc("/detay", Detay)
	http.ListenAndServe(":8080", nil)
}

func AnaSayfa(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ana Sayfa Çalıştı")
	view, _ := template.ParseFiles("View/index.html")
	view.Execute(w, nil)
}
func Detay(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Detay Sayfa Çalıştı")
	view, _ := template.ParseFiles("View/detay.html")
	view.Execute(w, nil)
}

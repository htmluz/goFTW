package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "metodo n suportado", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "oieee")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse form err", err)
		return
	}
	fmt.Fprintf(w, "POST bem sucedido")
	name := r.FormValue("nome")
	address := r.FormValue("endereco")
	fmt.Fprintf(w, "nome = %s\n", name)
	fmt.Fprintf(w, "endereco = %s\n", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server pronto na porta 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"Api-Aula1/router"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	r := router.New()
	const addr = ":8080"
	log.Printf("Servidor ouvindo em %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func buscarLivros(w http.ResponseWriter, r *http.Request) {
	nomeLivro := r.URL.Query().Get("nome")

	if nomeLivro == "" {
		http.Error(w, "Por favor, forneça o nome do livro", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("CHAVEBOOKS")
	urlPesquisa := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&key=%s", url.QueryEscape(nomeLivro), apiKey)

	resposta, err := http.Get(urlPesquisa)
	if err != nil {
		http.Error(w, "Erro ao consultar API do Google", http.StatusInternalServerError)
		return
	}
	defer resposta.Body.Close()

	corpo, _ := ioutil.ReadAll(resposta.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Write(corpo)
}

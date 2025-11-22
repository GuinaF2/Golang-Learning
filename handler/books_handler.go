package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type LivroParaSalvar struct {
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Nota   int    `json:"nota"`
}

type GoogleResponse struct {
	Items []struct {
		VolumeInfo struct {
			Title   string   `json:"title"`
			Authors []string `json:"authors"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

type LivroRespostaGET struct {
	Titulo  string   `json:"titulo"`
	Autores []string `json:"autores"`
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	nomeLivro := r.URL.Query().Get("nome")

	if nomeLivro == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("CHAVEBOOKS")
	urlPesquisa := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&key=%s", url.QueryEscape(nomeLivro), apiKey)

	resp, err := http.Get(urlPesquisa)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var googleData GoogleResponse
	json.Unmarshal(body, &googleData)

	var livros []LivroRespostaGET
	for _, item := range googleData.Items {
		l := LivroRespostaGET{
			Titulo:  item.VolumeInfo.Title,
			Autores: item.VolumeInfo.Authors,
		}
		livros = append(livros, l)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livros)
}

func HandleSave(w http.ResponseWriter, r *http.Request) {
	var livroEscolhido LivroParaSalvar

	json.NewDecoder(r.Body).Decode(&livroEscolhido)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livroEscolhido)
}

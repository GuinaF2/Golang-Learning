# 📚 API de Livros com Go (Atividade de Extensão)

API REST desenvolvida em **Go (Golang)** que integra com a Google Books API.

---

## 🛠️ Tecnologias Utilizadas

* **Linguagem:** Go (Golang) 1.20+
* **Roteador:** [Gorilla Mux](https://github.com/gorilla/mux)
* **Config:** [Godotenv](https://github.com/joho/godotenv) (para variáveis de ambiente)
* **Versionamento:** Git & GitHub

---

## ⚡ Guia de Instalação e Configuração

Siga estes passos para rodar o projeto na sua máquina.

### 1. Clone o Repositório
Baixe o código para sua máquina (substitua pelo link do seu repositório):

```bash
git clone [https://github.com/SEU_USUARIO/api-livros-golang.git](https://github.com/SEU_USUARIO/api-livros-golang.git)
cd api-livros-golang
```

### 2\. Instale as Dependências

O Go baixará automaticamente as bibliotecas listadas no `go.mod` (Gorilla Mux e Godotenv):

```bash
go mod tidy
```

### 3\. Configuração de Segurança (.env)

Este projeto requer uma **API Key** do Google Books.

1.  Crie um arquivo chamado `.env` na raiz do projeto.
2.  Adicione sua chave no arquivo:

<!-- end list -->

```env
API_KEY=SuaChaveGiganteDoGoogleAqui
```

> **Nota:** O arquivo `.env` já está configurado no `.gitignore` para não vazar sua senha no GitHub.

-----

## 🚀 Como Rodar

Com tudo configurado, inicie o servidor:

```bash
go run main.go
```

O terminal deverá exibir:

```text
Servidor rodando na porta 8080...
```

-----

## 🔗 Documentação das Rotas (Endpoints)

Você pode testar as rotas usando o **Postman** ou **Insomnia**.

### 1\. Consultar Livro (GET)

Busca livros diretamente na API do Google e retorna uma lista simplificada.

  * **Rota:** `/books/search`
  * **Método:** `GET`
  * **Query Param:** `nome`

**Exemplo de URL:**
`http://localhost:8080/books/search?nome=O Hobbit`

**Resposta (JSON):**

```json
[
    {
        "titulo": "O Hobbit",
        "autores": ["J.R.R. Tolkien"]
    }
]
```

### 2\. Salvar Livro (POST)

Simula o salvamento de um livro escolhido no banco de dados local.

  * **Rota:** `/books/save`
  * **Método:** `POST`
  * **Header:** `Content-Type: application/json`

**Corpo da Requisição (Body JSON):**

```json
{
    "titulo": "O Hobbit",
    "autores": "J.R.R. Tolkien",
}
```

**Resposta:**
O servidor retornará o mesmo objeto JSON confirmando o recebimento.

-----

*Desenvolvido durante o Workshop de Backend com Golang.*

```
```

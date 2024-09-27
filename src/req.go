package src

import (
    "net/http"
)

// Request representa uma requisição HTTP básica
type Request struct {
    Method  string            // Método HTTP, por exemplo, "GET", "POST"
    URL     string            // URL para a qual a requisição será feita
    Headers http.Header       // Headers da requisição
    Body    []byte            // Corpo da requisição (opcional, usado em POST, PUT, etc.)
}

// Função para criar uma nova requisição
func NewRequest(method, url string) *Request {
    return &Request{
        Method:  method,
        URL:     url,
        Headers: make(http.Header),
    }
}

func MakeRequest() {
    // Exemplo de criação de uma requisição GET
    req := NewRequest("GET", "https://api.exemplo.com/dados")
    req.Headers.Set("Content-Type", "application/json")
    req.Headers.Set("Authorization", "Bearer seu_token_aqui")

    // Aqui você pode adicionar lógica para enviar a requisição usando o pacote net/http
    // Por enquanto, apenas imprimimos os detalhes da requisição
    println("Método:", req.Method)
    println("URL:", req.URL)
    for key, values := range req.Headers {
        for _, value := range values {
            println("Header:", key, "=", value)
        }
    }
}


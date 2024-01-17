# Rate Limiter Middleware for Go
## Overview
Este middleware implementa um sistema de limitação de taxa de requisições (rate limiting) para aplicações web em Go. Ele é projetado para prevenir o excesso de uso dos recursos do servidor limitando o número de requisições ou ações permitidas por um token ou endereço IP em um determinado período de tempo.

# Features
- Configuração flexível de armazenamento para o controle do rate limit.
- Extração automática do endereço IP do solicitante.
- Suporte para cabeçalhos X-Forwarded-For e X-Real-IP.
- Fácil integração com aplicações Go existentes.

# Installation
Para utilizar este middleware, inclua o pacote em seu projeto Go:

```go
import "github.com/booscaaa/desafio-rate-limiter-go-expert-pos/ratelimiter"
```

# Configuração
Antes de utilizar o middleware, você precisa inicializá-lo com as opções desejadas. Por exemplo, para definir o repositório de armazenamento:

## Exemplo de uso com GIN-GONIC
```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/booscaaa/desafio-rate-limiter-go-expert-pos/ratelimiter"
)

func main() {
	ratelimiter.Initialize()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		ratelimiter.Middleware(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	router.Run(":8080")
}
```

## Exemplo de uso com GORILLA MUX
```go
package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/booscaaa/desafio-rate-limiter-go-expert-pos/ratelimiter"
)

func main() {
	ratelimiter.Initialize()
	router := mux.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return ratelimiter.Middleware(next)
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", router)
}
```

## Exemplo de uso com NET/HTTP

```go
package main

import (
	"net/http"
	"github.com/booscaaa/desafio-rate-limiter-go-expert-pos/ratelimiter"
)

func main() {
	ratelimiter.Initialize()
	http.Handle("/", ratelimiter.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})))

	http.ListenAndServe(":8080", nil)
}
```

## Exemplo de uso com GO-CHI

```go
package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/booscaaa/desafio-rate-limiter-go-expert-pos/ratelimiter"
)

func main() {
	ratelimiter.Initialize()
	router := chi.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return ratelimiter.Middleware(next)
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", router)
}
```

O middleware vai verificar se o limite de requisições foi atingido para o token ou IP específico. Se o limite for excedido, uma resposta HTTP 429 (Too Many Requests) será enviada.

# Customization
Você pode personalizar o comportamento do middleware ajustando as opções de LimiterOpts e implementando sua própria lógica de armazenamento se necessário.
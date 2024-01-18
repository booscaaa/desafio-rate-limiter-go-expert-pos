# Rate Limiter Middleware for Go
Este middleware implementa um sistema de limitação de taxa de requisições (rate limiting) para aplicações web em Go. Ele é projetado para prevenir o excesso de uso dos recursos do servidor limitando o número de requisições ou ações permitidas por um token ou endereço IP em um determinado período de tempo.

# Estrutura

- ratelimiter: pasta com o fonte do package para ser usado em seu router.

- api: exemplo de uso do package em uma api com gin-gonic.

# Features
- Configuração flexível de armazenamento para o controle do rate limit.
- Extração automática do endereço IP do solicitante.
- Suporte para cabeçalhos X-Forwarded-For e X-Real-IP.
- Fácil integração com aplicações Go existentes.

# Como usar
Para utilizar este middleware, inclua o pacote em seu projeto Go:

```go
import "github.com/booscaaa/desafio-rate-limiter-go-expert-pos/ratelimiter"
```

# Configuração
Antes de utilizar o middleware, você precisa inicializá-lo com as opções desejadas. Ele tem suporte para `json` e `.env`. 

OBS: O `.env` é utilizado como padrão caso tenha os dois configurados.

## Configuração em JSON
Por exemplo, para definir todas as opções disponíveis no json use:

| Chave        | Descrição                                                       |
|--------------|-----------------------------------------------------------------|
| `limiter`    | Objeto raiz para as configurações do limitador de requisições.  |
| `database`   | Define as configurações do banco de dados para armazenar dados de limitação. |
| `inMemory`   | Se `true`, usa a memória para armazenar dados de limitação. Se `false`, não usa a memória. |
| `redis`      | Se `true`, usa o Redis para armazenar dados de limitação. Se `false`, não usa o Redis. |
| `default`    | Define limites padrões para requisições.                         |
| `requests`   | Número máximo de requisições permitidas.                          |
| `every`      | Intervalo de tempo (em segundos) para o limite de requisições.    |
| `ips`        | Array de objetos especificando limites para IPs individuais.    |
| `ip`         | Endereço IP específico para aplicar limites de requisição.      |
| `tokens`     | Array de objetos especificando limites para tokens de acesso.   |
| `token`      | Token de acesso específico para aplicar limites de requisição.  |

### Exemplo

Crie um arquivo na raiz do seu projeto chamado `config.json`.

```json
{
  "limiter": {
    "database": {
      "inMemory": true,
      "redis": false
    },
    "default" : {
      "requests": 10,
      "every": 1
    },
    "ips": [
      {
        "ip": "127.0.0.1",
        "requests": 10,
        "every": 1
      }
    ],
    "tokens": [
      {
        "token": "123456",
        "requests": 4,
        "every": 3
      }
    ]
  }
}
```

## Configuração em .env
Por exemplo, para definir todas as opções disponíveis no .env use:

| Variável                         | Descrição                                                       |
|----------------------------------|-----------------------------------------------------------------|
| `LIMITER_DATABASE_INMEMORY`      | Se `false`, não usa a memória para armazenar dados de limitação. Se `true`, usa a memória. |
| `LIMITER_DATABASE_REDIS`         | Se `true`, usa o Redis para armazenar dados de limitação. Se `false`, não usa o Redis. |
| `LIMITER_DEFAULT_REQUESTS`       | Número máximo padrão de requisições permitidas.                 |
| `LIMITER_DEFAULT_EVERY`          | Intervalo de tempo padrão (em segundos) para o limite de requisições. |
| `LIMITER_IPS_0_IP`               | Endereço IP específico (ex.: 127.0.0.1) para aplicar limites de requisição. |
| `LIMITER_IPS_0_REQUESTS`         | Número máximo de requisições permitidas para o IP especificado. |
| `LIMITER_IPS_0_EVERY`            | Intervalo de tempo (em segundos) para o limite de requisições para o IP especificado. |
| `LIMITER_TOKENS_0_TOKEN`         | Token de acesso específico (ex.: 123456) para aplicar limites de requisição. |
| `LIMITER_TOKENS_0_REQUESTS`      | Número máximo de requisições permitidas para o token especificado. |
| `LIMITER_TOKENS_0_EVERY`         | Intervalo de tempo (em segundos) para o limite de requisições para o token especificado. |
| `LIMITER_TOKENS_1_TOKEN`         | Outro token de acesso específico (ex.: abc543) para aplicar limites de requisição. |
| `LIMITER_TOKENS_1_REQUESTS`      | Número máximo de requisições permitidas para o segundo token especificado. |
| `LIMITER_TOKENS_1_EVERY`         | Intervalo de tempo (em segundos) para o limite de requisições para o segundo token especificado. |

### Exemplo

Crie um arquivo na raiz do seu projeto chamado `.env`.

```env
LIMITER_DATABASE_INMEMORY=false
LIMITER_DATABASE_REDIS=true

LIMITER_DEFAULT_REQUESTS=10
LIMITER_DEFAULT_EVERY=1

LIMITER_IPS_0_IP=127.0.0.1
LIMITER_IPS_0_REQUESTS=10
LIMITER_IPS_0_EVERY=1

LIMITER_TOKENS_0_TOKEN=123456
LIMITER_TOKENS_0_REQUESTS=4
LIMITER_TOKENS_0_EVERY=3

LIMITER_TOKENS_1_TOKEN=abc543
LIMITER_TOKENS_1_REQUESTS=5
LIMITER_TOKENS_1_EVERY=1
```

# Adicionando o middlewarre ao seu router preferido

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

	router.Use(ratelimiter.Middleware)

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

	router.Use(ratelimiter.Middleware)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", router)
}
```

O middleware vai verificar se o limite de requisições foi atingido para o token ou IP específico. Se o limite for excedido, uma resposta HTTP 429 (Too Many Requests) será enviada.

# Rodando

Clone ou baixe o repositório para sua máquina local.

```bash
git clone https://github.com/booscaaa/desafio-rate-limiter-go-expert-pos.git
```

### Navegue até o diretório do projeto

```bash
cd desafio-rate-limiter-go-expert-pos
```

### Execute com o docker
Irá subir o servidor na pasta api na porta 8080 e também o redis para realizar os testes.

```bash
docker compose up --build -d
```

# Testando
Para realizar os testes e verificar as configurações feitas no arquivo `json` ou `.env` podemos usar o desafio do stress test criado.

Vamos configurar o `.env` da seguinte forma:

```bash
LIMITER_DATABASE_INMEMORY=false
LIMITER_DATABASE_REDIS=true

LIMITER_DEFAULT_REQUESTS=10
LIMITER_DEFAULT_EVERY=1

LIMITER_IPS_0_IP=127.0.0.1
LIMITER_IPS_0_REQUESTS=10
LIMITER_IPS_0_EVERY=1

LIMITER_TOKENS_0_TOKEN=123456
LIMITER_TOKENS_0_REQUESTS=4
LIMITER_TOKENS_0_EVERY=3

LIMITER_TOKENS_1_TOKEN=abc543
LIMITER_TOKENS_1_REQUESTS=5
LIMITER_TOKENS_1_EVERY=1
```

### Testando as configurações default
Para realizar o teste podemos rodar passando tokens aleatórios ou em um ip que não seja o 127.0.0.1, rode:

```bash
docker run --network=host --rm booscaaa/desafio-stress-test-go-expert-pos --url http://localhost:8080 --concurrency 2 --requests 40
```

### Saída

```bash
Test completed
----------------------------------------------------------------------
Total requests: 40
----------------------------------------------------------------------
Successful requests: status code 200; total 10
----------------------------------------------------------------------
Requests with error: status code 429; total 30
----------------------------------------------------------------------
Total execution time: 13.651292ms
```
Como o default está configurado para 10 requests por segundo, 10 funcionaram corretamente e 30 retornaram status code `429`.

### Testando as configurações default com um token aleatório não cadastrado

```bash
 docker run --network=host --rm booscaaa/desafio-stress-test-go-expert-pos --url http://localhost:8080?token=3444444 --concurrency 2 --requests 40
```

### Saída

```bash
Test completed
----------------------------------------------------------------------
Total requests: 40
----------------------------------------------------------------------
Successful requests: status code 200; total 10
----------------------------------------------------------------------
Requests with error: status code 429; total 30
----------------------------------------------------------------------
Total execution time: 13.740945ms
```

Como o token `3444444` não está no cadastro, o rate limiter usa a configuração default. Sendo assim, novamente 10 funcionaram corretamente e 30 retornaram status code `429`.

### Testando com o token 123456
O token tem a seguinte configuração. `4` requests a cada `3` segundos

```bash
LIMITER_TOKENS_0_TOKEN=123456
LIMITER_TOKENS_0_REQUESTS=4
LIMITER_TOKENS_0_EVERY=3
```

Rode:

```bash
docker run --network=host --rm booscaaa/desafio-stress-test-go-expert-pos --url http://localhost:8080?token=123456 --concurrency 2 --requests 40
```

### Saída

```bash
Test completed
----------------------------------------------------------------------
Total requests: 40
----------------------------------------------------------------------
Successful requests: status code 200; total 4
----------------------------------------------------------------------
Requests with error: status code 429; total 36
----------------------------------------------------------------------
Total execution time: 12.779442ms
```

# Storage

O middleware suporta atualmente o `redis` e também `inmemory`. Configurando aqui:

```bash
LIMITER_DATABASE_INMEMORY=false
LIMITER_DATABASE_REDIS=true
```

Também é possível implementar seu próprio storage implementando a interface:

```go
type DatabaseRepository interface {
	Create(context.Context, RateLimiterInfo, time.Duration) error
	Read(context.Context, string) (*RateLimiterInfo, error)
	CheckLimit(context.Context, string, int, time.Duration) (bool, error)
}
```

E injetando sua implementação dessa forma ao inicializar o rate limiter: 

```go
func main() {
->	ratelimiter.Initialize(
->		ratelimiter.Storage(IMPLEMTANTACAO_DO_STORAGE),
->	)

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
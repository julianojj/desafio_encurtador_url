# desafio_encurtador_url
Desafio encurtador de url proposto pela plataforma [Devgym](https://www.devgym.com.br/)

## Tecnologias utilizadas:
```
uuid v1.3.0
testify v1.8.1
pq v1.10.7
godotenv v1.4.0
gin v1.8.2
```

## Como rodar o projeto?
Devemos criar um arquivo 2 arquivos, .env e .db.env com as seguintes configurações:

.env deve conter o `BASE_URL` para a conexão com o banco de dados
```
BASE_URL="postgres://username:password@host:port/dbname?sslmode=disable"
```

.db.env
```
POSTGRES_USER=username
POSTGRES_PASSWORD=password
POSTGRES_DB=dbname
```
Devemos rodar o comando `docker compose up -d` para subir o banco de dados

## Comandos utilizados:
```
go test ./...
go run ./...
```

## API
Há 2 rotas disponíveis:

POST /cut passando um json: { "url": "myurl" } retorna { "short": "mycode" }

GET /uncut?code=mycode retorna { "url": "myurl" }

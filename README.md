# PDI 2024-2

Este é um projeto para o plano de desenvolvimento do segundo semestre de 2024.

### Iniciar projeto

Primeiramente, utilize clone o repositório:
`git clone https://github.com/CarlosEduardoNop/pdi-2024-2-api.git`

Após clonar o repositório, é necessário criar o arquivo .env:
`cp .env.example env`

### Comandos iniciais

Para subir o projeto basta rodar o comando `docker compose up -d` e após rodar `go mod tidy`


### Comandos

- `go run ./cmd/artisan migrate` - Comando para rodar as migrations que ainda não foram executadas.
- `go run ./cmd/artisan migration --name=` - Comando para criar uma nova migration, podendo passar o nome. Será criado um arquivo .sql na pasta *migrations*.

### Versões

- Go: 1.22.2
- MYSql: 8.0.0
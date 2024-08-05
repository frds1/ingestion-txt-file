FROM golang:1.22.5 AS builder

# Defina o diretório de trabalho
WORKDIR /app

# Copie o go.mod e go.sum e baixe as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copie o código-fonte da aplicação
COPY . .

# Compile a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Defina a porta que a aplicação irá expor
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]

# Usar uma imagem base do Go
FROM golang:1.20-alpine

# Definir o diretório de trabalho
WORKDIR /backend

# Copiar o código fonte
COPY . .

# Baixar dependências
RUN go mod download

# Compilar a aplicação
RUN go build -o main ./cmd/backend/main.go

# Expor a porta
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]
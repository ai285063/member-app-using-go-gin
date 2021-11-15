FROM golang:latest
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN chmod +x ./wait-for-it.sh

EXPOSE 8080
CMD ["go", "run", "."]

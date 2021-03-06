FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go get "github.com/go-chi/chi"
RUN go get "github.com/go-chi/chi/middleware"
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/go-chi/jwtauth"
RUN go get "github.com/dgrijalva/jwt-go"
RUN go get "github.com/go-chi/cors"

EXPOSE 3000

RUN go build -o main .

CMD ["/app/main"]

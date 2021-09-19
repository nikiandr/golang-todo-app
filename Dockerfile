FROM golang:1.17.1

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /golang-todo-app

CMD [ "/golang-todo-app" ]
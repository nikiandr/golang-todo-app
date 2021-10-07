# To Do REST Api

REST API written in Golang using standard net/http package and [Gin](https://github.com/gin-gonic/gin "Gin Web Framework") framework. PostgreSQL DB used as a storage. This application implements To-do List model with CRUD operations.

## Authentication

Access to Sign-in and Sign-up endpoints doesn't require any authentication. All other endpoints require [JWT](https://jwt.io/ "Json Web Token") token that could be acquired on the first two endpoints.

## Database

![schema](/schema/schema.png)

On the picture above you can see DB schema implemented in PostgreSQL syntax ([here](schema/up.sql)). Tables ***users***, ***todo_items*** and ***todo_lists*** are tables which contain information about users, items - parts of lists and lists accordingly. Table ***users_lists*** is used to model many-to-many relationship between ***users*** and ***todo_lists*** tables.

## API endpoints

Particular information on endpoints available could be found in [Postman generated documentation](https://documenter.getpostman.com/view/8056238/UUy4cQqr).

## Build & Run

To prepare for build firstly you need to create *.env* file in working directory:

```bash
DB_PASSWORD=<password from database>
JWT_SIGNING_SECRET=<secret for jwt token generation>
PORT=<database port>
```

Then you need too deploy Postgres DB with schema from **up.sql** file and specify connection details in *configs/config.yml* file. 

To deploy project locally you may use Docker image from schema folder. Firstly you need to build your docker image by running command (using *schema* folder as your active one):

```bash
docker build -t postgres_todo .
```

Then you will need to run container from built image specifying you DB port, password and container name of choice e.g.:

```bash
docker run --name=todo_postgres -e POSTGRES_PASSWORD='<password of choice>' -p 5432:5432 -d postgres_todo
```

After that you can build project by running from project folder:

```bash
go build cmd/main.go
```

and then running executable file created.

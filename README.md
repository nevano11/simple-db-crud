# simple-db-crud
## Pet-проект на ЯП Golang с поддержкой CRUD операций на PostgreSql

## Используемые технологии
### Логгирование
    logrus github.com/sirupsen/logrus
### Web framework
    gin github.com/gin-gonic/gin

### База данных
Postgres on docker

    docker run --name simple-crud-postgres -v ./postgres-data:/var/lib/postgresql/data -e POSTGRES_DB=simple-crud -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5000:5432 -d postgres
**Database migrations**

Installing

    scoop install migrate
Up/down commands

    migrate -source file://schema -database postgres://postgres:password@localhost:5000/simple-crud?sslmode=disable up 2
    migrate -source file://schema -database postgres://postgres:password@localhost:5000/simple-crud?sslmode=disable down
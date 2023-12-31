# simple-db-crud - Pet-проект Golang с поддержкой CRUD операций

## Руководство по запуску
1 - Бд в докере

    docker run --name simple-crud-postgres -v ./postgres-data:/var/lib/postgresql/data -e POSTGRES_DB=simple-crud -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5000:5432 -d postgres
2 - Накатить скрипты из schema (Вручную или с помощью migrate)

    migrate -source file://schema -database postgres://postgres:password@localhost:5000/simple-crud?sslmode=disable up 2
3 - Наслаждаться приложением


## Используемые технологии
### Logging
    logrus github.com/sirupsen/logrus
### Config
    viper github.com/spf13/viper
### Web framework
    gin github.com/gin-gonic/gin

### Database
**Postgres on docker**

    docker run --name simple-crud-postgres -v ./postgres-data:/var/lib/postgresql/data -e POSTGRES_DB=simple-crud -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5000:5432 -d postgres
**Database migrations**

Installing with scoop (for Windows)

    scoop install migrate
Up/down commands

    migrate -source file://schema -database postgres://postgres:password@localhost:5000/simple-crud?sslmode=disable up 2
    migrate -source file://schema -database postgres://postgres:password@localhost:5000/simple-crud?sslmode=disable down
Using sql on code

    sqlx github.com/jmoiron/sqlx
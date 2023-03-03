# Cash Cash
This project is a finance api built with :heart: and Go.

## Tools

- [Go](https://go.dev/)
- [Echo](https://echo.labstack.com/)
- [Gorm](https://gorm.io/)
- [Postgres](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)

## Getting started

1. **Setup the database**
    
    How: change to deployments folder and do: `docker compose up -d`

2. **Run project**

    How: from the root of the project, do: `go run cmd/main/main.go`

## API

<details open>
<summary>Incomes</summary>

### List all incomes

**GET** /income

CURL:

`curl --request GET \
--url http://localhost:1323/income`


</details>

<details open>
<summary>Outcomes</summary>

</details>
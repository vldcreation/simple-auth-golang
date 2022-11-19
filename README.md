# Simple-Auth-Golang

## Instalation
1. copy .env.example and setup your own env.make sure you fill in all required env
```
cp .env.example .env
```
2. download all depedencies
```
go mod download
```
### Instalation - Without Docker
4. First you need to Make sure your psql connection is connected.Run app
```
DB_HOST=<DB_HOST> go run main.go
```
5. While development, just repeat step 4 to get changes

### Instalation - With Docker
6. First you need to make sure you have added your own env as listed in the service environment inside docker-compose

7. Build app using docker
```
make docker-build
```

7. Run all container
```
make docker-up
```
Also you can combine those two command with
``` docker-compose up --build ``` and if you need to run in background proccess add argument -d.see [documentation](https://docs.docker.com/engine/reference/commandline) for details.

8. While development to get changes from application, just rebuild your service-app
```
make edufund-rebuild
```
and restart then
```
make edufund-restart
```

If you want to rebuild all containers, do these steps sequentially
* ``` make docker-down ```
* ``` make docker-build ```
* ``` make docker-up ```


## Project Layer
- Models
    - Same as Entities, will used in all layer. This layer, will store any Object’s Struct and its method
- Repository
    - Repository will store any Database handler. Querying, or Creating/ Inserting into any database will stored here. This layer will act for CRUD to database only. No business process happen here. Only plain function to Database.
- Feature & Usecase
    - This layer will act as the business process handler. Any process will handled here. This layer will decide, which repository layer will use. And have responsibility to provide data to serve into delivery. Process the data doing calculation or anything will done here
- Delivery
    - This layer will act as the presenter. Decide how the data will presented.

## Guideline
- [Clean Code Using SOLID Principles][#solid-baeldung]
	- uncle bob's [explanation][#solid-unclebob]
- [Standard Go Project Layout][#go-layout] :3
	- multiple entrypoints of building a binary executable
- [Makefile][#makefile]
	- productive command dictionary
- [UnitTest](https://go.dev/doc/tutorial/add-a-test)
    - validate that each unit of the software code performs as expected

See coverage testing result
```
make test-unit
```

## Library That I Use
- sqlx/db
    - for sql driver
- godotenv
    - for local environtment management
- gin
    - for serve http
- jwt
    - for auth security issue
- go-sqlmock, gomonkey, testify
    - for testing purpose

## API Contract

| Endpoint        | Path           | Interface  |
| ------------- |:-------------:| -----:|
| `POST` /api/v1/user/account-creation      | [internal/feature/contract.user.go] | SetupUser |
| `POST` /api/v1/user/account-login      | [internal/feature/contract.user.go]      |   AccountLogin |


<!-- REFERENCE LINKS -->
[#solid-baeldung]:	https://www.baeldung.com/solid-principles
[#solid-unclebob]:	https://blog.cleancoder.com/uncle-bob/2020/10/18/Solid-Relevance.html
[#go-layout]:		https://github.com/golang-standards/project-layout
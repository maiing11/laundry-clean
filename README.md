# Laundry clean

a submit project clean architecture with gin framework, go-fx as dependency container

To learn about project structure and dependency injection please go through [here](https://medium.com/wesionary-team/dependency-injection-with-go-fx-b698a6585cf0?source=friends_link&sk=26f391ae41c493946ee3434be2ed4971)

## Running the project
- Make sure you have docker installed.
- Copy `.env.example` to `.env` adjust your configuration
- Run `docker-compose up -d`

If youre running without docker be sure database configuration is provided in `.env` file.

- Run `go run main.go app:serve` in terminal 


#### Environment Variables

<details>
    <summary>Variables Defined in the project </summary>

| Key            | Value                    | Desc                                        |
| -------------- | ------------------------ | ------------------------------------------- |
| `API_PORT`     | `8080`                   | Port at which app runs                      |
| `LOG_OUTPUT`   | `./server.log`           | Output Directory to save logs               |
| `POSTGRES_USER`| `username`               | Database Username                           |
| `POSTGRES_PASS`| `password`               | Database Password                           |
| `POSTGRES_HOST`| `0.0.0.0`                | Database Host                               |
| `POSTGRES_PORT`| `5432`                   | Database Port                               |
| `POSTGRES_NAME`| `laundry`                | Database Name                               |
| `JWT_SECRET`   | `secret`                 | JWT Token Secret key                        |

</details>


## Implemented Features

- Dependency Injection (go-fx)
- Routing (gin web framework)
- Environment Files
- Database Setup (postgresql)
- Models Setup 
- Repositories
- Implementing Basic CRUD Operation
- Cobra Commander CLI Support. try: `go run . --help`

## Todos

- [x] COBRA Commander CLI Support 
- [x] Use of Interfaces
- [ ] Logging with logrus (file saving on local file ./server.log)
- [ ] Middleware cors
- [ ] File upload middelware. 
- [ ] Authentication (JWT)
- [ ] Unit Testing

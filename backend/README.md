# Backend

## Starting the Docker Container & Server Application

Start the docker container for the PostgreSQL database:

```sh
docker compose up -d
```

To start the backend, use the following command where `<your_postgres_dsn>` is replaced with your actual PostgreSQL DSN:

```sh
DSN=<your_postgres_dsn> go run server.go
```

## Initial Setup

To start the backend application, you have to perform a few steps:

1. Make sure [Docker](https://docs.docker.com/get-docker/) and [Go](https://go.dev/doc/install) are installed on your system. Also, install `sqlc`. Mac/Linux `brew install sqlc` or download an executable, (e.g. `sqlc.exe` on Windows) and add it to your path.
2. `cd .\backend\` (change your CLI in the backend directory)
3. Start Docker Desktop, then boot up a PostgreSQL database in a docker container with: `docker compose up -d` [^1]
4. For start the backend application the `DSN` environment variable needs to be set (basically a DB connection string). You can set the environment variable for the current CLI session, use the Goland IDE [^2], or set it permanently on your system. `<your_postgres_dsn>` is a placeholder for the PostgreSQL DSN from the `.env.template` file:
   - With Windows Powershell: `$env:DSN="<your_postgres_dsn>"`. Linux / MacOS / Wind (CMD) `set DSN=<your_postgres_dsn>`. Then run the server with `go run server.go` or set the environment variable and run the server in one command, e.g. `DSN=<your_postgres_dsn> go run server.go`
   - You can permanently set environmental variables in the settings of your operating system (on Windows) or with commands in the footnote (on any OS) [^3]
   - troubleshooting [^4]
5. you are required to connect to the database and execute the migration scripts which can be found in the directory: [/internal/db/migrations](./internal/db/migrations). Fill the database with some dummy seed data. `go run .\cmd\app\app.go`
6. open [`http://localhost:8080/`](http://localhost:8080/) and try to fetch some entity, e.g. `{ todos { id } }`.

[^1]: After running `docker compose up -d`, "course-recommender" > "db-1" should show up as running in Docker Desktop. up: This command is used to create and start containers. When you run `docker compose up -d`, Docker Compose will start all the services defined in your `docker-compose.yml` file. This includes creating and starting any containers, networks, or volumes that are needed. 
The `-d`: flag stands for "detached" mode. It starts the containers in the background so that the CLI is not occupied and can still be used. `docker-compose.yml` defines, which postgres db image version is used and the host port (physical hardware computer) to the container port (virtualized computer). Also, environment variables for the container are predefined etc.

[^2]: For those using the JetBrains IDE GoLand (which we recommended for development), you can utilize the [EnvFile](https://plugins.jetbrains.com/plugin/7861-envfile) plugin. This plugin simplifies the process of binding environment variables from your `.env` file to the application. To configure this, edit your run configuration, add a Go build task, and in the 'EnvFile' tab, link to your `.env` file by copying the contents from `.env.example` into a new `.env` file.

[^3]: Linux and macOS (in bash shell): `echo 'export DSN=<your_postgres_dsn>' >> ~/.bash_profile`.  zsh shell (default in recent versions of macOS): `echo 'export DSN=<your_postgres_dsn>' >> ~/.zshrc`. Powershell `[System.Environment]::SetEnvironmentVariable('DSN', '<your_postgres_dsn>', [System.EnvironmentVariableTarget]::User)`.

[^4]: If you encounter database connection or authentication errors in the Graphi/QL Playground, make sure the environment variable is properly set and the host port, which is specified in the `docker-compose.yml` file, is not used yet (e.g. an already existing postgres installation and [service](C:\Windows\System32\services.msc) might have the same port). In that case, change the host port in the file and DSN.

## Development Instructions

### GraphQL Generation with gqlgen

For our GraphQL server, we employ `gqlgen`. This tool automatically generates Go code according to your GraphQL schema definitions.

When you make changes to the GraphQL schema, you will need to regenerate the associated resolvers and related code. To do so, run the following command:

```sh
go run github.com/99designs/gqlgen generate
```

### Code Generation with sqlc

We use `sqlc` for generating code based on our SQL queries. This generates go type definitions for db entities.

If you modify or introduce new SQL queries, execute the command below to regenerate the necessary code:

```sh
sqlc generate
```

Make sure to run this whenever you make changes to ensure the database interactions are up to date. Tools like nodemon could be configured to listen for file changes in the relevant directories to rerun these generative commands.

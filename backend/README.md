# Backend Setup

## Starting the Backend

To start the backend, use the following command where `<your_postgres_dsn>` is replaced with your actual PostgreSQL DSN:

```sh
DSN=<your_postgres_dsn> go run server.go
```

Before running this command, ensure that both Docker and Go are installed on your system. To initiate the PostgreSQL database, execute:

```sh
docker compose up -d
```

This command will set up a Postgres database container. Following that, you are required to connect to the database and execute the migration scripts which can be found in the directory: [/internal/db/migrations](./internal/db/migrations).

For those using GoLand (which is recommended for this project) to develop, you can utilize the [EnvFile](https://plugins.jetbrains.com/plugin/7861-envfile) plugin. This plugin simplifies the process of binding environment variables from your `.env` file to the application. To configure this, edit your run configuration, add a Go build task, and in the 'EnvFile' tab, link to your `.env` file by copying the contents from `.env.example` into a new `.env` file.

## Development Instructions

### GraphQL Generation with gqlgen

For our GraphQL server, we employ `gqlgen`. This tool automatically generates Go code according to your GraphQL schema definitions.

When you make changes to the GraphQL schema, you will need to regenerate the associated resolvers and related code. To do so, run the following command:

```sh
go run github.com/99designs/gqlgen generate
```

### Code Generation with sqlc

We use `sqlc` for generating code based on our SQL queries.

If you modify or introduce new SQL queries, execute the command below to regenerate the necessary code:

```sh
sqlc generate
```

Make sure to run this whenever you make changes to ensure the database interactions are up to date.
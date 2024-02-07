# Backend

## How to start the backend

    DSN=<your postgres dsn> go run server.go

Before you do that you must install Docker and of course Go. Then to start up the database, you can run `docker compose up -d`. This spins up a Postgres database. Now you need to manually (for now) connect to the database, and run the migrations. You can find the migrations in [/internal/db/migrations](./internal/db/migrations).

If you are using GoLand (recommended) you can use the [EnvFile](https://plugins.jetbrains.com/plugin/7861-envfile) plugin to auto bind the env variables from you .env file. For that, just edit your run configuration, add a go build and there you should see an extra tap (EnvFile) where you can point to an env file (.env). Just copy `.env.example` to `.env`.

## Development

We use `gqlgen` for our graphql server. It auto generates Go code for you based on your graphql schema definitions. 

Everytime you change the GraphQL schema, you can have to regenerate your resolvers etc. To do that, just run:

    go run github.com/99designs/gqlgen generate

We use `sqlc` to generate based on our 

Everytime you adjust or add sql queries, you have to run:

    sqlc generate
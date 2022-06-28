# MealPlannerApi

A MealPlanner API
version: 0.9.0

#### This is a MealPlanner application which can work with REST, GraphQL, gRPC and CLI interfaces.

#### The App's latest version is 1.0.0

## Requirements

- Docker Engine
- DockerCompose script aka docker-compose
- *nix system

## Technical stack

- [go-chi](https://github.com/go-chi/chi/)
- [JWT](https://github.com/go-chi/jwtauth/)
- [UUID](https://github.com/google/uuid/)
- [GraphQL](https://github.com/99designs/gqlgen/)
- [gRPC](https://google.golang.org/grpc/)
- [MongoDB](https://go.mongodb.org/mongo-driver/)
- [Redis](https://github.com/redis/go-redis/)
- [MessageBus](https://github.com/vardius/message-bus/)
- [Logging](https://github.com/sirupsen/logrus/)
- [SwaggerUI](https://hub.docker.com/r/swaggerapi/swagger-ui)
- [MongoExpress](https://hub.docker.com/_/mongo-express)

## Development stack

- Domain-Driven-Design
  - Aggregates
  - Entities
  - DataTransferObjects
  - Repositories
  - EntityManager
  - CacheManager
  - Factories
  - UserInterface
- CMD

## Release

- For farther information please go to [CHANGELOG.md](./CHANGELOG.md)
- Dependencies are updated before the last commit

## Contribute

- There is no contribution scheme for the projects. Sorry.

## Preparing to start

- create a .env.local file in the root
- put and change everything you need there
- wisely choose every value

## Start services

### DockerCompose

#### Use MongoDB

- Bring service alive

```sh
docker-compose -f service/docker/docker-compose-mongo.yml up -d
```

- Pop-up to the application

```sh
docker exec -ti docker-api-N sh
```

where N is a number which is chosen by docker engine

### CLI

- Start the CLI application for using

```sh
go run ui/cmd/cli/main.go
```

You can use -command= for applying commands immediately after starting.
Anyway, you can use any command in the application and the `help` command for listing of all commands.

### REST

- Start the REST application for using

```sh
go run ui/cmd/rest/main.go -cors=true -dev=true -jwtAuthentication=true -contentTypeJSON=true
```

You can use swagger for testing API.

### GraphQL

- Start GraphQL application for using

```sh
go run ui/cmd/graphql/main.go -cors=true -dev=true -jwtAuthentication=true -contentTypeJSON=true
```

You can use the GraphQL's playground for testing API.

### gRPC

- Start gRPC application for using

```sh
go run ui/cmd/grpc/main.go -gRPCPort=50051 -dev=true
```

You can use the gRPC's client `ui/grpc/client/main.go` for testing API.

## Testing

- Go to the app container as

```sh
docker exec -ti docker-api-N sh
```

- Start the tests as

```sh
go test -v -cover ./...
```

## Migrations

- For MongoDB no migrations are required.

## ToDo

- in collection add validation
  data https://www.mongodb.com/docs/manual/core/schema-validation/update-schema-validation/#std-label-schema-update-validation
- cover validating for any entity and aggregate
- covering at least 90% of tests
- cover GraphQL functionality
- cover gRPC functionality
- cover a wizard for the CLI application

## Licence

- see [Apache 2.0](./LICENSE)
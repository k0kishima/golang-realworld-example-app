# ![RealWorld Example App](https://github.com/gothinkster/realworld-starter-kit/raw/master/logo.png)

> ### Golang implementation of RealWorld app that adheres to the [RealWorld](https://github.com/gothinkster/realworld) spec and API.

### [Demo](https://github.com/gothinkster/realworld)&nbsp;&nbsp;&nbsp;&nbsp;[RealWorld](https://github.com/gothinkster/realworld)

This is a RealWorld example application built with Gin and the Ent framework. It demonstrates a backend implementation for a blogging platform with features such as articles, comments, tags, and user profiles.

You might also check out [the Frontend implementation in Nuxt3](https://github.com/k0kishima/nuxt3-realworld-example-app).

For more information on how to this works with other frontends/backends, head over to the [RealWorld](https://github.com/gothinkster/realworld) repo.


## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.16 or higher

### Setting Up the Development Environment

#### 1. Start the MySQL database container:

Use the provided `docker-compose.yml` file to start a MySQL container:

```bash
$ docker-compose up -d
```

This will create a new MySQL database named golang_realworld.

#### 2. Configure environment variables:

Copy the example environment file and edit it to match your local setup:

```bash
$ cp .env.example .env
$ vim .env
```

#### 3. Start the server:

Run the following command to start the server:

```bash
$ go run main.go
```

The server will automatically create the required database schema if the connection to the database is successful.

### Running Tests

To run the end-to-end tests, execute the following command:

```bash
$ make test-e2e
```

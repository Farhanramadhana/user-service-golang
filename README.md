# Backend Engineering Interview Assignment (Golang)

## Requirements

To run this project you need to have the following installed:

1. [Go](https://golang.org/doc/install) version 1.19
2. [Docker](https://docs.docker.com/get-docker/) version 20
3. [Docker Compose](https://docs.docker.com/compose/install/) version 1.29
4. [GNU Make](https://www.gnu.org/software/make/)
5. [oapi-codegen](https://github.com/deepmap/oapi-codegen)

    Install the latest version with:
    ```
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    ```
6. [mock](https://github.com/golang/mock)

    Install the latest version with:
    ```
    go install github.com/golang/mock/mockgen@latest
    ```

## Initiate The Project
Before we begin, please follow these steps:

1. Copy the values from env.example to a new file named .env, which includes examples of JWT private and public keys.
2. Generate a private and public RSA key pair using the following commands:

a. openssl genrsa -out jwtRSA256-private.pem 2048
b. openssl rsa -in jwtRSA256-private.pem -pubout -outform PEM -out jwtRSA256-public.pem

Note: please generate on PKCS8 format, for testing purpose, the key can generate online from https://acte.ltd/utils/openssl

3. Copy the value from jwtRSA256-private.pem to the JWT_PRIVATE_KEY variable in the .env file.
4. Copy the value from jwtRSA256-public.pem to the JWT_PUBLIC_KEY variable in the .env file.

To start working, execute

```
make init
```

## Running

To run the project, run the following command:

```
docker-compose up --build
```

You should be able to access the API at http://localhost:8080

If you change `database.sql` file, you need to reinitate the database by running:

```
docker-compose down --volumes
```

## Testing

To run test, run the following command:

```
make test
```

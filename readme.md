## How to setup

1. Clone this repository.
2. Copy the `.env.example` file to `.env` and fill in the required environment variables.
3. Run `docker-compose -f ./deployments/docker-compose.yaml --env-file .env up -d` to start the server.

## Tech Stack

- Golang as the backend language
- PostgreSQL as the database server
- Gorm as the ORM
- Fiber as the web framework
- Postman as the frontend

## Features

- Daftar Nasabah Baru
- Tabung Saldo Rekening
- Tarik Saldo Rekening
- Cek saldo Rekening

## Entity Relationship Diagram

![ERD](docs/erd.png)

## API Endpoints

The API endpoints can be found in the [OpenAPI Definition](docs/account.yml) and the [Postman Collection](docs/postman_collection.json).


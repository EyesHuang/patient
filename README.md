# Patient Service

This project primarily provides two web services for interacting with patient data.

## Available Services

- `GET /patients`: Retrieve information for all patients.
- `PUT /orders/{id}`: Update a specified medical order using its ID.

## Getting Started

### Setting Up PostgreSQL

Follow these steps to establish a PostgreSQL database:

1. Navigate to the `deployment/postgresql` directory and create a new `.env` file.
2. In the `.env` file, add the following environment variables. You may modify the values according to your requirements:

    ```env
    POSTGRES_DB=mydatabase
    POSTGRES_USER=myuser
    POSTGRES_PASSWORD=mypassword
    ```

3. Start the PostgreSQL server using Docker Compose:

    ```shell
    docker-compose up -d
    ```

### Starting the Web Services

To start the web services, navigate to the `cmd/patient` directory and run the Go program:

```shell
cd cmd/patient
go run main.go

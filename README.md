# Book Keeper Computer Management System Called Book Keeper

This project provides a REST API for managing computers issued by a Sample Company. It allows the system administrator to store and retrieve details of computers, assign them to employees, and get notifications when an employee has more than three devices.

## Table of Contents

- [Getting Started](#getting-started)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Notes](#notes)

## Getting Started

Follow these instructions to set up and run the project on your local machine.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/majorchork/book-keeper.git
    cd book-keeper
    ```

2. Install Go dependencies:

    ```go
    go mod tidy
    ```

3. Install PostgreSQL and set up a database:

    ```sh
    sudo apt-get install postgresql postgresql-contrib
    sudo -u postgres createuser --interactive
    sudo -u postgres createdb book-keeper
    ```

4. Install Swaggo for generating Swagger documentation:

    ```sh
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

## Configuration

1. Create a `.env` file in the project root directory and add the following environment variables:

    ```env
    DATABASE_URL="postgres://youruser:yourpassword@localhost:5432/book-keeper"
    PORT=8081
    ```

    Replace `youruser` and `yourpassword` with your PostgreSQL user and password.

## Usage

1. Run the server:

    ```sh
    make run
    ```

    The server will start on `http://localhost:8081`.

## Endpoints

- `POST /computers/create`: Add a new computer
- `GET /computers/viewAll`: Get all computers
- `GET /computers/employee/:abbr`: Get all computers assigned to an employee
- `GET /computers/viewInfo/:id`: Get data of a single computer
- `DELETE /computers/:id`: Delete a computer
- `PUT /computers/assign`: Assign(Unassign) a computer to(from) an employee

## Given more time

1. i'd have added unit and integration tests

2. implemented a middleware to ensure admin permissions 

3. implemented a basic auth middleware to validate authorised users

4. ensure proper api documentation with swagger

## Notes

- Ensure PostgreSQL is running and properly configured before starting the application.
- Set the `DATABASE_URL` environment variable correctly to avoid connection issues.
- Make sure to run `swag init` whenever you update the API documentation comments in the source code.
- To test the endpoints, use tools like Postman or cURL.

## Amendments

- **Database Configuration**: If you change the database configuration, update the `DATABASE_URL` in the `.env` file accordingly.
- **Port Configuration**: The server runs on port 8080 by default. You can change this by setting the `PORT` environment variable.
- **Swagger Documentation**: Keep the Swagger comments up-to-date with the latest changes in the API.

## Conclusion

This README provides detailed instructions for setting up, running, and documenting the Book Keeper Computer Management System. Following these steps ensures that the project is properly configured and that the documentation remains up-to-date.

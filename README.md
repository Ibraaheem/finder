## Finder 

Finder is a simple Go application that provides directory listing functionality via a RESTful API.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)

## Installation

### Prerequisites
- Go (at least Go 1.19)
- Docker 

### Clone the Repository
The repositoy can be found at https://gitlab.com/ibraaheem/finder.git

### Build and Run Locally
To build and run the application please make use of the Makefile

- `make build` builds a Docker image

- `make run` runs the Docker image
    
The application will be accessible at `http://localhost:8080`.

- `make stop` stops and purges the image

### Using Docker

You can also manually run the commands to build and run the application in a Docker container:

- `docker build -t finder .`
- `docker run -p 8080:8080 finder` 

## Usage

Once the application is running and the contents are copied onto the pod, please make use of Postman or similar to interact with the API.

Below are the available API endpoints: 

## API Endpoints

### Health Check

-   **Endpoint**: `/health`
-   **Method**: GET
-   **Description**: Checks the health status of the application.
-   **Response**: Returns a JSON response indicating the health status.

Example Request:

`GET http://localhost:8080/health` 

### List Directory

-   **Endpoint**: `/explore`
-   **Method**: GET
-   **Parameters**:
    -   `path` (string, required): The path of the directory you want to list.
    -   `page` (int, optional): The page number for pagination (default is 1).
    -   `limit` (int, optional): The maximum number of entries per page (default is 10).
-   **Description**: Lists the contents of the specified directory.
-   **Response**: Returns a JSON response containing the directory entries.

Example Request:

`GET http://localhost:8080/explore?path=your/directory/path` 

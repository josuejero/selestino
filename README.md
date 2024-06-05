# Selestino

Selestino is a recipe website where users can input available ingredients and get relevant recipes for Peruvian dishes. This project leverages a variety of technologies including Golang, PostgreSQL, Docker, Kubernetes, and Jenkins for CI/CD.

## Table of Contents

- [Project Overview](#project-overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Setup and Installation](#setup-and-installation)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Running the Application](#running-the-application)
  - [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [CI/CD Pipeline](#cicd-pipeline)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

Selestino is designed to help users find Peruvian recipes based on the ingredients they have on hand. Users can register, log in, and search for recipes. The application supports JWT-based authentication and includes CI/CD pipelines for continuous integration and deployment.

## Features

- User registration and authentication
- Search recipes by ingredients and other criteria
- RESTful API for interacting with the application
- CI/CD pipeline with Jenkins
- Dockerized application for easy deployment
- Kubernetes for container orchestration

## Technologies Used

- Golang
- PostgreSQL
- Docker
- Kubernetes
- Jenkins
- JWT (JSON Web Tokens)
- Gorilla Mux
- Testify (testing framework)
- SQLMock (mocking SQL for tests)

## Setup and Installation

### Prerequisites

Ensure you have the following installed:

- Golang 1.20 or later
- Docker
- Docker Compose
- Kubernetes (Minikube or any other local Kubernetes setup)
- Jenkins (with necessary plugins)
- Git

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/josuejero/selestino.git
    cd selestino
    ```

2. Set up the environment variables:

    Create a `.env` file in the root directory with the following content:

    ```env
    DB_HOST=postgres
    DB_PORT=5432
    DB_USER=selestino_user
    DB_PASSWORD=your_password
    DB_NAME=selestino
    JWT_SECRET_KEY=my_secret_key
    ```

3. Build and run the Docker containers:

    ```bash
    docker-compose up --build
    ```

4. Set up the Kubernetes environment:

    ```bash
    minikube start
    eval $(minikube -p minikube docker-env)
    docker build -t selestino:latest .
    kubectl apply -f k8s/postgres-deployment.yaml
    kubectl apply -f k8s/selestino-deployment.yaml
    ```

5. Access the application:

    ```bash
    minikube service selestino --url
    ```

## Usage

### Running the Application

After following the installation steps, you can access the application using the Minikube service URL. Use a tool like Postman or cURL to interact with the API endpoints.

### API Endpoints

- **Register User**
    ```http
    POST /register
    ```

    **Request Body:**
    ```json
    {
        "username": "exampleuser",
        "password": "password123"
    }
    ```

- **Login User**
    ```http
    POST /login
    ```

    **Request Body:**
    ```json
    {
        "username": "exampleuser",
        "password": "password123"
    }
    ```

- **Get All Recipes**
    ```http
    GET /recipes
    ```

- **Add Recipe**
    ```http
    POST /recipes
    ```

    **Request Body:**
    ```json
    {
        "name": "Lomo Saltado",
        "ingredients": "Beef, Onion, Tomato, Soy Sauce",
        "instructions": "Stir-fry ingredients and serve with rice"
    }
    ```

- **Search Recipes by Ingredients**
    ```http
    GET /recipes/search?ingredients=Beef,Tomato
    ```

## Testing

Run the tests to ensure all functionality is working correctly:

```bash
go test ./internal/repository
```

## CI/CD Pipeline

The CI/CD pipeline is set up using Jenkins. It automates the build, push, and deployment processes. The `Jenkinsfile` in the root directory defines the pipeline stages.

To configure the Jenkins pipeline:

1. Install Jenkins and necessary plugins.
2. Create credentials in Jenkins for Docker Hub and Kubernetes.
3. Create a new pipeline job and configure it to use the repository's `Jenkinsfile`.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

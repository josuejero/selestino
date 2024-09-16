# Selestino

Selestino is a recipe recommendation website that generates personalized Peruvian dish suggestions based on user input. It is built using Python and PostgreSQL, with a focus on scalability and automation, utilizing Google Cloud, Docker, Kubernetes, and Jenkins for Continuous Integration and Continuous Deployment (CI/CD). The website also integrates RESTful APIs and is thoroughly tested with Selenium for browser-based functional testing.

## Table of Contents
- [Features](#features)
- [Technologies](#technologies)
- [Setup](#setup)
- [Running the Application](#running-the-application)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features
- Generates personalized Peruvian dish recommendations based on user input.
- Deployed on Google Cloud for scalability and reliability.
- Uses Docker and Kubernetes for containerization and orchestration.
- Jenkins CI/CD pipeline for automated testing, building, and deployment.
- Integration of RESTful APIs for dynamic data interaction.
- Comprehensive browser-based functional testing with Selenium.
  
## Technologies
The Selestino project is built using the following technologies:
- **Python**: Backend logic and API integrations.
- **PostgreSQL**: Database management.
- **Docker**: Containerization of the application.
- **Kubernetes**: Orchestration for scaling and managing the containerized services.
- **Google Cloud Platform (GCP)**: Hosting the application and managing cloud infrastructure.
- **Jenkins**: CI/CD pipeline for automating testing and deployment.
- **Selenium**: Browser automation for testing.
  
## Setup

### Prerequisites
Ensure you have the following installed:
- [Python 3.12](https://www.python.org/downloads/)
- [Docker](https://docs.docker.com/get-docker/)
- [Kubernetes](https://kubernetes.io/docs/setup/)
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)
- [Jenkins](https://www.jenkins.io/download/)
- [PostgreSQL](https://www.postgresql.org/download/)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/josuejero/selestino.git
    cd selestino
    ```

2. Set up a virtual environment and install dependencies:
    ```bash
    python3 -m venv venv
    source venv/bin/activate
    pip install -r requirements.txt
    ```

3. Set up PostgreSQL:
    ```bash
    psql -U postgres -c "CREATE DATABASE selestino_db;"
    psql -U postgres -c "CREATE USER selestino_user WITH PASSWORD 'password';"
    psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE selestino_db TO selestino_user;"
    ```

4. Set environment variables for Django:
    ```bash
    export DJANGO_SECRET_KEY='your-secret-key'
    export DB_NAME='selestino_db'
    export DB_USER='selestino_user'
    export DB_PASSWORD='password'
    export DB_HOST='localhost'
    ```

## Running the Application

1. Apply database migrations:
    ```bash
    python manage.py migrate
    ```

2. Run the development server:
    ```bash
    python manage.py runserver
    ```

3. Access the app at [http://localhost:8000](http://localhost:8000).

## Docker and Kubernetes

### Building and Running with Docker
1. Build the Docker image:
    ```bash
    docker build -t selestino:latest .
    ```

2. Run the Docker container:
    ```bash
    docker run -d -p 8000:8000 selestino:latest
    ```

### Deploying to Kubernetes
1. Apply Kubernetes configurations:
    ```bash
    kubectl apply -f k8s/configmap.yaml
    kubectl apply -f k8s/secret.yaml
    kubectl apply -f k8s/deployment.yaml
    kubectl apply -f k8s/service.yaml
    ```

2. Verify the deployment:
    ```bash
    kubectl rollout status deployment/selestino-deployment
    ```



## Testing

### Running Unit Tests
To run the unit tests with coverage:
```bash
pytest --cov=selestino tests/
```

### Browser Testing with Selenium
The project includes comprehensive browser testing using Selenium. Make sure to have [Google Chrome](https://www.google.com/chrome/) and [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/downloads) installed.

To run the Selenium tests:
```bash
python manage.py test
```

## Contributing
Contributions are welcome! Please submit a pull request or open an issue for discussion.

### Steps to Contribute
1. Fork the repository.
2. Create a new feature branch:
    ```bash
    git checkout -b feature-branch
    ```
3. Commit your changes:
    ```bash
    git commit -m "Add new feature"
    ```
4. Push to your fork and submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
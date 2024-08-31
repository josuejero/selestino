pipeline {
    agent any

    environment {
        PATH = "/usr/bin:$PATH"
        // Environment variables
        DOCKER_IMAGE = "selestino-web"
        // Commented out credentials for testing purposes
        // REGISTRY_CREDENTIALS = credentials('docker')
        // GITHUB_CREDENTIALS = credentials('github')
        // GOOGLE_CLOUD_CREDENTIALS = credentials('GOOGLE_CLOUD_CREDENTIALS')
        DB_NAME = "selestino"
        DB_USER = "josuejero"
        // DB_PASSWORD = credentials('DB_PASSWORD')
        GOOGLE_PROJECT_ID = "selestino-434015"
        GOOGLE_COMPUTE_ZONE = "us-central1-a"
    }

    stages {
        stage('Check Environment Variables') {
            steps {
                script {
                    echo "Checking environment variables... [DEBUG-000]"
                    sh 'echo DOCKER_IMAGE: $DOCKER_IMAGE'
                    // Echo placeholders for credentials
                    echo "REGISTRY_CREDENTIALS: ****"
                    echo "GITHUB_CREDENTIALS: ****"
                    echo "GOOGLE_CLOUD_CREDENTIALS: ****"
                    sh 'echo DB_NAME: $DB_NAME'
                    sh 'echo DB_USER: $DB_USER'
                    echo "DB_PASSWORD: ****"
                    sh 'echo GOOGLE_PROJECT_ID: $GOOGLE_PROJECT_ID'
                    sh 'echo GOOGLE_COMPUTE_ZONE: $GOOGLE_COMPUTE_ZONE'
                }
            }
        }

        stage('Check Environment') {
            steps {
                script {
                    sh 'echo "User: $(whoami)" [DEBUG-001]'
                    sh 'echo "Current directory: $(pwd)" [DEBUG-002]'
                    sh 'echo "PATH: $PATH" [DEBUG-003]'
                    sh 'echo "Checking if Docker is in PATH: $(which docker || echo \'Docker not found in PATH\')" [DEBUG-004]'
                    sh 'ls -la /var/run/docker.sock'
                    sh 'ls -la /usr/bin/docker || echo "/usr/bin/docker not found" [DEBUG-006]'
                    sh 'echo $SHELL [DEBUG-007]'
                }
            }
        }

        stage('Check Docker Installation') {
            steps {
                script {
                    try {
                        sh 'docker --version'
                        echo "Docker is installed and accessible. [DEBUG-008]"
                    } catch (Exception e) {
                        echo "Docker is not installed or not accessible: ${e.message} [ERROR-108]"
                        error("Failed at stage: Check Docker Installation [ERROR-108]")
                    }
                }
            }
        }

        stage('Checkout Code') {
            steps {
                echo "Skipping code checkout... [DEBUG-010]"
                // Skipping actual Git checkout since credentials are not being used
            }
        }

        stage('Build Docker Image') {
            steps {
                echo "Starting Docker image build... [DEBUG-012]"
                script {
                    try {
                        dockerImage = docker.build("${DOCKER_IMAGE}:${env.BUILD_ID}", "selestino/")
                        echo "Docker image built successfully: ${DOCKER_IMAGE}:${env.BUILD_ID} [DEBUG-013]"
                    } catch (Exception e) {
                        echo "Error during Docker image build: ${e.message} [ERROR-102]"
                        error("Failed at stage: Build Docker Image [ERROR-102]")
                    }
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo "Starting tests... [DEBUG-014]"
                script {
                    try {
                        dockerImage.inside('-v /var/run/docker.sock:/var/run/docker.sock') {
                            sh 'pytest selestino/tests/'
                        }
                        echo "Tests completed successfully. [DEBUG-015]"
                    } catch (Exception e) {
                        echo "Error during testing: ${e.message} [ERROR-103]"
                        error("Failed at stage: Run Tests [ERROR-103]"
                    }
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                echo "Skipping Docker image push... [DEBUG-016]"
                // Skipping Docker push since registry credentials are not being used
            }
        }

        stage('Deploy to Google Cloud') {
            steps {
                echo "Skipping deployment to Google Cloud... [DEBUG-018]"
                // Skipping Google Cloud deployment since credentials are not being used
            }
        }
    }

    post {
        success {
            echo 'Build was successful! [DEBUG-020]'
        }
        failure {
            echo 'Build failed. Please check the logs. [ERROR-106]'
        }
    }
}

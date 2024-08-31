pipeline {
    agent any

    environment {
        PATH = "/usr/local/bin:$PATH"
        DB_NAME = "selestino"
        DB_USER = "josuejero"
        DB_PASSWORD = "peruano1"
        DB_HOST = "db"
        GOOGLE_PROJECT_ID = "selestino-434015"
        GOOGLE_COMPUTE_ZONE = "us-central1-a"
        GITHUB_CREDENTIALS = credentials('GITHUB_CREDENTIALS_ID')
    }

    stages {
        stage('Check Environment Variables') {
            steps {
                script {
                    echo "Checking environment variables... [DEBUG-000]"
                    sh 'echo DB_NAME: $DB_NAME'
                    sh 'echo DB_USER: $DB_USER'
                    sh 'echo DB_PASSWORD: $DB_PASSWORD'
                    sh 'echo DB_HOST: $DB_HOST'
                    sh 'echo GOOGLE_PROJECT_ID: $GOOGLE_PROJECT_ID'
                    sh 'echo GOOGLE_COMPUTE_ZONE: $GOOGLE_COMPUTE_ZONE'
                    sh 'echo GITHUB_CREDENTIALS: ****'
                }
            }
        }

        stage('Check Environment') {
            steps {
                script {
                    sh 'echo "User: $(whoami)" [DEBUG-001]'
                    sh 'echo "Current directory: $(pwd)" [DEBUG-002]'
                    sh 'echo "PATH: $PATH" [DEBUG-003]'
                }
            }
        }

        stage('Install Docker Compose') {
            steps {
                script {
                    echo "Installing Docker Compose... [DEBUG-004]"
                    sh '''
                        sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
                        sudo chmod +x /usr/local/bin/docker-compose
                    '''
                }
            }
        }

        stage('Checkout Code') {
            steps {
                script {
                    echo "Starting code checkout... [DEBUG-010]"
                    try {
                        git branch: 'master', credentialsId: "${GITHUB_CREDENTIALS}", url: 'https://github.com/josuejero/selestino.git'
                        echo "Code checkout completed successfully. [DEBUG-011]"
                    } catch (Exception e) {
                        echo "Error during code checkout: ${e.message} [ERROR-101]"
                        error("Failed at stage: Checkout Code [ERROR-101]")
                    }
                }
            }
        }

        stage('Build Docker Images') {
            steps {
                script {
                    echo "Listing files and folders in the current directory... [DEBUG-012]"
                    sh 'ls -la'
                    
                    echo "Checking for docker-compose.yml file and building Docker images... [DEBUG-013]"
                    try {
                        if (fileExists('docker-compose.yml')) {
                            echo "docker-compose.yml found in the current directory. Building Docker images... [DEBUG-014]"
                            sh 'docker-compose build'
                        } else if (fileExists('selestino/docker-compose.yml')) {
                            echo "docker-compose.yml found in selestino directory. Changing directory and building Docker images... [DEBUG-015]"
                            dir('selestino') {
                                sh 'ls -la'  // List files in the selestino directory for debugging
                                sh 'docker-compose build'
                            }
                        } else {
                            error("docker-compose.yml not found in either current or selestino directory [ERROR-104]")
                        }
                    } catch (Exception e) {
                        echo "Error during Docker image build: ${e.message} [ERROR-103]"
                        error("Failed at stage: Build Docker Images [ERROR-103]")
                    }
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    echo "Listing files and folders in the current directory... [DEBUG-016]"
                    sh 'ls -la'

                    echo "Checking for docker-compose.yml file and running tests... [DEBUG-017]"
                    try {
                        if (fileExists('docker-compose.yml')) {
                            echo "docker-compose.yml found in the current directory. Running tests... [DEBUG-018]"
                            sh 'docker-compose run --rm web pytest tests/'
                        } else if (fileExists('selestino/docker-compose.yml')) {
                            echo "docker-compose.yml found in selestino directory. Changing directory and running tests... [DEBUG-019]"
                            dir('selestino') {
                                sh 'ls -la'  // List files in the selestino directory for debugging
                                sh 'docker-compose run --rm web pytest tests/'
                            }
                        } else {
                            error("docker-compose.yml not found in either current or selestino directory [ERROR-105]")
                        }
                    } catch (Exception e) {
                        echo "Error during testing: ${e.message} [ERROR-104]"
                        error("Failed at stage: Run Tests [ERROR-104]")
                    }
                }
            }
        }

        stage('Deploy to Google Cloud') {
            steps {
                script {
                    echo "Skipping deployment to Google Cloud... [DEBUG-020]"
                    // Include your deployment logic here when ready
                }
            }
        }
    }

    post {
        success {
            echo 'Build was successful! [DEBUG-022]'
        }
        failure {
            echo 'Build failed. Please check the logs. [ERROR-106]'
        }
    }
}

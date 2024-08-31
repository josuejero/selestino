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

        stage('Install Python 3 and Setup Virtual Environment') {
            steps {
                script {
                    echo "Checking for Python 3 and setting up virtual environment if necessary... [DEBUG-004]"
                    sh '''
                        if ! command -v python3 &> /dev/null
                        then
                            echo "Python 3 not found, installing Python 3... [DEBUG-005]"
                            sudo apt-get update -y
                            sudo apt-get install -y python3 python3-pip python3-venv
                        else
                            echo "Python 3 is already installed. [DEBUG-006]"
                        fi

                        python3 -m venv venv
                        source venv/bin/activate
                    '''
                }
            }
        }

        stage('Install Python Dependencies') {
            steps {
                script {
                    echo "Installing Python dependencies, including pytest... [DEBUG-007]"
                    sh '''
                        source venv/bin/activate
                        pip install --upgrade pip
                        pip install pytest
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

        stage('Run Tests') {
            steps {
                script {
                    echo "Running tests... [DEBUG-016]"
                    try {
                        sh '''
                            source venv/bin/activate
                            pytest tests/
                        '''
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

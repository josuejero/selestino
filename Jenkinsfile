pipeline {
    agent any

    environment {
        DATABASE_URL = "postgres://josuejero:peruano1@localhost:5432/selestino"
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    echo 'Cleaning the workspace...'
                    cleanWs()
                    echo 'Workspace cleaned!'
                    echo 'Checking out the code from GitHub...'
                    git branch: 'master', url: 'https://github.com/josuejero/selestino.git'
                    echo 'Checked out code successfully!'
                }
            }
        }

        stage('Setup') {
            steps {
                script {
                    sh '''
                        echo "Creating virtual environment..."
                        python3 -m venv env
                        source env/bin/activate
                        echo "Upgrading pip..."
                        pip install --upgrade pip
                        pwd
                        echo "Installing dependencies from requirements.txt..."
                        pip install -r requirements.txt
                        echo "Setup completed!"
                    '''
                }
            }
        }

        stage('Migrate Database') {
            steps {
                script {
                    sh '''
                        echo "Navigating to the Django project directory..."
                        cd selestino
                        echo "Activating virtual environment..."
                        source ../env/bin/activate
                        echo "Running database migrations..."
                        python manage.py migrate
                        echo "Migrations completed!"
                    '''
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    sh '''
                        echo "Navigating to the Django project directory..."
                        cd selestino
                        echo "Activating virtual environment..."
                        source ../env/bin/activate
                        echo "Running Django tests..."
                        python manage.py test
                        echo "Tests completed!"
                    '''
                }
            }
        }

        stage('Build') {
            steps {
                echo 'Build stage (currently no specific build actions configured)'
            }
        }

        stage('Deploy') {
            when {
                branch 'master'
            }
            steps {
                echo 'Deploy stage (currently no specific deploy actions configured)'
            }
        }
    }

    post {
        always {
            script {
                node {
                    echo 'Cleaning workspace after job completion...'
                    cleanWs()
                }
            }
        }
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed. Check the logs for details.'
        }
    }
}

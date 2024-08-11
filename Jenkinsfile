pipeline {
    agent any

    stages {
        stage('Setup') {
            steps {
                script {
                    // Clean the workspace
                    cleanWs()
                }

                // Debugging: Show the current directory and its contents
                sh '''
                    echo "Current Directory:"
                    pwd
                    echo "Contents of the Directory:"
                    ls -la
                '''

                // Install dependencies in a virtual environment
                sh '''
                    python3 -m venv env
                    source env/bin/activate
                    echo "Looking for requirements.txt:"
                    ls -la
                    pip install --upgrade pip
                '''
            }
        }

        stage('Code Quality Check') {
            steps {
                // Run a linter (e.g., flake8) to check code quality
                sh '''
                    source env/bin/activate
                    pip install flake8
                    flake8 selestino/ recipeservice/
                '''
            }
        }

        stage('Run Tests') {
            steps {
                // Run Django tests
                sh '''
                    source env/bin/activate
                    python manage.py test
                '''
            }
        }

        stage('Build') {
            steps {
                // You can add steps here if you plan to build the project, such as Dockerizing the app
                echo 'Build stage (optional)'
            }
        }

        stage('Deploy') {
            when {
                branch 'main'
            }
            steps {
                // Deploy the application (e.g., push Docker image to a registry, deploy to AWS, etc.)
                echo 'Deploy stage (optional)'
            }
        }
    }

    post {
        always {
            // Clean up the workspace after the pipeline is complete
            cleanWs()
        }
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}

pipeline {
    agent any

    environment {
        IMAGE_NAME = 'test'          // Replace with your Docker image name
        IMAGE_TAG = 'latest'                    // Replace with your tag if needed
        DOCKER_REGISTRY = 'your_registry_url'   // Replace with your Docker registry URL
        DOCKER_CREDENTIALS_ID = 'docker-credentials-id' // Replace with your Jenkins credentials ID for Docker registry
    }

    stages {
        stage('Clone Repository') {
            steps {
                checkout scm
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${IMAGE_NAME}:${IMAGE_TAG}")
                }
            }
        }
    }
}

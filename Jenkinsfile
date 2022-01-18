pipeline {
    agent any
    tools {
        go 'go-1.17'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Setting up dependencies'){
            agent {
                label "api2"
            }
            steps {
                echo 'Setting up dependencies'
                sh 'go mod vendor'
            }
        }
        stage('Build') {
            agent {
                label "api2"
            }
            steps {
                echo 'Running build automation'
                sh 'go build -o main'
            }
        }
    }
}
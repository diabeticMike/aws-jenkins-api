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
		steps {
			echo 'Setting up dependencies'
			sh 'go mod vendor'
		}
	}
        stage('Build') {
            steps {
                echo 'Running build automation'
                sh 'go build -o main'
            }
        }
    }
}

pipeline {
    agent any
    
    environment {
        // Enforce PowerShell Core as the default shell step execution
        JENKINS_POWERSHELL_COMMAND = 'pwsh'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Go Binary') {
            steps {
                powershell '''
                    # Go the module directory
                    cd ./example/lambda-func

                    # Set Linux target architecture for AWS Lambda
                    $Env:GOOS = "linux"
                    $Env:GOARCH = "amd64"
                    $Env:CGO_ENABLED="0"
                    
                    # Compile the binary
                    go build -tags lambda.norpc -o bootstrap main.go
                '''
            }
        }

        stage('Package Zip') {
            steps {
                powershell '''
                    # Package the binary into a zip deployment artifact
                    # Compress-Archive -Path bootstrap -DestinationPath main.zip -Force
                    $env:USERPROFILE
                    #& "$env:USERPROFILE/Go/bin/build-lambda-zip.exe" -o myFunction.zip bootstrap
                    
                    # Clean up the raw Linux binary from workspace
                    ls
                    #Remove-Item bootstrap
                '''
            }
        }
    }

    post {
        success {
            // Archive the zip file in Jenkins for record-keeping
            archiveArtifacts artifacts: 'myFunction.zip', fingerprint: true
        }
    }
}
pipeline{
    agent{
        docker {
            image 'beastapple/base-images:v1.0'
            args '--user root -v /var/run/docker.sock:/var/run/docker.sock' // mount Docker socket to access the host's Docker daemon
        }   
    }
    stages{
        stage("checkout"){
            steps{
                sh "echo checkout successful"
            }
            post{
                always{
                    echo "========always========"
                }
                success{
                    echo "========A executed successfully========"
                }
                failure{
                    echo "========A execution failed========"
                }
            }
        }
        stage("build"){
            steps{
                sh "ls -ltr"
                sh "docker build . -t integration"
            }
        }
    }
    post{
        always{
            echo "========always========"
        }
        success{
            echo "========pipeline executed successfully ========"
        }
        failure{
            echo "========pipeline execution failed========"
        }
    }
}






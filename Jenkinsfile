pipeline {
    agent any
    
    stages {
        stage('Clone') {
            steps {
                git 'https://github.com/CraftyCancer/BB-App.git'
            }
        }
        
        stage('Build') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'divya' passwordVariable: '123@1')]) {
                    script {
                        docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                            def customImage = docker.build("d1034/aliceclient", "--file Aliceclient/Dockerfile .")
                            customImage.push()
                            
                            def customImage_1 = docker.build("d1034/bobclient", "--file Bobclient/Dockerfile .")
                            customImage_1.push()
                            
                            def customImage_2 = docker.build("d1034/bbdb", "--file BBDB/Dockerfile .")
                            customImage_2.push()
                            
                            def customImage_3 = docker.build("d1034/bbserverclient", "--file BBServerclient/Dockerfile .")
                            customImage_2.push()
                        }
                    }
                }
            }
        }
        
        stage('Deploy') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'DOCKERHUB_USERNAME', passwordVariable: 'DOCKERHUB_PASSWORD')]) {
                    script {
                        docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                            def customImage = docker.image('d1034/aliceclient')
                            customImage.run('-p 8095:80 -d')
                            
                            def customImage_1 = docker.image('d1034/bbserverclient')
                            customImage_1.run()
                            
                            def customImage_2 = docker.image('d1034/bobclient')
                            customImage_2.run()
                            
                            def customImage_3 = docker.image('d1034/bbdb')
                            customImage_3.run()
                        }
                    }
                }
            }
        }
    }
    
    post {
        always {
            sh 'docker logout'
        }
    }
}

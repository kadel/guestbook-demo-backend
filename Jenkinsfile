node {
  stage ('Setup') {
    sh 'env'
    sh 'ls -lah'
    sh 'pwd'
  }
  stage('Build') {
    sh 'go build github.com/kadel/guestbook-demo-backend'
  }
  stage('Unit Test') {
    sh 'go test github.com/kadel/guestbook-demo-backend'
  }
}
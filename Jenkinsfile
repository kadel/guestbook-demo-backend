node {
  // stage('Checkout') {
  //   checkout scm
  // }
  stage('Build') {
    sh 'go build github.com/kadel/guestbook-demo-backend'
  }
  stage('Unit Test') {
    sh 'go test github.com/kadel/guestbook-demo-backend'
  }
}
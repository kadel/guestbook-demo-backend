node ('golang') {
  stage ('Setup') {
    sh 'go version'
    sh 'go get github.com/kadel/guestbook-demo-backend'
  }
  stage('Build') {
    sh 'go build github.com/kadel/guestbook-demo-backend'
  }
  stage('Unit Test') {
    sh 'go test github.com/kadel/guestbook-demo-backend'
  }
}
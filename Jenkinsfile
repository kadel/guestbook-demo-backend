node ('golang') {
  stage ('Setup') {
    sh 'ls -lah'
    sh 'pwd'
    sh 'env'
    sh 'ls -lah /root/go/bin/'
    sh 'go version'
  }
  stage('Build') {
    sh 'go build github.com/kadel/guestbook-demo-backend'
  }
  stage('Unit Test') {
    sh 'go test github.com/kadel/guestbook-demo-backend'
  }
}
node ('maven') {
  stage ('Setup') {
    sh 'ls -lah'
    sh 'ls -lah guestbook-demo-backend'
    sh 'git clone https://github.com/kadel/guestbook-demo-backend'
    sh 'curl -L https://github.com/kedgeproject/kedge/releases/download/v0.1.0/kedge-linux-amd64 -o kedge'
    sh 'chmod +x kedge'
  }

  state('Run Kedge') {
    sh './kedge generate -f Kedge/'
  }
}
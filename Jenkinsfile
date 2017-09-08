node ('maven') {
  stage ('Setup') {
    sh 'ls -lah'
    sh 'git clone https://github.com/kadel/guestbook-demo-backend'
    sh 'curl -L https://github.com/kedgeproject/kedge/releases/download/v0.1.0/kedge-linux-amd64 -o kedge'
    sh 'chmod +x kedge'
  }

  stage('Run Kedge') {
    sh './kedge generate -f guestbook-demo-backend/Kedge/'
    sh 'oc status'
  }
}
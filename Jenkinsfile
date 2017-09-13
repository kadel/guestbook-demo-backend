node ('maven') {
  stage ('Setup') {
    sh 'git clone https://github.com/kadel/guestbook-demo-backend'
    sh 'curl -L https://github.com/kedgeproject/kedge/releases/download/v0.1.0/kedge-linux-amd64 -o kedge'
    sh 'chmod +x kedge'
  }

  stage('Build') {
    sh 'oc apply -f guestbook-demo-backend/OpenShift/s2i-build.yaml'
    openshiftBuild(buildConfig: 'guestbook-backend', showBuildLogs: 'true')
  }

  stage('Deploy') {
    sh './kedge generate -f guestbook-demo-backend/Kedge/ | oc apply -f -'
    // hack to trigger new deployment (https://github.com/kubernetes/kubernetes/issues/27081)
    Date date = new Date()
    String timestamp = date.format("HH:mm:ss.dd/MM/yyyy")
    String patch = '{\"spec\":{\"template\":{\"metadata\":{\"annotations\":{\"date\":\"' + timestamp + '\"}}}}}'
    print patch
    sh 'kubectl patch deployment/backend -p ' + patch
  }
}
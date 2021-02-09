def repository = getJobName()
def registry = getLocalRegistry()
def namespace = 'engineering-portal'

node(['buildah', 'helm']) {
  stage('Checkout') {
    checkout scm
  }

  def imageList
  stage('Build') {
    imageList = buildah(repository: repository)
  }

  generateDefaultHelmChart(repository, [image: pickBestImageTag(imageList)])
  
  if (env.BRANCH_NAME == 'develop') {
    stage('Deploy') {
      deployWithStandardChart(repository, namespace)
    }
  }
}

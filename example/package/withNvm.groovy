def call(final String version = 'node', final Closure body) {
  sh "running with nvm: ${node}"

  body()
}

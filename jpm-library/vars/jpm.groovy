def install() {
  println "Installing jenkins scripts..."
}

def load(final String name) {

}

def call(final String name) {
  load "./jpm_modules/${name}"
}

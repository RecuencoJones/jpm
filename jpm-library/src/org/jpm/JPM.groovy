class JPM {
  private final steps

  JPM(steps) {
    this.steps = steps
  }

  void install() {
    steps.echo "Installing jenkins scripts..."

    steps.sh(
      script: """
      cat jpm.txt | jpm install
      """,
      returnStdout: true
    )

    steps.echo "Done!"
  }

  def load(final String name) {
    return steps.load "./jpm_modules/${name}"
  }
}

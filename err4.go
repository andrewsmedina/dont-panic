err := readfile(“.bashrc”)
if strings.Contains(err.Error(), "not found") {
  // handle error
}

err := readfile(“.bashrc”)
if strings.Contains(error.Error(), "not found") {
  // handle error
}

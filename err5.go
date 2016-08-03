	var ErrNotFound = stderr.New("Not Found")

	err := readfile(".bashrc")
	if err == ErrNotFound {
		// handle error
	}

type ErrNotFound struct {
  Message string
  Err error
}

func (e *ErrNotFound) Error() string {}

err := readfile(".bashrc")
if err, ok := err.(ErrNotFound); ok {
  // handle error
}

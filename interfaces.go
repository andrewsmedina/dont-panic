type ErrNotFound interface {
	Error
	NotFound()
}

func isNotFound(err error) bool {
	return err.(ErrNotFound)
}

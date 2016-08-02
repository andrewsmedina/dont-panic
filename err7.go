type notFound interface {
	NotFound()
}

func IsNotFound(err error) bool {
	_, ok := err.(notFound)
	return ok
}

package repository

// TODO Use NotFoundError to handle errors in the repository layer and propagate them to the api layer
type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{
		message: message,
	}
}

func (nf NotFoundError) Error() string {
	return nf.message
}

type UniqueIndexError struct {
	message string
}

func NewUniqueIndexError(message string) UniqueIndexError {
	return UniqueIndexError{
		message: message,
	}
}

func (ui UniqueIndexError) Error() string {
	return ui.message
}

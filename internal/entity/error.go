package entity

type ErrValidation struct {
	Message string
	Errors  []map[string]string
	Err     error
}

func NewErrValidation(message string, errors []map[string]string, err error) *ErrValidation {
	return &ErrValidation{
		Message: message,
		Errors:  errors,
		Err:     err,
	}
}

func (ve ErrValidation) Error() string {
	return ve.Err.Error()
}

type ErrNotFound struct {
	Message string
	Err     error
}

func NewErrNotFound(message string, err error) *ErrNotFound {
	return &ErrNotFound{
		Message: message,
		Err:     err,
	}
}

func (ent ErrNotFound) Error() string {
	return ent.Err.Error()
}

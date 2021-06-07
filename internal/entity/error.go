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
	if ve.Err != nil {
		return ve.Err.Error()
	}

	return ve.Message
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
	if ent.Err != nil {
		return ent.Err.Error()
	}

	return ent.Message
}

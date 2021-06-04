package entity

type ErrValidation struct {
	Message string
	Errors  []map[string]string
	Err     string
}

func (ve ErrValidation) Error() string {
	return ve.Err
}

type ErrNotFound struct {
	Message string
	Err     string
}

func (ent ErrNotFound) Error() string {
	return ent.Err
}

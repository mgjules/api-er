package repository

import "errors"

// Repository generic errors
var (
	ErrCreateRecord   = errors.New("can't create record")
	ErrGetRecord      = errors.New("can't retrieve record")
	ErrListRecords    = errors.New("can't retrieve records")
	ErrUpdateRecord   = errors.New("can't update record")
	ErrDeleteRecord   = errors.New("can't delete record")
	ErrRecordNotFound = errors.New("record not found")
)

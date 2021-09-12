package project

import "errors"

var (
	ErrUUID = errors.New("failed to create uuid")
	ErrSave = errors.New("failed to save item")
)

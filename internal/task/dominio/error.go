package dominio

import (
	"errors"
)

var (
	ErrInvalidTaskID          = errors.New("invalid Task ID")
	ErrInvalidTaskName        = errors.New("invalid Task Name")
	ErrInvalidTaskDescription = errors.New("invalid Task Description")
	ErrInvalidTaskStatus      = errors.New("invalid Task Status")
	ErrInvalidTaskDateStart   = errors.New("invalid Task Date Start")
	ErrInvalidTaskDateFinally = errors.New("invalid Task Date Finally")
	ErrInvalidTaskPriority    = errors.New("invalid Task Priority")
)

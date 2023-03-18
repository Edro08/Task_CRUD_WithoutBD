package dominio

import "strconv"

type ID struct {
	value int
}

func NewTaskID(value string) (ID, error) {

	v, err := strconv.Atoi(value)

	if err != nil {
		return ID{}, ErrInvalidTaskID
	}

	return ID{value: v}, nil
}

type Name struct {
	value string
}

func NewtaskName(value string) (Name, error) {
	if value == "" {
		return Name{}, ErrInvalidTaskName
	}

	return Name{value: value}, nil
}

type Description struct {
	value string
}

func NewTaskDescription(value string) (Description, error) {
	if value == "" {
		return Description{}, ErrInvalidTaskDescription
	}

	return Description{value: value}, nil
}

type Status struct {
	value int
}

func NewTaskStatus(value string) (Status, error) {

	v, err := strconv.Atoi(value)

	if err != nil {
		return Status{}, ErrInvalidTaskStatus
	}

	return Status{value: v}, nil
}

type DateStart struct {
	value string
}

func NewTaskDateStart(value string) (DateStart, error) {
	if value == "" {
		return DateStart{}, ErrInvalidTaskDateStart
	}

	return DateStart{value: value}, nil
}

type DateFinally struct {
	value string
}

func NewTaskDateFinally(value string) (DateFinally, error) {
	if value == "" {
		return DateFinally{}, ErrInvalidTaskDateFinally
	}

	return DateFinally{value: value}, nil
}

type Priority struct {
	value int
}

func NewTaskPriority(value string) (Priority, error) {

	v, err := strconv.Atoi(value)

	if err != nil {
		return Priority{}, ErrInvalidTaskPriority
	}

	return Priority{value: v}, nil
}

type Task struct {
	Id          ID
	Name        Name
	Description Description
	Status      Status
	DateStart   DateStart
	DateFinally DateFinally
	Priority    Priority
}

func NewTask(id, name, description, status, dateStart, dateFinally, priority string) (Task, error) {
	Id, errId := NewTaskID(id)
	Name, errName := NewtaskName(name)
	Description, errDescription := NewTaskDescription(description)
	Status, errStatus := NewTaskStatus(status)
	DateStart, errDateStart := NewTaskDateStart(dateStart)
	DateFinally, errDateFinally := NewTaskDateFinally(dateFinally)
	Priority, errPriority := NewTaskPriority(priority)

	if errId != nil {
		return Task{}, errId
	} else if errName != nil {
		return Task{}, errName
	} else if errDescription != nil {
		return Task{}, errDescription
	} else if errStatus != nil {
		return Task{}, errStatus
	} else if errDateStart != nil {
		return Task{}, errDateStart
	} else if errDateFinally != nil {
		return Task{}, errDateFinally
	} else if errPriority != nil {
		return Task{}, errPriority
	}

	return Task{
		Id:          Id,
		Name:        Name,
		Description: Description,
		Status:      Status,
		DateStart:   DateStart,
		DateFinally: DateFinally,
		Priority:    Priority,
	}, nil
}

func (c Task) IDs() int {
	return c.Id.value
}

func (c Task) Prioritys() int {
	return c.Priority.value
}

func (c Task) PriorityName(id int) string {
	var value = ""
	switch c.Priority.value {
	case 1:
		value = "Baja"
	case 2:
		value = "Media"
	case 3:
		value = "Alta"
	}

	return value
}

func (c Task) Statuses() int {
	return c.Status.value
}

package add

import (
	"garcia-eduardo-0323/internal/task/dominio"
	"reflect"
	"testing"
)

func TestAddData_Add(t *testing.T) {
	type args struct {
		id          string
		nombre      string
		description string
		status      string
		dateStart   string
		dateFinally string
		priority    string
		tasks       []dominio.Task
	}

	var tasks []dominio.Task
	task, _ := dominio.NewTask("1", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	tasks = append(tasks, task)

	testArgs := args{
		id:          "2",
		nombre:      "Leer",
		description: "Leer un Libro",
		status:      "1",
		dateStart:   "2023-03-08",
		dateFinally: "2023-03-08",
		priority:    "3",
		tasks:       tasks}

	testArgsDuplicateID := args{
		id:    "1",
		tasks: tasks}

	testArgsInvalidID := args{
		id:    "MM",
		tasks: tasks}

	testArgsInvalidName := args{
		id:     "3",
		nombre: "",
		tasks:  tasks}

	task2, _ := dominio.NewTask("2", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	var wantTasks []dominio.Task
	wantTasks = append(wantTasks, task)
	wantTasks = append(wantTasks, task2)

	tests := []struct {
		name    string
		args    args
		want    []dominio.Task
		wantErr bool
	}{
		{
			name:    "Test Add Successful",
			args:    testArgs,
			want:    wantTasks,
			wantErr: false,
		},
		{
			name:    "Test Add Failure Invalid Id Task",
			args:    testArgsInvalidID,
			want:    tasks,
			wantErr: true,
		},
		{
			name:    "Test Add Failure Invalid Name",
			args:    testArgsInvalidName,
			want:    tasks,
			wantErr: true,
		},
		{
			name:    "Test Add Failure Duplicate ID",
			args:    testArgsDuplicateID,
			want:    tasks,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AddData{}
			got, err := p.Add(tt.args.id, tt.args.nombre, tt.args.description, tt.args.status, tt.args.dateStart, tt.args.dateFinally, tt.args.priority, tt.args.tasks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAddDataImpl(t *testing.T) {
	tests := []struct {
		name string
		want *AddData
	}{
		{
			name: "Test NewAddDataImpl",
			want: NewAddDataImpl(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAddDataImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAddDataImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

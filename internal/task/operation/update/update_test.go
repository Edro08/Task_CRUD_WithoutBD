package update

import (
	"garcia-eduardo-0323/internal/task/dominio"
	"reflect"
	"testing"
)

func TestNewUpdateDataImpl(t *testing.T) {
	tests := []struct {
		name string
		want *UpdateData
	}{
		{
			name: "Test NewUpdateDataImpl",
			want: NewUpdateDataImpl(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUpdateDataImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdateDataImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateData_Update(t *testing.T) {
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
	task, _ = dominio.NewTask("2", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	tasks = append(tasks, task)

	testArgs := args{
		id:          "1",
		nombre:      "Leer 2",
		description: "Leer un Libro 2",
		status:      "1",
		dateStart:   "2023-03-08",
		dateFinally: "2023-03-08",
		priority:    "3",
		tasks:       tasks}

	testArgsInvalidID := args{
		id:    "MM",
		tasks: tasks}

	testArgsInvalidName := args{
		id:     "1",
		nombre: "",
		tasks:  tasks}

	var wantTasks []dominio.Task
	task, _ = dominio.NewTask("1", "Leer 2", "Leer un Libro 2", "1", "2023-03-08", "2023-03-08", "3")
	wantTasks = append(wantTasks, task)
	task, _ = dominio.NewTask("2", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	wantTasks = append(wantTasks, task)

	tests := []struct {
		name    string
		args    args
		want    []dominio.Task
		wantErr bool
	}{
		{
			name:    "Test Update Successful",
			args:    testArgs,
			want:    wantTasks,
			wantErr: false,
		},
		{
			name:    "Test Update Failure Invalid Id Task",
			args:    testArgsInvalidID,
			want:    tasks,
			wantErr: true,
		},
		{
			name:    "Test Update Failure Invalid Name",
			args:    testArgsInvalidName,
			want:    tasks,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &UpdateData{}
			got, err := p.Update(tt.args.id, tt.args.nombre, tt.args.description, tt.args.status, tt.args.dateStart, tt.args.dateFinally, tt.args.priority, tt.args.tasks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

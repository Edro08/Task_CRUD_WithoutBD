package delete

import (
	"garcia-eduardo-0323/internal/task/dominio"
	"reflect"
	"testing"
)

func TestDeleteData_Delete(t *testing.T) {
	type args struct {
		id    string
		tasks []dominio.Task
	}

	var tasks []dominio.Task
	task, _ := dominio.NewTask("1", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	tasks = append(tasks, task)
	task2, _ := dominio.NewTask("2", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	tasks = append(tasks, task2)

	testArgs := args{
		id:    "1",
		tasks: tasks}

	testArgsInvalidID := args{
		id:    "MM",
		tasks: tasks}

	var wantTasks []dominio.Task
	wantTasks = append(wantTasks, task2)

	tests := []struct {
		name    string
		args    args
		want    []dominio.Task
		wantErr bool
	}{
		{
			name:    "Test Delete Successful",
			args:    testArgs,
			want:    wantTasks,
			wantErr: false,
		},
		{
			name:    "Test Delete Failure Invalid Id Task",
			args:    testArgsInvalidID,
			want:    tasks,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DeleteData{}
			got, err := p.Delete(tt.args.id, tt.args.tasks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDeleteDataImpl(t *testing.T) {
	tests := []struct {
		name string
		want *DeleteData
	}{
		{
			name: "Test NewDeleteDataImpl",
			want: NewDeleteDataImpl(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeleteDataImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeleteDataImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

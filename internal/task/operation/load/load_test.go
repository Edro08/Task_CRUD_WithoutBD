package load

import (
	"garcia-eduardo-0323/internal/task/dominio"
	"reflect"
	"testing"
)

func TestLoadData_Load(t *testing.T) {
	type args struct {
		tasks []dominio.Task
	}

	testArgs := args{
		tasks: []dominio.Task{},
	}

	var wantTasks []dominio.Task
	task, _ := dominio.NewTask("1", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	wantTasks = append(wantTasks, task)

	task, _ = dominio.NewTask("2", "Escuchar", "Escuchar Musica", "2", "2023-03-08", "2023-03-08", "2")
	wantTasks = append(wantTasks, task)

	task, _ = dominio.NewTask("3", "Trabajar", "Trabajar en la casa", "3", "2023-03-08", "2023-03-08", "1")
	wantTasks = append(wantTasks, task)

	tests := []struct {
		name    string
		args    args
		want    []dominio.Task
		wantErr bool
	}{
		{
			name:    "Test Load Successful",
			args:    testArgs,
			want:    wantTasks,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &LoadData{}
			got, err := p.Load(tt.args.tasks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLoadDataImpl(t *testing.T) {
	tests := []struct {
		name string
		want *LoadData
	}{
		{
			name: "Test NewLoadDataImpl",
			want: NewLoadDataImpl(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoadDataImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoadDataImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

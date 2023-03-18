package list

import (
	"garcia-eduardo-0323/internal/task/dominio"
	"reflect"
	"testing"
)

func TestListData_All(t *testing.T) {
	type args struct {
		tasks []dominio.Task
	}

	var tasks []dominio.Task
	task, _ := dominio.NewTask("1", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	tasks = append(tasks, task)
	task2, _ := dominio.NewTask("2", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	tasks = append(tasks, task2)

	testArgs := args{tasks: tasks}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test List All Successful",
			args:    testArgs,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ListData{}
			if err := p.All(tt.args.tasks); (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListData_List(t *testing.T) {
	type args struct {
		option string
		id     string
		tasks  []dominio.Task
	}
	var tasks []dominio.Task
	task, _ := dominio.NewTask("1", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	tasks = append(tasks, task)
	task2, _ := dominio.NewTask("2", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	tasks = append(tasks, task2)

	testArgsAll := args{option: "1", id: "", tasks: tasks}
	testArgsOrderBy := args{option: "2", id: "", tasks: tasks}
	testArgsOnly := args{option: "4", id: "1", tasks: tasks}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test List All Successful",
			args:    testArgsAll,
			wantErr: false,
		},
		{
			name:    "Test List OrderBy Successful",
			args:    testArgsOrderBy,
			wantErr: false,
		},
		{
			name:    "Test List onlyTask Successful",
			args:    testArgsOnly,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ListData{}
			if err := p.List(tt.args.option, tt.args.id, tt.args.tasks); (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListData_OrderByPriorityAndStatus(t *testing.T) {
	type args struct {
		tasks []dominio.Task
	}

	var tasks []dominio.Task
	task, _ := dominio.NewTask("1", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "1")
	tasks = append(tasks, task)
	task2, _ := dominio.NewTask("2", "Leer", "Leer un Libro", "2", "2023-03-08", "2023-03-08", "3")
	tasks = append(tasks, task2)

	testArgs := args{tasks: tasks}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test List OrderBy Successful",
			args:    testArgs,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ListData{}
			if err := p.OrderByPriorityAndStatus(tt.args.tasks); (err != nil) != tt.wantErr {
				t.Errorf("OrderByPriorityAndStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListData_onlyTask(t *testing.T) {
	type args struct {
		id    string
		tasks []dominio.Task
	}

	var tasks []dominio.Task
	task, _ := dominio.NewTask("1", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "1")
	tasks = append(tasks, task)
	task2, _ := dominio.NewTask("2", "Leer", "Leer un Libro", "2", "2023-03-08", "2023-03-08", "3")
	tasks = append(tasks, task2)

	testArgs := args{id: "1", tasks: tasks}
	testArgsErr := args{id: "MM", tasks: tasks}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test List onlyTask Successful",
			args:    testArgs,
			wantErr: false,
		},
		{
			name:    "Test List onlyTask Failure Invalid Id",
			args:    testArgsErr,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ListData{}
			if err := p.onlyTask(tt.args.id, tt.args.tasks); (err != nil) != tt.wantErr {
				t.Errorf("onlyTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewListDataImpl(t *testing.T) {
	tests := []struct {
		name string
		want *ListData
	}{
		{
			name: "Test NewListDataImpl",
			want: NewListDataImpl(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListDataImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListDataImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

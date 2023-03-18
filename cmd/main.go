package main

import (
	"fmt"
	"garcia-eduardo-0323/internal/task/dominio"
	"garcia-eduardo-0323/internal/task/operation/add"
	"garcia-eduardo-0323/internal/task/operation/delete"
	"garcia-eduardo-0323/internal/task/operation/list"
	"garcia-eduardo-0323/internal/task/operation/load"
	"garcia-eduardo-0323/internal/task/operation/update"
	"log"
)

func main() {

	// Creacion de variables y estructuras a usar
	var option string
	var optionList string
	var id string
	var tasks []dominio.Task

	// Inyeccion de dependencias
	ListData := list.NewListDataImpl()
	AddData := add.NewAddDataImpl()
	DeleteData := delete.NewDeleteDataImpl()
	LoadData := load.NewLoadDataImpl()
	UpdateData := update.NewUpdateDataImpl()

	// Carga de datos inicial arreglo
	tasks, err := LoadData.Load(tasks)
	if err != nil {
		log.Println("Error in load data: " + err.Error())
		return
	}

	//Bucle de menu
	for option != "5" {
		Menu()

		fmt.Println("Digitar Opcion:")
		_, err := fmt.Scanln(&option)
		if err != nil {
			log.Println("Error in Scanner: " + err.Error())
		}

		switch option {
		case "1":
			id, name, description, status, dateStart, dateFinally, priority := RetrieverTask("add")
			tasks, err = AddData.Add(id, name, description, status, dateStart, dateFinally, priority, tasks)
			if err != nil {
				log.Println("Error in add task: " + err.Error())
			}

		case "2":
			ListTaskMenu()

			fmt.Println("Digitar Opcion:")
			_, err := fmt.Scanln(&optionList)
			if err != nil {
				log.Println("Error in Scanner: " + err.Error())
			}

			err = ListData.List(optionList, "", tasks)
			if err != nil {
				log.Println("Error in list tasks: " + err.Error())
			}
		case "3":
			fmt.Println("Digitar Id:")
			_, err := fmt.Scanln(&id)
			if err != nil {
				log.Println("Error in Scanner: " + err.Error())
			}

			err = ListData.List(optionList, "4", tasks)
			if err != nil {
				log.Println("Error in list task: " + err.Error())
			} else {
				_, name, description, status, dateStart, dateFinally, priority := RetrieverTask("update")
				tasks, err = UpdateData.Update(id, name, description, status, dateStart, dateFinally, priority, tasks)
				if err != nil {
					log.Println("Error in update task: " + err.Error())
				}
			}

		case "4":
			fmt.Println("Digitar Id:")
			_, err := fmt.Scanln(&id)
			if err != nil {
				log.Println("Error in Scanner: " + err.Error())
			}
			tasks, err = DeleteData.Delete(id, tasks)
			if err != nil {
				log.Println("Error in delete task: " + err.Error())
			}

		case "5":
			fmt.Println("Adios!")
		}
	}

}

func Menu() {
	fmt.Println("Task Management: ")
	fmt.Println("1- Add Task")
	fmt.Println("2- List Task")
	fmt.Println("3- Update Task")
	fmt.Println("4- Delete Task")
	fmt.Println("5- Salir")
	fmt.Println("")
}

func ListTaskMenu() {
	fmt.Println("	List Task:")
	fmt.Println("	1- All")
	fmt.Println("	2- Order by Priority and Status Pendiente")
	fmt.Println("")
}

func RetrieverTask(flow string) (id, name, description, status, dateStart, dateFinally, priority string) {

	if flow == "add" {
		fmt.Println("Digitar Id:")
		_, err := fmt.Scanln(&id)
		if err != nil {
			return
		}
	}

	fmt.Println("Digitar Name:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}

	fmt.Println("Digitar Description:")
	_, err = fmt.Scanln(&description)
	if err != nil {
		return
	}

	fmt.Println("Digitar Status:")
	_, err = fmt.Scanln(&status)
	if err != nil {
		return
	}
	fmt.Println("Digitar Fecha Inicio:")
	_, err = fmt.Scanln(&dateStart)
	if err != nil {
		return
	}

	fmt.Println("Digitar Fecha FInalizacion:")
	_, err = fmt.Scanln(&dateFinally)
	if err != nil {
		return
	}

	fmt.Println("Digitar Prioridad:")
	_, err = fmt.Scanln(&priority)
	if err != nil {
		return
	}

	fmt.Println("")

	return id, name, description, status, dateStart, dateFinally, priority
}

package app

import (
	"encoding/json"
	"fmt"
	"os"

	todolist "github.com/JMitchell159/CLIToDo/internal/toDoList"
)

type ListCollection map[string]todolist.TaskList

func CreateCollection() ListCollection {
	return make(ListCollection)
}

func (listCol ListCollection) AddList(name string, list todolist.TaskList) {
	listCol[name] = list
}

func (listCol ListCollection) Save(listName string) error {
	jsonList, err := json.Marshal(listCol)
	if err != nil {
		return fmt.Errorf("something went wrong while marshalling map to json: %v", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("something went wrong while finding home directory: %v", err)
	}

	if _, err = os.Stat(home+"/toDoLists"); os.IsNotExist(err) {
		err = os.Mkdir(home+"/toDoLists", 0750)
		if err != nil {
			return fmt.Errorf("something went wrong while creating to do list directory: %v", err)
		}
	}

	fileName := fmt.Sprintf("%s/toDoLists/%s.json", home, listName)
	err = os.WriteFile(fileName, jsonList, 0660)
	if err != nil {
		return fmt.Errorf("something went wrong while writing to do list to file: %v", err)
	}

	return nil
}

func Load(listName string) (ListCollection, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("something went wrong while finding home directory: %v", err)
	}

	fileName := fmt.Sprintf("%s/toDoLists/%s.json", home, listName)
	jsonList, err := os.ReadFile(fileName)
	if os.IsNotExist(err) {
		return CreateCollection(), nil
	}
	if err != nil {
		return nil, fmt.Errorf("something went wrong while reading to do list: %v", err)
	}

	result := CreateCollection()
	err = json.Unmarshal(jsonList, &result)
	if err != nil {
		return nil, fmt.Errorf("something went wrong while unmarshalling json to map: %v", err)
	}

	return result, nil
}


package project

import (
	"encoding/json"
	"fmt"
	"golang-rest-api-portfolio/api/v1/models"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

type inMemoryProjectDataBase struct {
	projects []models.Project
}

var singleInstance *inMemoryProjectDataBase

func getInstance() *inMemoryProjectDataBase {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &inMemoryProjectDataBase{}
			singleInstance.projects = *jsonToProjects(fetchJson())
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func fetchJson() []byte {
	jsonFile, err := os.Open("./assets/json/projects.json")

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func jsonToProjects(data []byte) *[]models.Project {
	var project []models.Project
	err := json.Unmarshal(data, &project)

	if err != nil {
		log.Fatal(err)
	}

	return &project
}

func GetProjectByID(id string) *models.Project {
	var result *models.Project = nil

	var iterate []models.Project = *&getInstance().projects
	for i, a := range iterate {
		if a.ID == id {
			result = &iterate[i]
		}
	}

	return result
}

func GetProjects() *[]models.Project {
	return &getInstance().projects
}

func AddProject(project models.Project) {
	*&getInstance().projects = append(*&getInstance().projects, project)
}

func RemoveProject(id string) string {
	var deleted_id string = ""

	array := *&getInstance().projects

	for i, a := range array {
		if a.ID == id {
			deleted_id = array[i].ID
			array[i] = *&getInstance().projects[len(array)-1]
			array[len(array)-1] = models.Project{}
			getInstance().projects = array[:len(array)-1]
		}
	}

	return deleted_id
}

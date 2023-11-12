package experience

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

type inMemoryExperienceDataBase struct {
	experiences []models.Experience
}

var singleInstance *inMemoryExperienceDataBase

func getInstance() *inMemoryExperienceDataBase {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &inMemoryExperienceDataBase{}
			singleInstance.experiences = *jsonToExperiences(fetchJson())
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func fetchJson() []byte {
	jsonFile, err := os.Open("./assets/json/experiences.json")

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func jsonToExperiences(data []byte) *[]models.Experience {
	var experience []models.Experience
	err := json.Unmarshal(data, &experience)

	if err != nil {
		log.Fatal(err)
	}

	return &experience
}

func GetExperienceByID(id string) *models.Experience {
	var result *models.Experience = nil

	var iterate []models.Experience = *&getInstance().experiences
	for i, a := range iterate {
		if a.ID == id {
			result = &iterate[i]
		}
	}

	return result
}

func GetExperiences() *[]models.Experience {
	return &getInstance().experiences
}

func AddExperience(experience models.Experience) {
	*&getInstance().experiences = append(*&getInstance().experiences, experience)
}

func RemoveExperience(id string) string {
	var deleted_id string = ""

	array := *&getInstance().experiences

	for i, a := range array {
		if a.ID == id {
			deleted_id = array[i].ID
			array[i] = *&getInstance().experiences[len(array)-1]
			array[len(array)-1] = models.Experience{}
			getInstance().experiences = array[:len(array)-1]
		}
	}

	return deleted_id
}

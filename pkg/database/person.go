package database

import (
	"encoding/json"
	"fmt"
	"github.com/lruggieri/graphqltest/pkg/person"
	"io/ioutil"
	"os"
	"path"
)

var allPeople = func() []person.Person {
	jsonFile, err := os.Open(path.Join("pkg", "database", "people.json"))

	if err != nil {
		fmt.Printf("failed to open json file, error: %v", err)
	}

	jsonDataFromFile, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	var peopleData []person.Person

	err = json.Unmarshal(jsonDataFromFile, &peopleData)

	if err != nil {
		fmt.Printf("failed to parse json, error: %v", err)
	}

	return peopleData
}()

func GetAllPeople() []person.Person {
	return allPeople
}

func GetPersonByID(id int) *person.Person {
	for i := range allPeople {
		if allPeople[i].ID == id {
			return &allPeople[i]
		}
	}
	return nil
}

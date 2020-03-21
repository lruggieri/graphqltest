package database

import (
	"encoding/json"
	"fmt"
	"github.com/lruggieri/graphqltest/pkg/company"
	"io/ioutil"
	"os"
	"path"
)

var allCompanies = func() []company.Company {
	jsonFile, err := os.Open(path.Join("pkg", "database", "companies.json"))

	if err != nil {
		fmt.Printf("failed to open json file, error: %v", err)
	}

	jsonDataFromFile, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	var companiesData []company.Company

	err = json.Unmarshal(jsonDataFromFile, &companiesData)

	if err != nil {
		fmt.Printf("failed to parse json, error: %v", err)
	}

	return companiesData
}()

func GetAllCompanies() []company.Company {
	return allCompanies
}

func GetCompanyByID(id int) company.Company {
	for i := range allCompanies {
		if allCompanies[i].ID == id {
			return allCompanies[i]
		}
	}
	return company.Company{}
}

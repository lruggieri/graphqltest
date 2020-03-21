package database

import (
	"encoding/json"
	"fmt"
	"github.com/lruggieri/graphqltest/pkg/job"
	"io/ioutil"
	"os"
	"path"
)

var allJobs = func()[]job.Job{
	jsonFile, err := os.Open(path.Join("pkg","database","jobs.json"))

	if err != nil {
		fmt.Printf("failed to open json file, error: %v", err)
	}

	jsonDataFromFile, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	var jobsData []job.Job

	err = json.Unmarshal(jsonDataFromFile, &jobsData)

	if err != nil {
		fmt.Printf("failed to parse json, error: %v", err)
	}

	return jobsData
}()

func GetAllJobs() []job.Job{
	return allJobs
}

func GetJobByID(id int) *job.Job{
	for i := range allJobs{
		if allJobs[i].ID == id{
			return &allJobs[i]
		}
	}
	return nil
}
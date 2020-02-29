package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const batchSize = 20
const preferRegion = "England"

var pool []string

// use https://uinames.com/api/
func GetName() string {
	if len(pool) <= 0 {
		names, err := fetchNames()
		if err != nil {
			fmt.Println(err)
			pool = append(pool, "can", "not", "fetch", "names", "use", "preset", "one")
		}
		pool = append(pool, names...)
	}
	return deQueue()
}

func fetchNames() ([]string, error) {
	var names []string
	url := fmt.Sprintf("https://uinames.com/api/?amount=%d&region=%s", batchSize, preferRegion)
	fmt.Printf("fetch names from %s\n", url)
	response, err := http.Get(url)
	defer response.Body.Close()
	if err == nil {
		var generatedNames []generatedName
		err = json.NewDecoder(response.Body).Decode(&generatedNames)
		if err == nil {
			for _, item := range generatedNames {
				names = append(names, fmt.Sprintf("%s-%s", item.Name, item.Surname))
			}
		}
	}
	return names, err
}

func deQueue() string {
	result := pool[0]
	pool = pool[1:]
	return result
}

type generatedName struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

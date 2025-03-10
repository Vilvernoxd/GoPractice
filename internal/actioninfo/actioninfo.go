package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(datastrings string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {

	for _, activity := range dataset {
		err := dp.Parse(activity)
		if err != nil {
			fmt.Printf("Parsing error: %v\n", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Failed to get activity information: %v\n", err)
			continue
		}

		fmt.Println(info)
	}
}

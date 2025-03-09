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
			fmt.Printf("Ошибка парсинга: %v\n", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Не удалось получить информацию об активности: %v\n", err)
			continue
		}

		fmt.Println(info)
	}
}

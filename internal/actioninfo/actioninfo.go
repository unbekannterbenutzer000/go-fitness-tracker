package actioninfo

import (
	"fmt"
	"log"
)

// DataParser определяет методы разбора данных и получения информации об активности.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info обрабатывает записи об активности, выводит результаты и регистрирует ошибки.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Print(info)
	}
}

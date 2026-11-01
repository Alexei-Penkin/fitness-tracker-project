package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает типы, умеющие разбирать строку с данными
// и формировать строку с информацией об активности.
// Методы интерфейса реализованы в пакетах trainings и daysteps.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info перебирает строки набора данных, парсит каждую и выводит
// сформированную информацию об активности. Ошибки парсинга и
// формирования информации логируются, обработка набора продолжается.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка парсинга данных %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации: %v", err)
			continue
		}

		fmt.Println(info)
	}
}

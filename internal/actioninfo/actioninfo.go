package actioninfo

import "fmt"

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, v := range dataset {

		err := dp.Parse(v)
		if err != nil {
			fmt.Println("ошибка:", err)
			continue
		}

		result, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("ошибка:", err)
			continue
		}

		fmt.Println(result)
	}
}

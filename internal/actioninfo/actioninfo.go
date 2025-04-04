// Пакет actioninfo реализует вывод общей информации обо всех тренировках и прогулках.
package actioninfo

import "fmt"

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info принимает слайс строк с данными о тренировках или прогулках и экземпляр одной структур и выводит итоговую информацию о проведенной активности.
func Info(dataset []string, dp DataParser) {
	for i := range dataset {
		err := dp.Parse(dataset[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		out, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out)
	}
}

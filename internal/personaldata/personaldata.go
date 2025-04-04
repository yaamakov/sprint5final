// Пакет persondata хранит персональные данные о человеке.
package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print метод выводит информацию о персоне.
func (p Personal) Print() {
	out := fmt.Sprintf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
	fmt.Print(out)
}

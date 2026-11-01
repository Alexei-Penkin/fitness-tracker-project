package personaldata

import "fmt"

// Personal содержит персональные данные пользователя.
// Структура встраивается в другие структуры проекта.
type Personal struct {
	Name   string  // имя пользователя
	Weight float64 // вес пользователя в килограммах
	Height float64 // рост пользователя в метрах
}

// Print выводит данные пользователя на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}

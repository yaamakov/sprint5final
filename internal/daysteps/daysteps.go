// Пакет daysteps реализует функционал для парсинга строки с данными о прогулках и формирования строки с информацией о них.
package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse парсит строку с данными и записывает данные в соответствующие поля структуры.
func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("step count convertation failed: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("negative or equal 0 number of steps")
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("time convertation failed: %w", err)
	}

	ds.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает строку с данными о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	var distance float64
	if ds.Duration <= 0 {
		return "", fmt.Errorf("duration must be greater than 0")
	}

	distance = spentenergy.Distance(ds.Steps)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Колличество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, distance, calories), nil

}

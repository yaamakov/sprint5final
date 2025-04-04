// Пакет trainings реализetn функционал по разбору строки с данными о тренировках и формирования строки с информацией о них.
package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse парсит строку с данными и записывает данные в соответствующие поля структуры.
func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("step count convertation failed: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("negative number of steps")
	}

	t.Steps = steps

	switch parts[1] {
	case "Бег":
		t.TrainingType = parts[1]
	case "Ходьба":
		t.TrainingType = parts[1]
	default:
		return fmt.Errorf("unknown training type")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("time convertation failed: %w", err)
	}

	t.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает строку с данными о тренировке.
func (t Training) ActionInfo() (string, error) {
	var distance, speed float64
	if t.Duration <= 0 {
		return "", fmt.Errorf("duration must be greater than 0")
	}

	distance = spentenergy.Distance(t.Steps)

	speed = spentenergy.MeanSpeed(t.Steps, t.Duration)

	switch t.TrainingType {
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nCожгли калорий %.2f ккал.", t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
	case "Ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nCожгли калорий %.2f ккал.", t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
	default:
		return "", fmt.Errorf("unknown training type")
	}
}

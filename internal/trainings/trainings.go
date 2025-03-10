package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {

	components := strings.Split(datastring, ",")
	if len(components) != 3 {
		err = errors.New("expected 3 parts in the string")
		return
	}

	steps, stepsErr := strconv.Atoi(components[0])
	if stepsErr != nil {
		err = errors.New("step count must be an integer")
		return
	}
	t.Steps = steps

	trainingType := components[1]
	if trainingType != "Бег" && trainingType != "Ходьба" {
		err = errors.New("invalid training type")
		return
	}
	t.TrainingType = trainingType

	duration, durationErr := time.ParseDuration(components[2])
	if durationErr != nil {
		err = errors.New("invalid duration format")
		return
	}
	t.Duration = duration

	return
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {

	if t.Duration <= 0 {
		return "", errors.New("duration must be greater than 0")
	}

	dist := spentenergy.Distance(t.Steps)

	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "неизвестный тип тренировки", errors.New("unknown training type")
	}

	if err != nil {
		return "", err
	}

	res := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), dist, avgSpeed, calories)

	return res, nil
}

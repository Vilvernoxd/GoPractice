package daysteps

import (
	"errors"
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

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {

	components := strings.Split(datastring, ",")
	if len(components) != 2 {
		err = errors.New("invalid string format")
		return
	}

	steps, err := strconv.Atoi(components[0])
	if err != nil {
		err = errors.New("step count must be an integer")
		return
	}

	duration, err := time.ParseDuration(components[1])
	if err != nil {
		err = errors.New("invalid duration format")
		return
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {

	if ds.Duration <= 0 {
		return "", errors.New("duration must be greater than 0")
	}

	if ds.Weight <= 0 {
		return "", errors.New("weight must be greater than 0")
	}

	dist := spentenergy.Distance(ds.Steps)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", errors.New("error calculating calories")
	}

	res := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, dist, calories)

	return res, nil
}

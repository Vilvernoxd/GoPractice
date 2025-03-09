package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
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
		err = errors.New("неверный формат строки")
		return
	}

	steps, err := strconv.Atoi(components[0])
	if err != nil {
		err = errors.New("количество шагов должно быть целым числом")
		return
	}

	duration, err := time.ParseDuration(components[1])
	if err != nil {
		err = errors.New("неверный формат длительности")
		return
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {

	if ds.Duration <= 0 {
		return "", errors.New("продолжительность должна быть больше 0")
	}

	if ds.Weight <= 0 {
		return "", errors.New("вес пользователя должен быть больше 0")
	}

	dist := (float64(ds.Steps) * StepLength) / 1000

	calories := ds.Weight * dist * 0.035

	res := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, dist, calories)

	return res, nil
}

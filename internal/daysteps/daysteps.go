package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	split := strings.Split(datastring, ",")
	if len(split) != 2 {
		return fmt.Errorf("Неверное кол-во вхождений")
	}

	ds.Steps, err = strconv.Atoi(split[0])
	if err != nil || ds.Steps <= 0 {
		return fmt.Errorf("Не смог перевести стрингу в шаги")
	}

	ds.Duration, err = time.ParseDuration(split[1])
	if err != nil || ds.Duration <= 0 {
		return fmt.Errorf("Не смог совладать со временем")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}

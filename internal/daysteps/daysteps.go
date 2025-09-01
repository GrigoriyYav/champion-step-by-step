package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	separatedData := strings.Split(data, ",")

	if len(separatedData) != 2 {
		return 0, 0, errors.New("wrong data format")
	}

	steps, err := strconv.Atoi(separatedData[0])
	if err != nil {
		return 0, 0, err
	} 

	if steps < 1 {
		return 0, 0, errors.New("zero number of steps")
	}

	duration, err := time.ParseDuration(separatedData[1])
	if err != nil {
		return 0, 0, err
	} 

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)

	if err != nil {
		fmt.Println(err)
		return ""
	} 

	if steps < 1 {
		return ""
	}

	distance := float64(steps) * float64(stepLength) / mInKm // дистанция в киллометрах

	caloriesBurned, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	} 

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
	steps, distance, caloriesBurned)
}

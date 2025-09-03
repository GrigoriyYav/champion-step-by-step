package daysteps

import (
	"errors"
	"fmt"
	"log"
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
	if steps <= 0 {
		return 0, 0, errors.New("steps must be positive")
	}

	duration, err := time.ParseDuration(separatedData[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, errors.New("duration must be positive")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)

	if err != nil {
		log.Println(err)
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

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distance, caloriesBurned)
}

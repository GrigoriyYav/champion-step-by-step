package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	separatedData := strings.Split(data, ",")

	if len(separatedData) != 3 {
		return 0, "", 0, errors.New("wrong data format")
	}

	steps, err := strconv.Atoi(separatedData[0])
	if err != nil || steps <= 0 {
    if err != nil {
        return 0, "", 0, err
    }
    return 0, "", 0, errors.New("steps must be positive")
	}


	duration, err := time.ParseDuration(separatedData[2])
	if err != nil || duration <= 0 {
    if err != nil {
        return 0, "", 0, err
    }
    return 0, "", 0, errors.New("duration must be positive")
	}

	return steps, separatedData[1], duration, nil
}

func distance(steps int, height float64) float64 {
	strideLength := stepLengthCoefficient * height
	return strideLength * float64(steps) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= time.Duration(0) {
		return 0
	}

	distance := distance(steps, height)

	hours := duration.Hours()
    
	averageSpeed := distance / hours
	
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}

	var dist, avgSpeed, calories float64
	
	switch activity {
	case "Ходьба":
			dist = distance(steps, height)
			avgSpeed = meanSpeed(steps, height, duration)
			calories, err = WalkingSpentCalories(steps, weight, height, duration)
			if err != nil {
					return "", fmt.Errorf("walking calories calculation failed: %w", err)
			}
			
	case "Бег":
			dist = distance(steps, height)
			avgSpeed = meanSpeed(steps, height, duration)
			calories, err = RunningSpentCalories(steps, weight, height, duration)
			if err != nil {
					return "", fmt.Errorf("running calories calculation failed: %w", err)
			}

	default: 
			return "", errors.New("неизвестный тип тренировки")
	}	


	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
        activity, 
        duration.Hours(), // Convert duration to hours
        dist, 
        avgSpeed, 
        calories)
    
    return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= time.Duration(0) {
		return 0, errors.New("data of running  cannot include zero")
	}

	meanSpeed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()
		
	calories := (weight * meanSpeed * durationInMinutes) / minInH 

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= time.Duration(0) {
		return 0, errors.New("data of walking cannot include zero")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH 

	return calories * walkingCaloriesCoefficient, nil
}	

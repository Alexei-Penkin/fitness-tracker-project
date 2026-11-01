package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимы для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories рассчитывает количество калорий, потраченных при ходьбе.
// Возвращает ошибку, если входные параметры некорректны.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateActivityParams(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient
	return calories, nil
}

// RunningSpentCalories рассчитывает количество калорий, потраченных при беге.
// Возвращает ошибку, если входные параметры некорректны.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateActivityParams(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

// MeanSpeed рассчитывает среднюю скорость в км/ч.
// Возвращает 0 при отрицательных шагах или неположительной продолжительности.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	return distance / duration.Hours()
}

// Distance рассчитывает дистанцию в километрах по количеству шагов и росту.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

// validateActivityParams проверяет параметры расчета калорий на корректность.
// Имеют смысл только положительные шаги, вес, рост и продолжительность.
func validateActivityParams(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	return nil
}
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
	stepLength = 0.65
	mInKm      = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	if strings.TrimSpace(data) == "" {
		return 0, 0, errors.New("пустая строка данных")
	}

	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("неверный формат данных: ожидалось два элемента, разделенных запятой")
	}

	stepsStr := parts[0]

	if stepsStr != strings.TrimSpace(stepsStr) {
		return 0, 0, errors.New("неверный формат количества шагов: пробелы не допускаются")
	}

	if strings.Contains(stepsStr, " ") {
		return 0, 0, errors.New("неверный формат количества шагов: пробелы внутри строки не допускаются")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil || steps <= 0 {
		return 0, 0, errors.New("ошибка преобразования количества шагов или значение меньше или равно нулю")
	}

	durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durationStr)
	if err != nil || duration <= 0 {
		return 0, 0, errors.New("ошибка преобразования длительности или продолжительность должна быть положительной")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println("Ошибка при парсинге данных:", err)
		return ""
	}

	if steps <= 0 {
		log.Println("Ошибка: количество шагов должно быть положительным")
		return ""
	}

	distance := float64(steps) * stepLength
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println("Ошибка при расчёте калорий:", err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distance/mInKm, calories)
}

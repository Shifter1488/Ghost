package daysteps

import (
 "fmt"
 "strconv"
 "strings"
 "time"
)

// parsePackage принимает строку с данными и возвращает количество шагов, продолжительность и ошибку
func parsePackage(data string) (int, time.Duration, error) {
 parts := strings.Split(data, ",") // Разделяем строку на части
 if len(parts) != 2 { // Проверяем, что получено две части
  return 0, 0, fmt.Errorf("неправильный формат данных")
 }

 steps, err := strconv.Atoi(parts[0]) // Преобразуем количество шагов в int
 if err != nil {
  return 0, 0, fmt.Errorf("не удалось преобразовать количество шагов: %v", err)
 }

 duration, err := time.ParseDuration(parts[1]) // Преобразуем продолжительность в time.Duration
 if err != nil {
  return 0, 0, fmt.Errorf("не удалось преобразовать продолжительность: %v", err)
 }

 return steps, duration, nil // Возвращаем количество шагов и продолжительность
}

// DayActionInfo парсит данные, вычисляет дистанцию и калории
func DayActionInfo(data string, weight, height float64) string {
 steps, duration, err := parsePackage(data) // Парсим данные
 if err != nil {
  return fmt.Sprintf("Ошибка: %v", err) // Возвращаем ошибку, если есть
 }

 distance := float64(steps) * 0.000762 // Примерный коэффициент для расчета расстояния (в км) на основе шагов
 calories := (weight * 0.035 * (float64(steps) / 1000)) + (0.029 * weight * (duration.Seconds() / 60)) // Расчет калорий

 return fmt.Sprintf("Вы прошли %d шагов, дистанция %.2f км, потрачено %.2f калорий за %s", steps, distance, calories, duration) // Формируем строку с результатами
}

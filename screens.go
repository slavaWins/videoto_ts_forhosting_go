package videoto

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Screenshots(inputFile string, outputDir string, numScreenshots int) {

	// Создаем папку для скриншотов, если её нет
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, 0755)
		if err != nil {
			fmt.Printf("Ошибка при создании папки: %v\n", err)
			return
		}
	}

	// Определение длительности видео с помощью ffprobe
	duration, err := getVideoDuration(inputFile)
	if err != nil {
		fmt.Printf("Ошибка при определении длительности видео: %v\n", err)
		return
	}

	// Если видео слишком короткое, делаем скриншоты от начала до конца
	if duration < float64(numScreenshots) {
		numScreenshots = int(duration) // Равняем количество скриншотов на длину в секундах
	}

	// Рассчитываем временные метки для каждого скриншота
	interval := duration / float64(numScreenshots)
	//fmt.Println("Длительность видео:", duration, "секунд. Интервал между скриншотами:", interval)

	for i := 0; i < numScreenshots; i++ {
		timeStamp := interval * float64(i)
		outputFile := fmt.Sprintf("%s/screenshot_%d.jpg", outputDir, i+1)

		// Генерация команды ffmpeg для скриншота
		cmd := exec.Command(
			"ffmpeg",
			"-ss", fmt.Sprintf("%f", timeStamp), // Временная метка
			"-i", inputFile, // Исходный файл
			"-frames:v", "1", // Один кадр
			"-vf", "scale=640:360:force_original_aspect_ratio=increase,crop=640:360",
			outputFile,                           // Выходной файл
			"-hide_banner", "-loglevel", "error", // Скрыть лишние сообщения
		)

		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("Ошибка при создании скриншота %d: %v\n%s\n", i+1, err, stderr.String())
			return
		}

		//	fmt.Printf("Скриншот %d успешно сохранен: %s\n", i+1, outputFile)
	}
}

// Функция для получения продолжительности видео через ffprobe
func getVideoDuration(filePath string) (float64, error) {
	cmd := exec.Command(
		"ffprobe",
		"-i", filePath,
		"-show_entries", "format=duration",
		"-v", "quiet",
		"-of", "csv=p=0",
	)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("ffprobe ошибка: %v\n%s", err, stderr.String())
	}

	// Преобразуем результат в float
	durationStr := strings.TrimSpace(out.String())
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, fmt.Errorf("некорректное значение длительности: %v", durationStr)
	}

	return duration, nil
}

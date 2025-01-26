package videoto

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func TsSegmentation(inputFile string, outputDir string, baseName string) {

	// Имя выходного плейлиста
	m3u8File := filepath.Join(outputDir, baseName+".m3u8")

	// Создаем директорию для выходных файлов
	if err := createDirIfNotExists(outputDir); err != nil {
		fmt.Println("Ошибка при создании директории:", err)
		return
	}

	// Формируем команду FFmpeg для создания HLS сегментов и .m3u8 файла
	// -i: входной файл
	// -c:v copy -c:a copy: копируем потоки без перекодирования
	// -hls_time 10: длина каждого сегмента (секции) в секундах
	// -hls_playlist_type vod: создаем VOD (по запросу) - это видеоконтент, доступный после завершения обработки.
	// output: это плейлист .m3u8, который автоматически включает ссылки на .ts сегменты
	cmd := exec.Command(
		"ffmpeg",
		"-i", inputFile, // Входной файл
		"-c:v", "copy", "-c:a", "copy", // Копирование видео и аудио без перекодирования
		"-hls_time", "10", // Длина сегментов в секундах
		"-hls_playlist_type", "vod", // Тип плейлиста (VOD для статичного контента)
		"-hls_segment_filename", // Шаблон имен для .ts сегментов
		filepath.Join(outputDir, baseName+"%03d.ts"),
		m3u8File, // Имя выходного плейлиста
	)

	// Выполняем команду
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при сегментации:", err)
		return
	}

	//fmt.Println("Сегментация завершена! Плейлист и сегменты сохранены.")
	//	fmt.Println("Путь к .m3u8:", m3u8File)
	//fmt.Println("Сегменты находятся в папке:", outputDir)
}

func WebpPreview(inputFile string, outputDir string) error {
	// Путь для выходного сжатого webp-файла
	outputFile := filepath.Join(outputDir, "preview.webp")

	cmd := exec.Command(
		"ffmpeg",
		"-i", inputFile,
		"-t", "5",
		"-an",
		"-vf", "scale=640:360:force_original_aspect_ratio=increase,crop=640:360",
		"-vcodec", "libwebp",
		"-q:v", "20",
		outputFile,
	)

	// Запускаем команду и проверяем на ошибки
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create webp: %v", err)
	}

	//fmt.Printf("WebP файл успешно создан: %s\n", outputFile)
	return nil
}

// Функция создаёт директорию, если её не существует
func createDirIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

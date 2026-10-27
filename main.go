// Пакет main — точка входа исполняемой программы
package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// printSection красиво оформляет заголовок секции вывода
func printSection(title string) {
	fmt.Println("── " + title + " " + strings.Repeat("─", 50-len([]rune(title))))
}

func main() {
	// ====================================================
	// БЛОК 1: Приветствие
	// ====================================================
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║         Hello, System! — Go v1.0             ║")
	fmt.Println("║   МДК.01.04 Системное программирование Go    ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Println()

	// ====================================================
	// БЛОК 2: Информация о рантайме Go
	// ====================================================
	printSection("🔧 Информация о Go-рантайме")
	fmt.Printf("  Версия Go:        %s\n", runtime.Version())
	fmt.Printf("  Операционная ОС:  %s\n", runtime.GOOS)
	fmt.Printf("  Архитектура CPU:  %s\n", runtime.GOARCH)
	fmt.Printf("  Логических ядер:  %d\n", runtime.NumCPU())
	fmt.Printf("  GOMAXPROCS:       %d\n", runtime.GOMAXPROCS(0))
	fmt.Println()

	// ====================================================
	// БЛОК 3: Пользователь и окружение
	// ====================================================
	printSection("👤 Пользователь и окружение")

	username := os.Getenv("USERNAME")
	if username == "" {
		username = os.Getenv("USER")
	}
	fmt.Printf("  Пользователь:     %s\n", username)

	homeDir := os.Getenv("USERPROFILE")
	if homeDir == "" {
		homeDir = os.Getenv("HOME")
	}
	fmt.Printf("  Домашний каталог: %s\n", homeDir)

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = "(не задан, используется ~/go по умолчанию)"
	}
	fmt.Printf("  GOPATH:           %s\n", gopath)

	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		goroot = runtime.GOROOT()
	}
	fmt.Printf("  GOROOT:           %s\n", goroot)

	// Задание 2: количество переменных окружения
	envVars := os.Environ()
	fmt.Printf("  Всего переменных окружения: %d\n", len(envVars))
	fmt.Println()

	// ====================================================
	// БЛОК 4: Файловая система
	// ====================================================
	printSection("📁 Файловая система")
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("  Ошибка получения текущей директории: %v\n", err)
	} else {
		fmt.Printf("  Текущая директория: %s\n", cwd)
	}
	fmt.Printf("  Временный каталог:  %s\n", os.TempDir())
	fmt.Println()

	// ====================================================
	// БЛОК 5: PATH
	// ====================================================
	printSection("🛤️  Переменная PATH")
	path := os.Getenv("PATH")

	var pathSeparator string
	if runtime.GOOS == "windows" {
		pathSeparator = ";"
	} else {
		pathSeparator = ":"
	}

	pathEntries := strings.Split(path, pathSeparator)
	fmt.Printf("  Всего записей в PATH: %d\n", len(pathEntries))

	limit := 5
	if len(pathEntries) < limit {
		limit = len(pathEntries)
	}
	for i, entry := range pathEntries[:limit] {
		fmt.Printf("  [%d] %s\n", i+1, entry)
	}
	if len(pathEntries) > 5 {
		fmt.Printf("  ... и ещё %d записей\n", len(pathEntries)-5)
	}
	fmt.Println()

	// ====================================================
	// БЛОК 6: Аргументы командной строки
	// ====================================================
	printSection("⌨️  Аргументы командной строки")
	fmt.Printf("  Программа: %s\n", os.Args[0])
	fmt.Printf("  Количество аргументов: %d\n", len(os.Args)-1)

	if len(os.Args) > 1 {
		fmt.Println("  Переданные аргументы:")
		for i, arg := range os.Args[1:] {
			fmt.Printf("    [%d] %s\n", i+1, arg)
		}
	} else {
		fmt.Println("  (аргументы не переданы)")
		fmt.Println("  Подсказка: запустите программу с аргументами:")
		fmt.Println("  go run main.go привет мир 123")
	}
	fmt.Println()

	// ====================================================
	// БЛОК 7 (доп.): Текущее время и дата
	// ====================================================
	printSection("🕒 Текущее время")
	now := time.Now()
	fmt.Printf("  Дата:  %s\n", now.Format("02.01.2006"))
	fmt.Printf("  Время: %s\n", now.Format("15:04:05"))
	fmt.Printf("  Полный формат (RFC3339): %s\n", now.Format(time.RFC3339))
	fmt.Println()

	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║          Программа завершена успешно         ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
}

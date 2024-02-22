package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var lines []string
	result := ""
	var n int32

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Читаем содержимое файла построчно
	scanner := bufio.NewScanner(file)

	fmt.Scanf("%d", &n)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			// Прибавляем  1 к ASCII коду каждого символа
			if (char < 123 && char > 96) || (char > 64 && char < 91) {
				char += n % 26
			}
			result += string(char)
		}
		lines = append(lines, result)
		result = ""
	}
	file, err = os.Create("text.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	// Записываем измененные строки обратно в файл
	writer := bufio.NewWriter(file)
	err = file.Truncate(0)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	fmt.Println("Текст в файле успешно изменен.")

}

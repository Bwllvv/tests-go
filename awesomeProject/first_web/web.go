package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите количество людей: ")
	countStr, _ := reader.ReadString('\n')
	count, _ := strconv.Atoi(strings.TrimSpace(countStr))

	people := make([]Person, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Введите имя и возраст человека %d (формат: имя возраст): ", i+1)
		input, _ := reader.ReadString('\n')
		parts := strings.Fields(input)
		people[i] = Person{Name: parts[0], Age: atoi(parts[1])}
	}

	fmt.Println("Введенные люди:")
	for _, person := range people {
		fmt.Printf("Имя: %s, Возраст: %d\n", person.Name, person.Age)
	}
}

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

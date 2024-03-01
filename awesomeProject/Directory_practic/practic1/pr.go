package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	str := "12124124"
	inte, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(0)
		return
	}
	inte *= 2
	stri := strconv.Itoa(inte)
	fmt.Println(stri + "sd")
	file_opn, erro := os.OpenFile("file.txt", os.O_RDWR, 0666)
	if erro != nil {
		fmt.Println(1)
		return
	}
	//defer file_opn.Close()
	scanne := bufio.NewScanner(file_opn)
	for scanne.Scan() {
		fmt.Printf("%s\n", scanne.Text())
	}
	fmt.Println(4)
	_, errore := file_opn.WriteString("helloeptads\n")
	if errore != nil {
		fmt.Println(2)
		return
	}

	file_opn.Seek(0, 0)
	defer file_opn.Close()
	for scanne.Scan() {
		fmt.Printf("%s\n", scanne.Text())
	}
	fmt.Println(4)

}

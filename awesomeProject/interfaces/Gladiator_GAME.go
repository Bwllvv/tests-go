package main

import (
	"fmt"
	"math"
	"math/rand"
)

type motions interface {
	run(a int) error
	jump(a int) error
	getX() int
	getY() int
}

type noob struct {
	X       int
	Y       int
	Name    string
	stamina int
}

func (n *noob) getX() int {
	if n.X > 80 {
		n.X = n.X % 80
	} else if n.X < 0 {
		n.X += 80
	}
	return n.X
}
func (n *noob) getY() int {
	if n.Y > 10 {
		n.Y = n.Y % 10
	} else if n.Y < 0 {
		n.Y += 10
	}
	return n.Y
}

func (n *noob) run(a int) error {
	if int(math.Abs(float64(a))) <= n.stamina {
		n.X += a
		n.stamina -= int(math.Abs(float64(a)))
		return nil
	} else {
		return fmt.Errorf("Не хватает стамины")
	}
}

func (n *noob) jump(a int) error {
	if int(math.Abs(float64(a)))*10 <= n.stamina {
		n.Y -= a
		n.stamina -= int(math.Abs(float64(a))) * 10
		return nil
	} else {
		return fmt.Errorf("Не хватает стамины")
	}
}

type pro struct {
	X       int
	Y       int
	Name    string
	old     int
	stamina int
}

func (n *pro) getX() int {
	return n.X
}
func (n *pro) getY() int {
	return n.Y
}

func (n *pro) run(a int) error {
	if a <= n.stamina {
		n.X += a
		n.stamina -= a
		return nil
	} else {
		return fmt.Errorf("Не хватает стамины")
	}
}

func (n *pro) jump(a int) error {
	if a*10 <= n.stamina {
		n.Y -= a
		n.stamina -= int(math.Abs(float64(a))) * 10
		return nil
	} else {
		return fmt.Errorf("Не хватает стамины")
	}
}

//func motion(m motions) {
//	switch m.(type) {
//	case *noob:
//		for (m.run()==nil||)
//
//
//	}
//}

func main() {
	var person motions
	var flags []int
	var shoot [][]int
	person = &noob{80 / 2, 10 / 2, "Stive", 100}
	pers := 0
	shooting := 0
	killing := 0
	var flag string
	slice := GenEnemy()
	shoot = [][]int{{person.getX() + 1, person.getY()}, {person.getX() - 1, person.getY()},
		{person.getX(), person.getY() + 1}, {person.getX(), person.getY() - 1}}
	for {
		fmt.Scan(&flag)
		switch flag {
		case "w":
			if person.jump(1) != nil {
				fmt.Println(person.jump(1))
			}
		case "a":
			if person.run(-5) != nil {
				fmt.Println(person.jump(1))
			}
		case "s":
			if person.jump(-1) != nil {
				fmt.Println(person.jump(1))
			}
		case "d":
			if person.run(5) != nil {
				fmt.Println(person.jump(1))
			}
		case "f":
			shooting = 1
		case "changeNoob":
			person = &noob{80 / 2, 10 / 2, "Stive", 100}
			pers = 0
		case "changePro":
			person = &pro{80 / 2, 10 / 2, "Lev", 19, 1000}
			pers = 1
		default:

		}
		if shooting == 0 {
			shoot = [][]int{{person.getX() + 1, person.getY()}, {person.getX() - 1, person.getY()},
				{person.getX(), person.getY() + 1}, {person.getX(), person.getY() - 1}}
		}
		flags = pole(person, pers, shooting, killing, slice, shoot)
		shooting = flags[0]
		killing = flags[1]
		//fmt.Println(person.(*noob).stamina)
		fmt.Println(shoot)
	}
}

func pole(person motions, pers int, shooting int, killing int, slice []int, shoot [][]int) []int {
	X := person.getX()
	Y := person.getY()
	if shooting == 1 {
		shoot[0][0] += 1
		shoot[1][0] -= 1
		shoot[2][1] += 1
		shoot[3][1] -= 1
		if shoot[0][0] >= 80 && shoot[1][0] <= 0 && shoot[2][1] >= 10 && shoot[3][1] <= 0 {
			shooting = 0
		}
	}
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 80; j++ {
			if shoot[0][0] == slice[0] && shoot[0][1] == slice[1] || shoot[1][0] == slice[0] && shoot[1][1] == slice[1] || shoot[2][0] == slice[0] && shoot[2][1] == slice[1] || shoot[3][0] == slice[0] && shoot[3][1] == slice[1] {
				killing = 1
			}
			if pers == 0 && j == X && i == Y {
				fmt.Print("*")
			} else if slice[0] == j && slice[1] == i && killing != 1 {
				fmt.Print("E")
			} else if slice[0] == j && slice[1] == i && killing == 1 {
				fmt.Print("X")
			} else if pers == 1 && j == X && i == Y {
				fmt.Print("&")
			} else if shoot[0][0] == j && shoot[0][1] == i || shoot[1][0] == j && shoot[1][1] == i || shoot[2][0] == j && shoot[2][1] == i || shoot[3][0] == j && shoot[3][1] == i {
				fmt.Print("#")
			} else if i == 0 || i == 10 {
				fmt.Print("-")
			} else if j == 0 || j == 80 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}

		}

		fmt.Print("\n")
	}
	if killing == 0 {
		slice[rand.Intn(2)] += rand.Intn(2)
	}
	if slice[0] > 80 {
		slice[0] = slice[0] % 80
	} else if slice[1] > 10 {
		slice[1] = slice[1] % 10
	}
	fmt.Println(slice)
	flags := []int{shooting, killing}
	return flags
}

func GenEnemy() []int {
	slice := []int{rand.Intn(80), rand.Intn(10)}
	return slice
}

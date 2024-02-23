package main

import "fmt"

type payer interface {
	pay(a int) error
}

type Card struct {
	number int
	money  int
}

func (c *Card) pay(a int) error {
	if a < c.money {
		c.money -= a
		return nil
	} else {
		return fmt.Errorf("недостаточно средств на карте")
	}
}

type Cash struct {
	money int
}

func (c *Cash) pay(a int) error {
	if a < c.money {
		c.money -= a
		return nil
	} else {
		return fmt.Errorf("недостаточно средств")
	}
}

type PhonePay struct {
	phone string
	money int
}

func (c *PhonePay) pay(a int) error {
	if a < c.money {
		c.money -= a
		return nil
	} else {
		return fmt.Errorf("недостаточно средств на телефоне")
	}
}

func buy(p payer) {
	switch p.(type) {
	case *Card:
		if p.pay(100) == nil {
			fmt.Println("Покупка успешна,спинано 100р")
		} else {
			fmt.Println(p.pay(100))
		}
	case *Cash:
		if p.pay(100) == nil {
			fmt.Println("Покупка успешна,принятно 100р")
		} else {
			fmt.Println(p.pay(100))
		}

	case *PhonePay:
		if p.pay(100) == nil {
			fmt.Println("Покупка успешна,с телефона спинано 100р")
		} else {
			fmt.Println(p.pay(100))
		}

	default:

	}
}
func main() {
	var person payer
	person = &Card{200, 530}
	buy(person)
	person = &Cash{800}
	buy(person)
	person = &PhonePay{"93485394573", 530}
	buy(person)
}

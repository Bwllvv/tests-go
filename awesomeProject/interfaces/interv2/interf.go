package main

import "fmt"

type payer interface {
	pay(a int) error
}

type CreditCard struct {
	number int
	money  int
}

func (CC *CreditCard) pay(a int) error {
	if CC.money >= a {
		CC.money -= a
		return nil
	} else {
		return fmt.Errorf("Нихуя бабок нет...")
	}
}

type Cash struct {
	value int
}

func (CC *Cash) pay(a int) error {
	if CC.value >= a {
		CC.value -= a
		return nil
	} else {
		return fmt.Errorf("Нихуя бабок нет... сходи к банкомату ")
	}
}

func buy(p payer) {
	switch p.(type) {
	case *CreditCard:
		if p.pay(100) == nil {
			fmt.Println("спасибо за покупку с карты")
		} else {
			fmt.Println(p.pay(100))
		}
	case *Cash:
		if p.pay(100) == nil {
			fmt.Println("спасибо за покупку и оплату наличными")
		} else {
			fmt.Println(p.pay(100))
		}

	default:
		fmt.Println("как ты заплатил????")
	}
}

func main() {
	card := &CreditCard{9737242, 1200}
	fmt.Println(card)
	buy(card)
	fmt.Println(card)
	cash := &Cash{value: 60}
	fmt.Println(cash)
	buy(cash)
	fmt.Println(cash)
}

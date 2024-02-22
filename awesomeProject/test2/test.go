package main

import "fmt"

type miner interface {
	Dig(string) string
}

type ore struct {
	//ore_type string;
	//weight int;
}

func (d *ore) Dig(ore_type string) string {
	return fmt.Sprintf("%s is digging", ore_type)
}
func motion(v miner, name string) {
	fmt.Printf(v.Dig(name))
}

func main() {
	motion(&ore{}, "Gold")
}

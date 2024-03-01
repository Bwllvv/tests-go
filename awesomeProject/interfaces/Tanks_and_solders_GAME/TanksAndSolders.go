package main

import (
	"fmt"
	"math/rand"
	"time"
)

type shooter interface {
	shoot(shootCount int) int
	getAmmo() int
	getHP() int
}

type solder struct {
	HP     int
	name   string
	weight int
	weapon weapon
}

func (s *solder) getAmmo() int {
	return s.weapon.ammo
}
func (s *solder) getHP() int {
	return s.HP
}

func (s *solder) shoot(shootCount int) int {
	return s.weapon.shoot(shootCount)
}

type weapon struct {
	gun    string
	ammo   int
	damage int
}

func (w *weapon) shoot(shootCount int) int {
	shoot := 0
	for i := 0; i < shootCount; i++ {
		shoot += w.damage * (rand.Intn(100-50) + 50)
	}
	return shoot
}

type cannon struct {
	size             string
	projectile       string
	projectile_count int
	damage           int
}

func (t *tank) shoot(shootCount int) int {
	return t.cannon.shoot(shootCount)
}
func (t *tank) getAmmo() int {
	return t.cannon.projectile_count
}
func (t *tank) getHP() int {
	return t.HP
}

func (c *cannon) shoot(shootCount int) int {
	shoot := 0
	for i := 0; i < shootCount; i++ {
		shoot += c.damage * (rand.Intn(100-50) + 50)
	}
	return shoot
}

type tank struct {
	HP     int
	model  string
	weight int
	cannon cannon
}

func main() {
	EndGame := true
	countUnitSolder := 0
	countUnitTank := 0
	sumdmg_solder := 0
	sumdmg_tank := 0
	var shooters shooter
	countUnitSolder = 15
	countUnitTank = 3
	savedamageSOLDER := 0
	savedamageTANK := 0
	for EndGame {
		EndGame = field(countUnitSolder, countUnitTank)
		if EndGame == false {
			return
		}
		shooters = &solder{10000, "Karl", 80, weapon{"M4A4", 30, 1}}
		sumdmg_solder = SummaryDamage(countUnitSolder, shooters.getAmmo(), shooters)
		fmt.Printf("Unit count:%d, Count shoot: %d, Damage: %d, Saved damage:%d\n", countUnitSolder, shooters.getAmmo(), sumdmg_solder, savedamageSOLDER)

		shooters = &tank{50000, "JacPanther", 15000, cannon{"44/22mm", "armor-piercing", 3, 50}}
		savedamageSOLDER += sumdmg_solder
		if savedamageSOLDER > shooters.getHP() {
			for savedamageSOLDER > shooters.getHP() {
				countUnitTank--
				savedamageSOLDER -= shooters.getHP()
			}
		}
		sumdmg_tank = SummaryDamage(countUnitTank, shooters.getAmmo(), shooters)
		fmt.Printf("Unit count:%d, Count shoot: %d, Damage: %d, Saved damage:%d\n", countUnitTank, shooters.getAmmo(), sumdmg_tank, savedamageTANK)
		shooters = &solder{10000, "Karl", 80, weapon{"M4A4", 30, 1}}
		savedamageTANK += sumdmg_tank
		if savedamageTANK > shooters.getHP() {
			for savedamageTANK > shooters.getHP() {
				countUnitSolder--
				savedamageTANK -= shooters.getHP()
			}
		}
		time.Sleep(5 * time.Second)
	}
	//shooters = &solder{"Karl", 80, weapon{"M4A4", 30, 10}}

}

func SummaryDamage(countUnit int, countShoot int, shooters shooter) int {
	sum := 0
	for i := 0; i < countUnit; i++ {
		sum += shooters.shoot(countShoot)
	}
	return sum
}

func field(countSolder int, countTank int) bool {
	if countSolder <= 0 {
		fmt.Println("=============")
		fmt.Println("Tanks win!")
		fmt.Println("=============")
		return false
	} else if countTank <= 0 {
		fmt.Println("=============")
		fmt.Println("Solders win!")
		fmt.Println("=============")
		return false
	}
	fmt.Println("__________________________________________")
	for i := 0; i < countSolder; i++ {
		fmt.Print("; ")
	}
	fmt.Print("\n")
	fmt.Println("*war field*")
	for i := 0; i < countTank; i++ {
		fmt.Print("{!} ")
	}
	fmt.Print("\n")
	fmt.Println("__________________________________________")
	return true
}

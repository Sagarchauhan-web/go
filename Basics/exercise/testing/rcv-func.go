package main

import "fmt"

const (
	MAX_HEALTH = 100
	MAX_ENERGY = 100
)

type Player struct {
	name string
	health int
	energy int
}

func (p *Player) addHealth(health int) {
	if (p.health + health) > MAX_HEALTH {
		p.health = MAX_HEALTH
	} else {
		p.health += health
	}
}

func (p *Player) applyDamage(health int) {
	if (p.health - health) < 0 {
		p.health = 0
	} else {
		p.health -= health
	}
	
}

func (p *Player) addEnergy(energy int) {	
	if (p.energy + energy) > MAX_ENERGY {
		p.energy = MAX_ENERGY
	} else {
		p.energy += energy
	}
}

func (p *Player) consumeEnergy(energy int) {	
	if (p.energy - energy) < 0 {
		p.energy = 0
	} else {
		p.energy -= energy
	}
}


func main() {
	player := Player{name: "Sagar", health: 10, energy: 10}

	fmt.Println("Initial:", player)

	player.addHealth(10)
	player.applyDamage(5)
	fmt.Println("After health modify:\n", player)

	player.addEnergy(10)
	player.consumeEnergy(5)
	fmt.Println("After health and energy modify:\n", player)

	player.applyDamage(999)
	player.consumeEnergy(999)
	fmt.Println("After health and energy modify:\n", player)
}

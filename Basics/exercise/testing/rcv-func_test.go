//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func newPlayer() Player {
	return Player{name: "Sagar", health: 10, energy: 10}
}

func TestHealth(t *testing.T) {
	p := newPlayer()
	p.addHealth(MAX_HEALTH)

	if p.health != MAX_HEALTH {
		t.Fatalf("Expected health to be %v, got %v", MAX_HEALTH, p.health)
	}

	if p.health > MAX_HEALTH {
		t.Fatalf("health can not be greater than %v", MAX_HEALTH)
	}

	p.applyDamage(MAX_HEALTH + 1)
	if p.health != 0 {
		t.Fatalf("Expected health to be 0, got %v", p.health)
	}

	if p.health < 0 {
		t.Fatalf("health can not be less than 0")
	}
}
func TestEnergy(t *testing.T) {
	p := newPlayer()
	p.addEnergy(MAX_ENERGY)

	if p.energy != MAX_ENERGY {
		t.Fatalf("Expected energy to be %v, got %v", MAX_ENERGY, p.energy)
	}

	if p.energy > MAX_ENERGY {
		t.Fatalf("energy can not be greater than %v", MAX_ENERGY)
	}

	p.consumeEnergy(MAX_ENERGY + 1)
	if p.energy != 0 {
		t.Fatalf("Expected energy to be 0, got %v", p.energy)
	}

	if p.energy < 0 {
		t.Fatalf("energy can not be less than 0")
	}
}

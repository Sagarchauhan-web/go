package main

import "fmt"

type Passenger struct {
	Name string
	TickerNumber int
	Boarded bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
		casey := Passenger{
			"Casey", 1, false,
		}

		fmt.Println(casey)

		var (
			bill = Passenger{Name: "bill", TickerNumber: 2,}
			ella = Passenger{Name: "ella", TickerNumber: 3,}
		)
		fmt.Println(bill, ella)

		var heidi Passenger
		heidi.Name = "heidi"
		heidi.TickerNumber = 4
		fmt.Println(heidi)

		casey.Boarded = true
		bill.Boarded = true

		if bill.Boarded {
			fmt.Println("Bill has boarded the bus")
		}
		if casey.Boarded {
			fmt.Println(casey.Name,"has boarded the bus")
		}

		heidi.Boarded = true
		bus := Bus{heidi}
		fmt.Println(bus)
		fmt.Println(bus.FrontSeat.Name, "is in the front of the bus")
}

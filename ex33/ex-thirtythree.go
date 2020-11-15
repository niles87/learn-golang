package main

import "fmt"

// structs
// create a vehicle struct with fields of doors and color
// create two new types that embed the vehicle type

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	trk := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}
	sdn := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}
	fmt.Println(trk, sdn)
	fmt.Println(trk.color)
	fmt.Println(sdn.doors)
}

package main

import (
	"fmt"

	"github.com/tarm/serial"
)

//
//	VIRTUAL KEYBOARD MAPPING
//

func Move(p [2]int, d int) ([2]int, string) {
	keyboardtv := [4][10]string{{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}, {"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}, {"a", "s", "d", "f", "g", "h", "j", "k", "l", "?"}, {"@", "z", "x", "c", "v", "b", "n", "m", ",", "."}}
	dimension := p[0]
	position := p[1]
	P := keyboardtv[dimension][position]
	//
	// Aa
	//
	if d == 97 {
		//fmt.Println("Left")
		P = keyboardtv[dimension][position-1]
		p[1] = position - 1
	}
	//
	// Dd
	//
	if d == 100 {
		//fmt.Println("Right")
		P = keyboardtv[dimension][position+1]
		p[1] = position + 1
	}
	//
	// Ww
	//
	if d == 119 {
		//fmt.Println("UP")
		P = keyboardtv[dimension-1][position]
		p[0] = dimension - 1
	}
	//
	//
	//
	if d == 115 {
		//fmt.Println("Down")
		P = keyboardtv[dimension+1][position]
		p[0] = dimension + 1
	}
	//
	//
	//
	//fmt.Println(P)
	return p, P
}

//
//
//

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Print(err)
	}
	for {
		buf := make([]byte, 2)
		_, err = s.Read(buf)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%x \n", buf[:len(buf)-1])
	}
}

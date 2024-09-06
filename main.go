package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tarm/serial"
)

//
//
//

func Move(p [2]int, d string) ([2]int, string) {
	keyboardtv := [4][10]string{{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}, {"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}, {"a", "s", "d", "f", "g", "h", "j", "k", "l", "?"}, {"@", "z", "x", "c", "v", "b", "n", "m", ",", "."}}
	dimension := p[0]
	position := p[1]
	P := keyboardtv[dimension][position]
	//
	//	<
	//
	if d == "04fb07" {
		//fmt.Println("Left")
		P = keyboardtv[dimension][position-1]
		p[1] = position - 1
	}
	//
	// >
	//
	if d == "04fb06" {
		//fmt.Println("Right")
		P = keyboardtv[dimension][position+1]
		p[1] = position + 1
	}
	//
	// up
	//
	if d == "04fb40" {
		//fmt.Println("UP")
		P = keyboardtv[dimension-1][position]
		p[0] = dimension - 1
	}
	//
	// down
	//
	if d == "04fb41" {
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
	//
	//
	//
	var strslice strings.Builder
	buf := make([]byte, 2)
	//
	//
	//
	var Point [2]int
	var Lastpoint [2]int
	var Char, Enter string
	var Passwd []string
	//
	//	Initial Position
	//
	Point[0] = 2
	Point[1] = 4
	//
	//
	//
	for {
		for i := 1; i <= 3; i++ {
			_, err = s.Read(buf)
			if err != nil {
				fmt.Print(err)
			}
			//
			//
			//
			tostring := hex.EncodeToString(buf[:len(buf)-1])
			// fmt.Printf("%x", buf[:len(buf)-1])

			// fmt.Print(".")
			strslice.WriteString(tostring)
		}
		// fmt.Println(strslice.String())
		//
		// Left
		//
		if strslice.String() == "04fb07" {
			Lastpoint, Char = Move(Point, "04fb07")
			//fmt.Println("Last point: ", Lastpoint)
			for i, _ := range Lastpoint {
				Point[i] = Lastpoint[i]
			}
			fmt.Print("Position: ", Point)
			fmt.Println(" Key: " + Char)
			//fmt.Println(Char)
			Enter = Char
			strslice.Reset()
		}
		//
		// Right
		//
		if strslice.String() == "04fb06" {
			Lastpoint, Char = Move(Point, "04fb06")
			//fmt.Println("Last point: ", Lastpoint)
			for i, _ := range Lastpoint {
				Point[i] = Lastpoint[i]
			}
			fmt.Print("Position: ", Point)
			fmt.Println(" Key: " + Char)
			Enter = Char
			strslice.Reset()
		}
		//
		//
		//
		if strslice.String() == "04fb40" {
			Lastpoint, Char = Move(Point, "04fb40")
			//fmt.Println("Last point: ", Lastpoint)
			for i, _ := range Lastpoint {
				Point[i] = Lastpoint[i]
			}
			//fmt.Println("Position: ", Point)
			fmt.Print("Position: ", Point)
			fmt.Println(" Key: " + Char)
			Enter = Char
			strslice.Reset()
		}
		//
		//
		//
		if strslice.String() == "04fb41" {
			Lastpoint, Char = Move(Point, "04fb41")
			//fmt.Println("Last point: ", Lastpoint)
			for i, _ := range Lastpoint {
				Point[i] = Lastpoint[i]
			}
			//fmt.Println("Point: ", Point)
			fmt.Print("Position: ", Point)
			fmt.Println(" Key: " + Char)
			Enter = Char
			strslice.Reset()
		}

		//
		// ENTER
		//
		if strslice.String() == "04fb44" {
			fmt.Println("Enter")
			Passwd = append(Passwd, Enter)
			strslice.Reset()
			// fmt.Println(Passwd)
		}
		//
		// Qq Quit
		//
		if strslice.String() == "04fb5b" {
			fmt.Println(Passwd)
			break
		}
	}
}

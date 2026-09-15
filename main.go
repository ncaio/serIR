package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tarm/serial"
)

const (
	// USBTERM CONFIG
	usbport = "/dev/ttyUSB0"
	baud    = 9600
	// Key Mapping
	left  = "04fb07"
	right = "04fb06"
	up    = "04fb40"
	down  = "04fb41"
	enter = "04fb44"
	exit  = "04fb5b"
	// Others
	banner = "serIR - por ncaio - v2 - https://github.com/ncaio/serIR"
)

var keyboardtv = [4][10]string{
	{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
	{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
	{"a", "s", "d", "f", "g", "h", "j", "k", "l", "?"},
	{"@", "z", "x", "c", "v", "b", "n", "m", ",", "."},
}

//
//
//

func Move(p [2]int, d string) ([2]int, string) {
	//
	// Aqui se constroi a Matriz do teclado virtual
	//
	dimension := p[0]
	position := p[1]
	P := keyboardtv[dimension][position]
	//
	//	<-
	//
	if d == left {
		P = keyboardtv[dimension][position-1]
		p[1] = position - 1
	}
	//
	// ->
	//
	if d == right {
		P = keyboardtv[dimension][position+1]
		p[1] = position + 1
	}
	//
	// up
	//
	if d == up {
		//fmt.Println("UP")
		P = keyboardtv[dimension-1][position]
		p[0] = dimension - 1
	}
	//
	// down
	//
	if d == down {
		P = keyboardtv[dimension+1][position]
		p[0] = dimension + 1
	}
	//
	//
	//
	return p, P
}

func destacar(tecla string) string {
	return fmt.Sprintf("\033[1;30;43m[%s]\033[0m", tecla)
}

func main() {
	//
	// Aqui configura o Serial USB e escuta
	//
	c := &serial.Config{Name: usbport, Baud: baud}
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
	fmt.Println(banner)
	//
	//
	//
	fmt.Println("Teclado Virtual:")
	for i, linha := range keyboardtv {
		for j, tecla := range linha {

			if i == Point[0] && j == Point[1] {
				fmt.Print(destacar(tecla), " ")
			} else {
				fmt.Printf("[%s] ", tecla)
			}
		}

		fmt.Println()
	}
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
			strslice.WriteString(tostring)
		}
		//
		// Left
		//
		if strslice.String() == left {
			Lastpoint, Char = Move(Point, left)
			for i, _ := range Lastpoint {
				Point[i] = Lastpoint[i]
			}
			fmt.Print("Position: ", Point)
			fmt.Println(" Key: " + Char)
			Enter = Char
			strslice.Reset()
		}
		//
		// Right
		//
		if strslice.String() == right {
			Lastpoint, Char = Move(Point, right)
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
		if strslice.String() == up {
			Lastpoint, Char = Move(Point, up)
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
		if strslice.String() == down {
			Lastpoint, Char = Move(Point, down)
			for i, _ := range Lastpoint {
				Point[i] = Lastpoint[i]
			}
			fmt.Print("Position: ", Point)
			fmt.Println(" Key: " + Char)
			Enter = Char
			strslice.Reset()
		}

		//
		// ENTER
		//
		if strslice.String() == enter {
			fmt.Println("Enter")
			Passwd = append(Passwd, Enter)
			strslice.Reset()
		}
		//
		// Qq Quit
		//
		if strslice.String() == exit {
			fmt.Println(Passwd)
			break
		}
	}
}

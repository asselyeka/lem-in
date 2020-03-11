package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	start              bool
	end                bool
	nameOfTheRoom      string
	connectedWithRooms []string
	ants               bool
	x, y               int
}

func CreateRooms(fileName string) ([]Room, int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()

	rawBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawBytes), "\n")
	var rooms []Room
	var ants int
	startRoom := 0
	endRoom := 0

	for indx, l := range lines {

		if strings.Contains(l, "##start") == true && indx+1 < len(lines) {
			startRoom = indx + 1
			var tmpl Room
			roomData := strings.Split(lines[indx+1], " ")
			tmpl.start = true
			tmpl.end = false
			tmpl.nameOfTheRoom = roomData[0]
			coorX, err1 := strconv.Atoi(roomData[1])
			coorY, err2 := strconv.Atoi(roomData[2])
			if err1 != nil || err2 != nil {
				fmt.Println("ERROR")
				os.Exit(0)
			}
			tmpl.x = coorX
			tmpl.y = coorY
			tmpl.connectedWithRooms = []string{}
			rooms = append(rooms, tmpl)
			continue
		}
		if strings.Contains(l, "##end") == true && indx+1 < len(lines) {
			endRoom = indx + 1
			var tmpl Room
			roomData := strings.Split(lines[indx+1], " ")
			tmpl.start = false
			tmpl.end = true
			tmpl.nameOfTheRoom = roomData[0]
			coorX, err1 := strconv.Atoi(roomData[1])
			coorY, err2 := strconv.Atoi(roomData[2])
			if err1 != nil || err2 != nil {
				fmt.Println("ERROR")
				os.Exit(0)
			}
			tmpl.x = coorX
			tmpl.y = coorY
			tmpl.connectedWithRooms = []string{}
			rooms = append(rooms, tmpl)
			continue
		}
		if indx != 0 && (indx == startRoom || indx == endRoom) {
			continue
		}
		roomData := strings.Split(l, " ")
		if len(roomData) == 3 {
			var tmpl Room
			tmpl.start = false
			tmpl.end = false
			tmpl.nameOfTheRoom = roomData[0]
			coorX, err1 := strconv.Atoi(roomData[1])
			coorY, err2 := strconv.Atoi(roomData[2])
			if err1 != nil || err2 != nil {
				continue
			}
			tmpl.x = coorX
			tmpl.y = coorY
			tmpl.connectedWithRooms = []string{}
			rooms = append(rooms, tmpl)
			continue
		}
		roomConnect := strings.Split(l, "-")
		if len(roomConnect) == 2 {
			room1 := roomConnect[0]
			room2 := roomConnect[1]
			fmt.Println(room1)
			fmt.Println(room2)
			for i, r := range rooms {
				if room1 == r.nameOfTheRoom {
					rooms[i].connectedWithRooms = append(rooms[i].connectedWithRooms, room2)
				}
				if room2 == r.nameOfTheRoom {
					rooms[i].connectedWithRooms = append(rooms[i].connectedWithRooms, room1)
				}
			}
			continue
		}

		ant, err := strconv.Atoi(l)
		if err == nil {
			ants = ant
		}
	}
	return rooms, ants
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please provide with ant-farm.txt")
		return
	}
	fileName := args[1]
	rooms, ants := CreateRooms(fileName)

	fmt.Println(ants)
	for _, r := range rooms {
		fmt.Print("Name:")
		fmt.Println(r.nameOfTheRoom)
		fmt.Print("CoordX:")
		fmt.Println(r.x)
		fmt.Print("CoordY:")
		fmt.Println(r.y)
		fmt.Print("start:")
		fmt.Println(r.start)
		fmt.Print("end:")
		fmt.Println(r.end)
		fmt.Print("Connected with:")
		for _, roo := range r.connectedWithRooms {

			fmt.Print(roo, ", ")
		}
		fmt.Println()
	}

}

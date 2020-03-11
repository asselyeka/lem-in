package main

type Room struct {
	start              bool
	end                bool
	nameOfTheRoom      string
	connectedWithRooms []*Room
	ants               bool
	x, y               int
}

func CreateRooms (fileName string) []Room, int {
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

		for indx, l := range lines {
			startRoom := 0
			endRoom := 0
			if strings.Contains(l,"##start") == true && indx+1 < len(lines){
				startRoom = indx+1
				var tmpl Room
				roomData := strings.Split(lines[indx+1]," ")
				tmpl.start = true
				tmpl.end = false
				tmpl.nameOfTheRoom = roomData[0]
				tmpl.x, err1 := strconv.Atoi(roomData[1])
				tmpl.y, err2 := strconv.Atoi(roomData[2])
				if err1 != nil || err2 != nil {
					fmt.Println("ERROR")
					os.Exit(0)
				}
				rooms = append(rooms, tmpl)
				continue
			}
			if strings.Contains(l,"##end") == true && indx+1 < len(lines){
				endRoom = indx+1
				var tmpl Room
				roomData := strings.Split(lines[indx+1]," ")
				tmpl.start = false
				tmpl.end = true
				tmpl.nameOfTheRoom = roomData[0]
				tmpl.x, err1 := strconv.Atoi(roomData[1])
				tmpl.y, err2 := strconv.Atoi(roomData[2])
				if err1 != nil || err2 != nil {
					fmt.Println("ERROR")
					os.Exit(0)
				}
				rooms = append(rooms, tmpl)
				continue
			}
			if indx == startRoom || indx == endRoom {
				continue
			}
			roomData := strings.Split(l, " ")
			if len(roomData) == 3 {
				var tmpl Room
				tmpl.start = false
				tmpl.end = false
				tmpl.nameOfTheRoom = roomData[0]
				tmpl.x, err1 := strconv.Atoi(roomData[1])
				tmpl.y, err2 := strconv.Atoi(roomData[2])
				if err1 != nil || err2 != nil {
					continue
				}
				rooms = append(rooms, tmpl)
			} else {
				roomData = strings.Split(l, "-")
				if len(roomData) == 2 {
					room1 := roomData[0]
					room2 := roomData[1]
					var room1Link *Room
					var room2Link *Room
					for i, r := range rooms {
						if room1 == r.nameOfTheRoom {
							room1Link = *r
						}
						if room2 == r.nameOfTheRoom {
							room2Link = *r
						}
					}
					
				} else {
					ants, err := strconv.Atoi(l)
					if err != nil {
						continue
					}

				}
			}

		}


}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please provide with ant-farm.txt")
		return
	}
	fileName := args[1]
	rooms, ants := CreateRooms(fileName)



}

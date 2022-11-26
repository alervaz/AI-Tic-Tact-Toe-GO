package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Global variables
var Player []int
var Running bool
var Strategies map[int][]int
var Strategy []int

func main() {
	//Initialize Variables
	Strategies = map[int][]int{
		1: []int{1, 2, 3},
		2: []int{4, 5, 6},
		3: []int{7, 8, 9},
		4: []int{1, 4, 7},
		5: []int{2, 5, 8},
		6: []int{3, 6, 9},
		7: []int{1, 5, 9},
		8: []int{3, 5, 7},
	}
	Grid := GetGrid()
	Box := []string{" ", "x", "o"}
	//Pick random Strategy for the main rival
	rand.Seed(time.Now().UnixNano())
	Strategy = Strategies[rand.Intn(8)+1]

	//Decide who is first
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		PlayerFirst(Grid, Box)
	} else {
		RivalFirst(Grid, Box)
	}
}
func PlayerFirst(Grid map[int]int, Box []string) {
	Running = true
	fmt.Println("Welcome to Tic-Tac-Toe")
	for Running {
		//It has 2 Check win and draw because if the AI dosnt know if can move the program crashes
		//Check it alredy won twice to shut down the app and dont make it crash
		PrintBoard(Grid, Box)
		PlayerInput(Grid)
		Won := CheckWin(Grid, Box)
		if Won {
			return
		}
		Draw := CheckIfDraw(Grid, Box)
		if Draw {
			return
		}
		RivalChoice(Grid)
		Won = CheckWin(Grid, Box)
		if Won {
			return
		}
		Draw = CheckIfDraw(Grid, Box)
		if Draw {
			return
		}
	}
}

func RivalFirst(Grid map[int]int, Box []string) {
	Running = true
	fmt.Println("Welcome to Tic-Tac-Toe")
	for Running {
		//It has 2 Check win and draw because if the AI dosnt know if can move the program crashes
		//Check it alredy won twice to shut down the app and dont make it crash
		RivalChoice(Grid)
		PrintBoard(Grid, Box)
		Won := CheckWin(Grid, Box)
		if Won {
			return
		}
		Draw := CheckIfDraw(Grid, Box)
		if Draw {
			return
		}
		PlayerInput(Grid)
		Won = CheckWin(Grid, Box)
		if Won {
			return
		}
		Draw = CheckIfDraw(Grid, Box)
		if Draw {
			return
		}
	}
}

// This map is mapped to the current Cell (btw I called Bow a Cell yes I am dumb but its too late now) the index its the Cell and the value
// goes from 0 to 2 if is 0 no one has that cell if its 1 the player has it and if its 2 the rival has it
func GetGrid() map[int]int {
	Grid := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}
	return Grid
}

// Prints the Board to the command line
func PrintBoard(Grid map[int]int, Box []string) {
	C1 := Box[Grid[1]]
	C2 := Box[Grid[2]]
	C3 := Box[Grid[3]]
	C4 := Box[Grid[4]]
	C5 := Box[Grid[5]]
	C6 := Box[Grid[6]]
	C7 := Box[Grid[7]]
	C8 := Box[Grid[8]]
	C9 := Box[Grid[9]]
	fmt.Printf("%s|%s|%s\n", C1, C2, C3)
	fmt.Println("-+-+-")
	fmt.Printf("%s|%s|%s\n", C4, C5, C6)
	fmt.Println("-+-+-")
	fmt.Printf("%s|%s|%s\n\n", C7, C8, C9)
}

// Checks if someone won
func CheckWin(Grid map[int]int, Box []string) bool {
	Player = []int{}
	Rival := []int{}

	for i, Cell := range Grid {
		if Cell == 1 {
			Player = append(Player, i)
		} else if Cell == 2 {
			Rival = append(Rival, i)
		}
	}

	if (Contains(Player, 1) && Contains(Player, 2) && Contains(Player, 3)) || (Contains(Player, 4) && Contains(Player, 6) && Contains(Player, 6)) || (Contains(Player, 7) && Contains(Player, 8) && Contains(Player, 9)) || (Contains(Player, 1) && Contains(Player, 4) && Contains(Player, 7)) || (Contains(Player, 2) && Contains(Player, 5) && Contains(Player, 8)) || (Contains(Player, 3) && Contains(Player, 6) && Contains(Player, 9)) || (Contains(Player, 1) && Contains(Player, 5) && Contains(Player, 9)) || (Contains(Player, 3) && Contains(Player, 5) && Contains(Player, 7)) {
		fmt.Println("Player Wins!")
		PrintBoard(Grid, Box)
		Running = false
		return true
	} else if (Contains(Rival, 1) && Contains(Rival, 2) && Contains(Rival, 3)) || (Contains(Rival, 4) && Contains(Rival, 6) && Contains(Rival, 6)) || (Contains(Rival, 7) && Contains(Rival, 8) && Contains(Rival, 9)) || (Contains(Rival, 1) && Contains(Rival, 4) && Contains(Rival, 7)) || (Contains(Rival, 2) && Contains(Rival, 5) && Contains(Rival, 8)) || (Contains(Rival, 3) && Contains(Rival, 6) && Contains(Rival, 9)) || (Contains(Rival, 1) && Contains(Rival, 5) && Contains(Rival, 9)) || (Contains(Rival, 3) && Contains(Rival, 5) && Contains(Rival, 7)) {
		fmt.Println("Rival Wins!")
		PrintBoard(Grid, Box)
		Running = false
		return true
	}
	return false
}

// A basic contains function
func Contains(Player []int, item int) bool {
	for _, v := range Player {
		if v == item {
			return true
		}
	}
	return false
}

// Checks for player input
func PlayerInput(Grid map[int]int) {
	fmt.Print("Pick a cell: ")
	Valid := false
	Choice := ""
	for !Valid {
		fmt.Scanln(&Choice)
		Number, err := strconv.Atoi(Choice)
		CanContinue := true
		if err != nil {
			CanContinue = false
		}

		if CanContinue {
			for i, v := range Grid {
				if i == Number && v == 0 {
					Valid = true
					Grid[i] = 1
				}
			}
		}
		if !Valid {
			fmt.Print("Invalid cell pick another one: ")
		}
	}
}

func RivalChoice(Grid map[int]int) {
	//This checks if the rival has one left to win and if it does applies moves to the missing one
	CanWin, Missing := CanWin(Grid, Strategy)
	if CanWin {
		if Grid[Missing] == 0 {
			Grid[Missing] = 2
			return
		}
	}

	//This checks if the player can win and if it does block him also it is random tat can fail to not make the game imposible if the rival
	//started and because humans make errors
	Has, Num := ContainsTwo()
	rand.Seed(time.Now().UnixNano())
	if Has && rand.Intn(4) == 2 {
		if Grid[Num] == 0 {
			Grid[Num] = 2
			return
		}
	}
	//This runs until the move of the rival is valid and it uses a strategy if that strategy cannot be used chamge to a different one
	Valid := false
	StrategiesTried := 0
	if !Running {
		return
	}
	for !Valid {
		//Checks if any of the Cells for the strategy is 1 if it is change else sais is valid
		Count := 0
		for _, v := range Strategy {
			if Grid[v] != 1 {
				Count += 1
			}
		}
		if Count == 3 {
			Valid = true
		}
		//If is valid move to the Cell its not occuped
		if Valid {
			for _, v := range Strategy {
				if Grid[v] == 0 {
					Grid[v] = 2
					return
				}
			}
		} else {
			//Changes strategy
			rand.Seed(time.Now().UnixNano())
			Strategy = Strategies[rand.Intn(8)+1]
			StrategiesTried += 1
		}

		if StrategiesTried == 8 {
			Move := false
			for !Move {
				rand.Seed(time.Now().UnixNano())
				RandNum := rand.Intn(9) + 1
				if Grid[RandNum] == 0 {
					Grid[RandNum] = 2
					Move = true
					return
				}
			}
			Valid = true
		}
	}
}

func CheckIfDraw(Grid map[int]int, Box []string) bool {
	C1 := Grid[1]
	C2 := Grid[2]
	C3 := Grid[3]
	C4 := Grid[4]
	C5 := Grid[5]
	C6 := Grid[6]
	C7 := Grid[7]
	C8 := Grid[8]
	C9 := Grid[9]
	if C1 != 0 && C2 != 0 && C3 != 0 && C4 != 0 && C5 != 0 && C6 != 0 && C7 != 0 && C8 != 0 && C9 != 0 {
		fmt.Println("Draw")
		PrintBoard(Grid, Box)
		Running = false
		return true
	}
	return false
}

// This cheks all the strategies and checks if the player have 2 in one so the AI can block it
// Not means which is the one who lefts to the player to win
func ContainsTwo() (bool, int) {
	for _, Strat := range Strategies {
		Count := 0
		Not := 0
		for _, Num := range Strat {
			if Contains(Player, Num) {
				Count += 1
			} else {
				Not = Num
			}

		}
		if Count == 2 {
			return true, Not
		}
	}
	return false, 0
}

// Checks if the rival can win
func CanWin(Grid map[int]int, Strategy []int) (bool, int) {
	Count := 0
	Missing := 0
	for _, v := range Strategy {
		if Grid[v] == 2 {
			Count += 1
		} else {
			Missing = v
		}
	}
	if Count == 2 {
		return true, Missing
	}
	return false, 0
}

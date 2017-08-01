package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type User string

func SwitchUser(user *User) {
	if *user == "X" {
		*user = "0"
	} else {
		*user = "X"
	}

}

func Clear() {
	cmd, err := exec.Command("clear").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(cmd))
}

func PrintFullTable() {
	fullTable := [][]string{
		{"|1|", "|2|", "|3|"},
		{"|4|", "|5|", "|6|"},
		{"|7|", "|8|", "|9|"},
	}

	fmt.Println()
	for _, line := range fullTable {
		fmt.Print("\t")
		for _, elem := range line {
			fmt.Print(elem)
		}
		fmt.Println()
	}
	fmt.Println()
}

func whereIsIt(mark string, table [][]string) map[int]bool {
	cells := make(map[int]bool)
	x := 0
	for _, line := range table {
		for _, item := range line {
			x++
			if mark == item {
				cells[x] = true
			}
		}
	}
	return cells
}

func contains(set [3]int, cells map[int]bool) (result bool) {
	var a bool
	for _, cell := range set {
		_, a = cells[cell]
		if a {
			result = true
		} else {
			return false
		}
	}
	return result
}

func hasWinner(table [][]string) (result string, anyWinner bool) {
	sets := [][3]int{{1, 2, 3}, {1, 4, 7}, {1, 5, 9}, {2, 5, 8}, {3, 6, 9}, {3, 5, 7}, {4, 5, 6}, {7, 8, 9}}
	for _, mark := range []string{"|X|", "|0|"} {
		cells := whereIsIt(mark, table)
		if len(cells) > 0 {
			for _, set := range sets {
				if contains(set, cells) {
					return mark, true
				}
			}
		}
	}
	return "", false
}

func SetMark(elem *string, user User) {
	*elem = strings.Join(
		append(append(strings.Split(*elem, "")[:1], string(user)), strings.Split(*elem, "")[2:]...),
		"",
	)
}

func main() {
	var hand int
	var winner string
	var any bool
	table := [][]string{
		{"| |", "| |", "| |"},
		{"| |", "| |", "| |"},
		{"| |", "| |", "| |"},
	}
	Clear()
	PrintFullTable()
	user := User("X")
	for {
		fmt.Println("* * * The old game * * *")
		fmt.Printf("This is the %s's time\n", user)
		fmt.Println("[Press 1-9 number to mark a cell]")
		fmt.Scanln(&hand)
		Clear()
		round := 0
		for i, line := range table {
			fmt.Print("\t")
			for j, cell := range line {
				round++
				if round == hand {
					elem := &table[i][j]
					if *elem == "| |" {
						SetMark(elem, user)
						SwitchUser(&user)
					}
					fmt.Print(*elem)
				} else {
					fmt.Print(cell)
				}
			}
			fmt.Println()
		}
		if winner, any = hasWinner(table); any {
			break
		}
	}
	if any {
		fmt.Printf("\n... and the Oscar goes to %s!\n", winner)
	} else {
		fmt.Println("\n... there is no winner among to losers.")
	}
}

package main

import "fmt"

func odd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

func getMark(idx int) string {
	if odd(idx) {
		return "|X|"
	}
	return "|0|"
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

func main() {
	var hand int
	var winner string
	var any bool
	table := [][]string{
		{"| |", "| |", "| |"},
		{"| |", "| |", "| |"},
		{"| |", "| |", "| |"},
	}
	fullTable := [][]string{
		{"|1|", "|2|", "|3|"},
		{"|4|", "|5|", "|6|"},
		{"|7|", "|8|", "|9|"},
	}

	fmt.Println("* * * The old game * * *")
	for _, line := range fullTable {
		fmt.Print("\t")
		for _, elem := range line {
			fmt.Print(elem)
		}
		fmt.Println()
	}
	fmt.Println()
	for x := 1; x <= 9; x++ {
		mark := getMark(x)
		fmt.Printf("This is the %s's time\n", mark)
		fmt.Println("[Press 1-9 number to mark a cell]")
		fmt.Scanln(&hand)
		item := 0
		for i, line := range table {
			for j, cell := range line {
				item++
				if item == hand {
					elem := &table[i][j]
					if *elem == "| |" {
						*elem = mark
					} else {
						x--
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
		fmt.Printf("And the Oscar goes to %s!\n", winner)
	} else {
		fmt.Println("There is no winner among to losers...")
	}
}

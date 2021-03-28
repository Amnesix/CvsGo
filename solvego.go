package main

import "fmt"

const tst294 string = "....12.8.5..8...........73..85......6..9.3....2..6..4..6..2....7..3.9..2..2..19.."
const tst295 string = "....14....3....2...7..........9...3.6.1.............8.2.....1.4....5.6.....7.8..."
const CASE_VIDE = '.'

type Sudoku [9][9]byte

var sudoku Sudoku

func (s *Sudoku) isok(r, c int, v byte) bool {
	for i := 0; i < 9; i++ {
		if s[r][i] == v {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if s[j][c] == v {
			return false
		}
	}
	for j := 3 * (r / 3); j < 3*(1+r/3); j++ {
		for i := 3 * (c / 3); i < 3*(1+c/3); i++ {
			if s[j][i] == v {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) chercheVide() (int, int, bool) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if s[r][c] == CASE_VIDE {
				return r, c, true
			}
		}
	}
	return 0, 0, false
}

func (s *Sudoku) Solve() bool {
	r, c, ok := s.chercheVide()
	if !ok {
		return true
	}
	var v byte
	for v = '1'; v <= '9'; v++ {
		if s.isok(r, c, v) {
			s[r][c] = v
			if s.Solve() {
				return true
			}
			s[r][c] = '.'
		}
	}
	return false
}

func (s *Sudoku) Affiche() {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			fmt.Printf("%c ", sudoku[r][c])
		}
		fmt.Println()
	}
}

func main() {
	for r := 0; r < 9; r++ {
		line := tst295[r*9 : (r+1)*9]
		for c := 0; c < 9; c++ {
			sudoku[r][c] = line[c]
		}
	}
	sudoku.Affiche()
	fmt.Println()
	sudoku.Solve()
	sudoku.Affiche()
}

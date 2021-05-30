package main

type Board struct {
	cols  int
	rows  int
	board [][]bool
	count [][]int
}

func newBoard(numCols, numRows int) *Board {
	var board = make([][]bool, numCols)
	var count = make([][]int, numCols)

	for col := range board {
		board[col] = make([]bool, numRows)
		count[col] = make([]int, numRows)
	}

	return &Board{
		cols:  numCols,
		rows:  numRows,
		board: board,
		count: count,
	}
}

func (b *Board) inBounds(x, y int) bool {
	return x >= 0 && x < b.cols && y >= 0 && y < b.rows
}

func (b *Board) at(x, y int) bool {
	return b.board[x][y]
}

func (b *Board) toggle(x, y int) {
	b.board[x][y] = !b.board[x][y]
}

func (b *Board) update() {
	b.updateCount()

	for i := 0; i < b.cols; i++ {
		for j := 0; j < b.rows; j++ {
			val := b.count[i][j]

			if val < 2 {
				b.board[i][j] = false
			} else if val == 3 {
				b.board[i][j] = true
			} else if val > 3 {
				b.board[i][j] = false
			}
		}
	}
}

func (b *Board) updateCount() {
	b.clearCount()

	for col, cols := range b.board {
		for row, cellAlive := range cols {

			if cellAlive {
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x == 0 && y == 0 {
							continue
						}
						if b.inBounds(col+x, row+y) {
							b.count[col+x][row+y]++
						}
					}
				}
			}
		}
	}
}

func (b *Board) clearCount() {
	for col, cols := range b.board {
		for row, _ := range cols {
			b.count[col][row] = 0
		}
	}
}

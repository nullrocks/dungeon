package dungeon

import (
	"strings"
	"os/exec"
	"os"
	"fmt"
)

type Position struct {
	X int
	Y int
}

type Field struct {
	Length int
	Width  int
	Keys   []Key
	Door   []Door
	Plane  [][]string
}

type Movable interface {
	Move(f *Field, dir string, dis int) bool
}

func NewField(m []string, ks []Key, ds []Door) (Field, Position) {
	length := len(m)
	width := len(m[0])
	matrix := make([][]string, length)
	var startingPoint Position
	for y, row := range m {
		matrix[y] = strings.Split(row, "")
		for x, block := range matrix[y] {
			if block == "*" {
				startingPoint = Position{x, y}
			}
		}
	}
	f := Field{length, width, ks, ds, matrix}
	return f, startingPoint
}

func (f Field) MoveToPosition(p Position, ks ...Key) (bool, Position, []Key) {

	canMove := false
	np := Position{p.X, p.Y}
	var keys []Key = []Key{}

	if p.Y >= f.Length {
		np.Y = 0
	} else if p.Y < 0 {
		np.Y = f.Length - 1
	}

	if p.X >= f.Width {
		np.X = 0;
	} else if p.X < 0 {
		np.X = f.Width - 1
	}

	if f.Plane[np.Y][np.X] != "#" {
		canMove = true
		fmt.Printf("p: %v, new: %v", p, np)
	}

	if f.Plane[np.Y][np.X] == KeyChar {
		// slice of keys in this position
		keys = f.pickKeys(np)
	} else if f.Plane[np.Y][np.X] == ClosedDoorChar {

	}

	return canMove, np, keys
}

func (f Field) Render() string {
	view := "\n"
	for _, row := range f.Plane {
		view += strings.Join(row, "") + "\n"
	}

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println(view)

	return view
}

func (f *Field) pickKeys(p Position) []Key {
	return []Key{}
}

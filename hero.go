package dungeon

type Hero struct {
	Name       string
	SightRange uint
	Position
	Keys       []Key
}

func New(n string, sr uint, p Position) (Hero, error) {
	return Hero{n, sr, p, nil}, nil
}

func (h *Hero) Move(f *Field, dir string, dis int) bool {

	np := Position{h.X, h.Y}
	moved := false

	switch dir {
	case "w":
		np.Y--
	case "d":
		np.X++
	case "a":
		np.X--
	case "s":
		np.Y++
	default:
		return false
	}

	moved, np := f.CanMoveToPosition(np)

	if moved {

	}

	return moved
}

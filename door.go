package dungeon

type Door struct {
	open bool
	Key
}

func (d *Door) Open(k Key) bool {
	if d.Type == k.Type {
		d.open = true
	}

	return d.open
}

func (d Door) IsOpen() bool {
	return d.open
}

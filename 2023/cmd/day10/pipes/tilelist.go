package pipes

type TileList []Tile

func NewTileList(capacity int) *TileList {
	tl := make(TileList, 0, capacity)
	return &tl
}

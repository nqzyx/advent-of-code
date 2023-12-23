package pipes

type TileList []Tile

func NewTileList(capacity int) *TileList {
	tl := make(TileList, 0, capacity)
	return &tl
}

func (l TileList) HasPathTiles() bool {
	for _, tile := range l {
		if tile.OnPath {
			return true
		}
	}
	return false
}

func (l TileList) IsInsidePath(d Direction) bool {
	type tuple [2]any
	// pathTileSeen := false
	var count int
	for _, tile := range l {
		if tile.OnPath {
			x := tuple{tile.PipeType, d}
			switch x {
			// Intersections
			case tuple{NorthSouth, East}, tuple{NorthSouth, West},
				tuple{EastWest, North}, tuple{EastWest, South}:
				count++
				// pathTileSeen = true
			default:
			}
		} /* else {
			if pathTileSeen {
				return count%2 == 1
			}
		} */
	}
	return count%2 == 1
}

func (l TileList) HasPipeType(pipeType PipeType) bool {
	for _, tile := range l {
		if tile.PipeType == pipeType {
			return true
		}
	}
	return false
}

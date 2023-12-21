package pipes

type Pipe string

const (
	NS Pipe = "|"
	EW Pipe = "-"
	NE Pipe = "L"
	SE Pipe = "F"
	SW Pipe = "7"
	NW Pipe = "J"
    Start Pipe = "S"
    Empty Pipe = "."
)

type (
	PipeConnections map[Pipe]Connections
	Connections     struct {
		North, South, East, West bool
	}
)

var pipeConnections PipeConnections = PipeConnections{
	EW: Connections{East: true, West: true},
	NE: Connections{North: true, East: true},
	NS: Connections{North: true, South: true},
	NW: Connections{North: true, West: true},
	SE: Connections{South: true, East: true},
	SW: Connections{South: true, West: true},
}

type Tile struct {
    Coords [2]int
    Pipe Pipe
}

func (p Tile) X() int {
    return p.Coords[0]
}

func (p Tile) Y() int {
    return p.Coords[1]
}

type Map [][]rune
/*struct {
    StartingTile Tile
    Tiles [][]Tile
}*/

func NewMap(input []string) (map Map) {
    for _, line := range input {
        row := make([]Tile, 0, 5)
        map.Tiles = append(map.Tiles, row)
        for _, point := range line {

            switch point {
                case Empty:
                    continue
            }
        }
    }
    return
}

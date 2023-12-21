package pipes

type Pipe rune

const (
	NS    Pipe = '|'
	EW    Pipe = '-'
	NE    Pipe = 'L'
	SE    Pipe = 'F'
	SW    Pipe = '7'
	NW    Pipe = 'J'
	Start Pipe = 'S'
	Empty Pipe = '.'
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

package inspect

//go:generate stringer -type=Dir
type Dir int

const (
	// ONC access between each children groups:
	ONC Dir = iota
	ENT
	EXT
)

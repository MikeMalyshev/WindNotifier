package bot

import "fmt"

// Location coordinates aggregated as latitude, longitude
type Coordinates [2]float32

type Location struct {
	Name  string
	Coord Coordinates
}

func (loc Location) String() string {
	return fmt.Sprintf("%s (%f,%f)", loc.Name, loc.Coord[0], loc.Coord[1])
}

func NewLocation(c Coordinates, name string) Location {
	return Location{Name: name, Coord: c}
}

func FindLocationByCoords(loc Coordinates) (Location, error) {
	// TODO
	return Location{}, nil
}

func FindLocationByName(name string) (Location, error) {
	// TODO
	return Location{}, nil
}

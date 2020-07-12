package portdomain

import (
	"context"
	"errors"
)

type PortID string

type Geo2DPoint struct {
	Lati  string
	Longi string
}

type Port struct {
	ID          PortID
	Name        string
	Country     string
	City        string
	Province    string
	Timezone    string
	Code        string
	Coordinates Geo2DPoint
	Regions     []string
	Alias       []string
	Unlocs      []string
}

var ErrUnkPort = errors.New("unknown port")

type PortRepository interface {
	Store(context.Context, *Port) error
	FindByID(context.Context, PortID) (*Port, error)
}

package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc/status"

	"github.com/walkline/shippingpg/portdomain"
	"github.com/walkline/shippingpg/portdomain/server/pb"
)

type PortServiceServer struct {
	pb.UnimplementedPortDomainServiceServer

	repo   portdomain.PortRepository
	logger log.Logger
}

func NewPortServiceServer(r portdomain.PortRepository, logger log.Logger) *PortServiceServer {
	return &PortServiceServer{
		repo:   r,
		logger: logger,
	}
}

func (s PortServiceServer) Store(ctx context.Context, p *pb.Port) (*pb.Empty, error) {
	defer func(t time.Time) {
		s.logger.Log(
			"action", "storing",
			"port_id", p.GetID(),
			"took", time.Since(t),
		)
	}(time.Now())

	port := &portdomain.Port{
		ID:       portdomain.PortID(p.GetID()),
		Name:     p.GetName(),
		Country:  p.GetCountry(),
		City:     p.GetCity(),
		Province: p.GetProvince(),
		Timezone: p.GetTimezone(),
		Code:     p.GetCode(),
		Regions:  p.GetRegions(),
		Alias:    p.GetAlias(),
		Unlocs:   p.GetUnlocs(),
	}

	if p.GetCoordinates() != nil {
		port.Coordinates = portdomain.Geo2DPoint{
			Lati:  p.GetCoordinates().GetLati(),
			Longi: p.GetCoordinates().GetLongi(),
		}
	}

	err := s.repo.Store(ctx, port)
	if err != nil {
		return nil, err
	}

	return new(pb.Empty), nil
}

func (s PortServiceServer) FindByID(ctx context.Context, id *pb.PortID) (*pb.Port, error) {
	defer func(t time.Time) {
		s.logger.Log(
			"action", "searching",
			"port_id", id.GetID(),
			"took", time.Since(t),
		)
	}(time.Now())

	p, err := s.repo.FindByID(ctx, portdomain.PortID(id.GetID()))
	if err != nil {
		if err == portdomain.ErrUnkPort {
			return nil, status.Error(http.StatusNotFound, err.Error())
		}

		return nil, err
	}

	port := &pb.Port{
		ID:       string(p.ID),
		Name:     p.Name,
		Country:  p.Country,
		City:     p.City,
		Province: p.Province,
		Timezone: p.Timezone,
		Code:     p.Code,
		Regions:  p.Regions,
		Alias:    p.Alias,
		Unlocs:   p.Unlocs,
		Coordinates: &pb.Geo2DPoint{
			Lati:  p.Coordinates.Lati,
			Longi: p.Coordinates.Longi,
		},
	}

	return port, nil
}

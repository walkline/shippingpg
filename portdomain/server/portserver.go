package server

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/walkline/shippingpg/portdomain"
	"github.com/walkline/shippingpg/portdomain/server/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		Unlocs:   p.GetAlias(),
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

func (PortServiceServer) FindByID(context.Context, *pb.PortID) (*pb.Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePort not implemented")
}

package port

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/walkline/shippingpg/clientapi"
	"github.com/walkline/shippingpg/clientapi/port/pb"
)

type KVRepository struct {
	client pb.PortDomainServiceClient
}

func NewKVRepository(client pb.PortDomainServiceClient) *KVRepository {
	return &KVRepository{
		client: client,
	}
}

type portJSON struct {
	Name        string        `json:"name"`
	Country     string        `json:"country"`
	City        string        `json:"city"`
	Province    string        `json:"province"`
	Code        string        `json:"code"`
	Timezone    string        `json:"timezone"`
	Alias       []string      `json:"alias"`
	Regions     []string      `json:"regions"`
	Unlocs      []string      `json:"unlocs"`
	Coordinates []json.Number `json:"coordinates"`
}

func (r *KVRepository) Store(ctx context.Context, kv *clientapi.KeyValue) error {
	b, ok := kv.Value.(json.RawMessage)
	if !ok {
		return errors.New("can't cast value to json.RawMessage")
	}

	portjson := portJSON{}
	err := json.Unmarshal(b, &portjson)
	if err != nil {
		return err
	}

	grpcPort := &pb.Port{
		ID:       kv.Key,
		Name:     portjson.Name,
		Province: portjson.Province,
		Country:  portjson.Country,
		City:     portjson.City,
		Code:     portjson.Code,
		Timezone: portjson.Timezone,
		Alias:    portjson.Alias,
		Regions:  portjson.Regions,
		Unlocs:   portjson.Unlocs,
	}

	if len(portjson.Coordinates) == 2 {
		grpcPort.Coordinates = &pb.Geo2DPoint{
			Lati:  string(portjson.Coordinates[0]),
			Longi: string(portjson.Coordinates[1]),
		}
	}
	_, err = r.client.Store(ctx, grpcPort)
	return err
}

func (r *KVRepository) FindByKey(ctx context.Context, key string) (*clientapi.KeyValue, error) {
	resp, err := r.client.FindByID(ctx, &pb.PortID{
		ID: key,
	})
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(&portJSON{
		Name:     resp.Name,
		Province: resp.Province,
		Country:  resp.Country,
		City:     resp.City,
		Code:     resp.Code,
		Timezone: resp.Timezone,
		Alias:    resp.Alias,
		Regions:  resp.Regions,
		Unlocs:   resp.Unlocs,
		Coordinates: []json.Number{
			json.Number(resp.Coordinates.GetLati()),
			json.Number(resp.Coordinates.GetLongi()),
		},
	})
	if err != nil {
		return nil, err
	}

	return &clientapi.KeyValue{
		Key:   key,
		Value: json.RawMessage(b),
	}, nil
}

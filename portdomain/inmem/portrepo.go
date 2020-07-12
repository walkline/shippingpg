package inmem

import (
	"context"
	"sync"

	"github.com/walkline/shippingpg/portdomain"
)

type PortRepository struct {
	mutex sync.RWMutex
	m     map[portdomain.PortID]portdomain.Port
}

func NewPortRepository() *PortRepository {
	return &PortRepository{
		m: map[portdomain.PortID]portdomain.Port{},
	}
}

func (r *PortRepository) Store(c context.Context, p *portdomain.Port) error {
	r.mutex.Lock()
	r.m[p.ID] = *p
	r.mutex.Unlock()
	return nil
}

func (r *PortRepository) FindByID(c context.Context, id portdomain.PortID) (*portdomain.Port, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	v, found := r.m[id]
	if !found {
		return nil, portdomain.ErrUnkPort
	}
	return &v, nil
}

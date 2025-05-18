package resolver

import (
	"context"
	"math/big"

	"github.com/0xPolygonID/verifier-backend/internal/config"
	"github.com/iden3/go-iden3-auth/v2/pubsignals"
	"github.com/iden3/go-iden3-auth/v2/state"
	log "github.com/sirupsen/logrus"
)

var _ pubsignals.StateResolver = &FabricResolver{}

type FabricResolver struct {
	Config config.ResolverSettingsAttrs
}

func NewFabricResolver(networkConfig config.ResolverSettingsAttrs) (*FabricResolver, error) {
	return &FabricResolver{Config: networkConfig}, nil
}

func (r *FabricResolver) Resolve(ctx context.Context, id *big.Int, state *big.Int) (*state.ResolvedState, error) {
	log.Info("Resolve", "id", id, "state", state)
	return nil, nil
}

func (r *FabricResolver) ResolveGlobalRoot(ctx context.Context, id *big.Int) (*state.ResolvedState, error) {
	log.Info("ResolveGlobalRoot", "id", id)
	return nil, nil
}

package reactor

import (
	"context"

	model "github.com/golang001/deltav/model/gomodel"
	"github.com/golang001/deltav/utils"
)

/* FusionReactor is a reactor that implaments the GRPC Reactor interface.
It works by reacting helium and deuterium and produces heat, energy.
TODO: For future realism, it should produce a small amount of gamma radiation
and should slowly degrade over time due to neutron bombardment because a
small number of deuterium-deuterium reactions should occur.
*/
type FusionReactor struct {
	heliumBuffer    *model.StorageTankClient
	deuteriumBuffer *model.StorageTankClient
}

// React implaments the GRPC interface for a Reactor.
func (mr *FusionReactor) React(ctx context.Context, req *model.ReactRequest) (res *model.ReactResponse,
	err error) {
	defer utils.CapturePanic(ctx, res, &err)
	// *res = mr.response
	// err = mr.err
	return
}

/*
TODO:
Options-
 this reactor needs to be topped off (self-contained, harder to manage, needs to implemnt storage interface)
 we can say fuck the buffer and it can request fuel from the main tanks directly (chaos dependencies, maybe a good thing)
 or we can have a monolith service for the ship to make requests from (maybe doesn't scale)
*/

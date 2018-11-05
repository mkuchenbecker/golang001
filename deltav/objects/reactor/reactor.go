package reactor

import (
	"context"
	"fmt"
	"math"

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
	heliumBuffer    model.StorageTankClient
	deuteriumBuffer model.StorageTankClient
	maxEfficency    float64 // Should never be above 0.7

}

const (
	// http://www.asi.org/adb/02/09/he3-intro.html
	tJPerKgH3                  float64 = 2157000
	tJPerKgDeut                float64 = tJPerKgH3 / 0.66
	inverseTeraJoulesPerKgH3           = 1 / tJPerKgH3
	inverseTeraJoulesPerKgDeut         = 1 / tJPerKgDeut
)

// React implaments the GRPC interface for a Reactor.
func (mr *FusionReactor) React(ctx context.Context, req *model.ReactRequest) (res *model.ReactResponse,
	err error) {
	defer utils.CapturePanic(ctx, res, &err)

	// To get the desired energy, we need to increase the reaction mass in proportion to the efficency of the reaction.
	invMaxEfficency := 1 / mr.maxEfficency
	h3Desired := req.DesiredEnergyTeraJoules * inverseTeraJoulesPerKgH3 * invMaxEfficency

	heliumRes, heliumErr := mr.heliumBuffer.WithdrawStorage(ctx,
		&model.StorageRequest{Storage: &model.Storage{Amount: h3Desired}},
	)
	if heliumErr != nil {
		err = fmt.Errorf("could not access helium: %s", heliumErr.Error())
		return
	}

	// We might get less helium back than anticipated.
	h3Fraction := heliumRes.Storage.Amount / h3Desired // Expensive, consider inverse math.

	deutDesired := req.DesiredEnergyTeraJoules * inverseTeraJoulesPerKgDeut * invMaxEfficency
	deutRes, deutErr := mr.deuteriumBuffer.WithdrawStorage(ctx,
		&model.StorageRequest{Storage: &model.Storage{Amount: deutDesired}},
	)
	if deutRes != nil {
		err = fmt.Errorf("could not access helium: %s", deutErr.Error())
		return
	}
	// We might get less deuterium back than anticipated.
	deutFraction := deutRes.Storage.Amount / deutDesired // Expensive, consider inverse math.

	// We take the min energy fraction as the true fraction.
	energyFraction := math.Min(h3Fraction, deutFraction)

	if energyFraction > 0.99 { // Correct for floating point errors.
		energyFraction = 1.0
	}

	energy := energyFraction * req.DesiredEnergyTeraJoules
	heat := (1 - mr.maxEfficency) * energy // A fraction of the total output is heat.

	res = &model.ReactResponse{
		Outputs: []*model.ReactorOutput{
			&model.ReactorOutput{Type: model.ReactorOutput_ENERGY_TJOULES, Amount: energy},
			&model.ReactorOutput{Type: model.ReactorOutput_HEAT_TJOULES, Amount: heat},
		},
	}
	err = nil

	return
}

/*
TODO:
Options-
 this reactor needs to be topped off (self-contained, harder to manage, needs to implemnt storage interface)
 we can say fuck the buffer and it can request fuel from the main tanks directly (chaos dependencies, maybe a good thing)
 or we can have a monolith service for the ship to make requests from (maybe doesn't scale)
*/

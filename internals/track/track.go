package track

import (
	"context"
	"encoding/json"

	"github.com/siddhantprateek/Trakd/internals/consignment"
)

type packageTrack struct {
	pc consignment.PackageClient
}

// TrackByVehicleID implements consignment.PackageInit.
func (p *packageTrack) TrackByVehicleID(ctx context.Context, id string) (*consignment.Package, error) {
	bytes, err := p.pc.ConsumeByVehicleID(ctx, id)
	if err != nil {
		return nil, err
	}

	var res consignment.Package
	err = json.Unmarshal(bytes, &res)
	return &res, err
}

func NewConsignmentTrack(pClient consignment.PackageClient) consignment.PackageInit {
	return &packageTrack{pc: pClient}
}

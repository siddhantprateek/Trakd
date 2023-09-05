package consignment

import "context"

type Package struct {
	From      string `json:"from"`
	To        string `json:"to"`
	VehicleID string `json:"vehicleId"`
}

type PackageInit interface {
	TrackByVehicleID(ctx context.Context, id string) (*Package, error)
}

type PackageClient interface {
	ConsumeByVehicleID(ctx context.Context, vehichleID string) ([]byte, error)
}

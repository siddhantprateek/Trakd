package consignment

import "context"

type Package struct {
	From      string `json:"from"`
	To        string `json:"to"`
	VehicleID string `json:"vehicleId"`
}

type PackageInit interface {
	TrackByVechileID(ctx context.Context, id string) (*Package, error)
}

type PackageClient interface {
	ConsumeByVechileID(ctx context.Context, vehichleID string) ([]byte, error)
}

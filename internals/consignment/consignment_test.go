package consignment

import (
	"context"
	"errors"
	"testing"
)

type MockPackageInit struct {
	PackageData map[string]*Package
}

func (m *MockPackageInit) TrackByVechileID(ctx context.Context, id string) (*Package, error) {
	pkg, found := m.PackageData[id]
	if !found {
		return nil, errors.New("Package not found")
	}
	return pkg, nil
}

type MockPackageClient struct {
	PackageData map[string][]byte
}

func (m *MockPackageClient) ConsumeByVechileID(ctx context.Context, vehichleID string) ([]byte, error) {
	data, found := m.PackageData[vehichleID]
	if !found {
		return nil, errors.New("Data not found")
	}
	return data, nil
}

func TestMockPackageInit_TrackByVechileID(t *testing.T) {
	mock := &MockPackageInit{
		PackageData: make(map[string]*Package),
	}

	// Create a test package
	testPackage := &Package{
		From:      "A",
		To:        "B",
		VehicleID: "123",
	}
	mock.PackageData["123"] = testPackage

	// Test case: Tracking an existing package
	result, err := mock.TrackByVechileID(context.Background(), "123")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != testPackage {
		t.Errorf("Expected the same package, but got a different one")
	}

	// Test case: Tracking a non-existent package
	_, err = mock.TrackByVechileID(context.Background(), "456")
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}

func TestMockPackageClient_ConsumeByVechileID(t *testing.T) {
	mock := &MockPackageClient{
		PackageData: make(map[string][]byte),
	}

	// Create test data
	testData := []byte("Test data")
	mock.PackageData["789"] = testData

	// Test case: Consuming data for an existing vehicle ID
	result, err := mock.ConsumeByVechileID(context.Background(), "789")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if string(result) != string(testData) {
		t.Errorf("Expected the same data, but got different data")
	}

	// Test case: Consuming data for a non-existent vehicle ID
	_, err = mock.ConsumeByVechileID(context.Background(), "999")
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}

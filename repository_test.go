package logitrack

import (
	"context"
	"testing"
)

var shipmentID string

func TestLogistics(t *testing.T) {
	ctx := context.Background()
	repo := NewPostgresRepository(testPool)

	MaliciousDispatchOrder := DispatchOrder{
		OriginWarehouseID:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
		DispatchWarehouseID: "b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22",
		TruckID:             "c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33",
		Items: []ShipmentItem{
			{
				InventoryID: "11111111-1111-1111-1111-111111111111",
				Quantity:    999,
			},
		},
	}

	_, err := repo.DispatchShipment(ctx, MaliciousDispatchOrder)
	if err == nil {
		t.Errorf("Expected an error because order quantity (999) exceeds available stock (50), but transaction succeeded.")
	}

	GoodDispatchOrder := DispatchOrder{
		OriginWarehouseID:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
		DispatchWarehouseID: "b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22",
		TruckID:             "c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33",
		Items: []ShipmentItem{
			{
				InventoryID: "11111111-1111-1111-1111-111111111111",
				Quantity:    50,
			},
		},
	}

	shipmentID, err = repo.DispatchShipment(ctx, GoodDispatchOrder)
	if err != nil {
		t.Errorf("Did not expect an error because quantity (50) is available but transaction failed.\nmessage:%v", err)
	}
}

func TestWarehouseInventory(t *testing.T) {
	ctx := context.Background()
	repo := NewPostgresRepository(testPool)

	items, err := repo.GetWarehouseInventory(ctx, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		t.Errorf("Warehouse inventory failed. Message: %v", err)
	}

	// I don't know how to compare two arrays of structs.
	found := false
	for _, item := range items {
		if item.ItemName == "MacBook Pro 14" && item.Quantity == 0 { // Since 50 Macbooks were shipped in the previous test
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Macbook Pro 14 not found in the returned items. Items:\n %v", items)
	}
}

func TestShipment(t *testing.T) {
	ctx := context.Background()
	repo := NewPostgresRepository(testPool)

	status, err := repo.GetShipment(ctx, shipmentID)
	if err != nil {
		t.Errorf("Shipment %v exists but is not found. Error: %v", shipmentID, err)
	}

	if status.TruckLicenseNumber != "DL-1CA-1234" {
		t.Errorf("License number expected DL-01-A-1234, found %v.", status.TruckLicenseNumber)
	}
}

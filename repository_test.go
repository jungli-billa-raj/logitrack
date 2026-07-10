package logitrack

import (
	"context"
	"testing"
)

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

	err := repo.DispatchShipment(ctx, MaliciousDispatchOrder)
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

	err = repo.DispatchShipment(ctx, GoodDispatchOrder)
	if err != nil {
		t.Errorf("Did not expect an error because quantity (50) is available but transaction failed.\nmessage:%v", err)
	}
}

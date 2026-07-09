package logitrack

import "context"

type ShipmentItem struct {
	InventoryID string
	Quantity    int
}

type DispatchOrder struct {
	OriginWarehouseID   string
	DispatchWarehouseID string
	TruckID             string
	Items               []ShipmentItem
}

type LogisticsRepository interface {
	DispatchShipment(cxt context.Context, order DispatchOrder) error
}

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

type InventoryItem struct {
	ItemName string
	Quantity int
}

type LogisticsRepository interface {
	DispatchShipment(cxt context.Context, order DispatchOrder) error
	GetWarehouseInventory(ctx context.Context, warehouseID string) ([]InventoryItem, error)
}

package logitrack

import "context"

type ShipmentItem struct {
	InventoryID string `json:"inventory_id"`
	Quantity    int    `json:"quantity"`
}

type DispatchOrder struct {
	OriginWarehouseID   string         `json:"origin_warehouse_id"`
	DispatchWarehouseID string         `json:"dispatch_warehouse_id"`
	TruckID             string         `json:"truck_id"`
	Items               []ShipmentItem `json:"items"`
}

type InventoryItem struct {
	ItemName string `json:"item_name"`
	Quantity int    `json:"quantity"`
}

type ShipmentStatus struct {
	OriginWarehouseID      string `json:"origin_warehouse_id"`
	DestinationWarehouseID string `json:"destination_warehouse_id"`
	TruckLicenseNumber     string `json:"truck_license_number"`
}
type LogisticsRepository interface {
	DispatchShipment(cxt context.Context, order DispatchOrder) (string, error)
	GetWarehouseInventory(ctx context.Context, warehouseID string) ([]InventoryItem, error)
	GetShipment(ctx context.Context, shipmentID string) (ShipmentStatus, error)
}

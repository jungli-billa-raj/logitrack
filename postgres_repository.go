package logitrack

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) DispatchShipment(ctx context.Context, order DispatchOrder) (string, error) {
	// tx is transaction block of pgx
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	defer tx.Rollback(ctx)

	// SQL queries go here

	// 1. loop through the items and deduct quantities from the inventory
	for _, item := range order.Items {
		query := `
		UPDATE inventory 
		SET quantity = quantity - $1
		WHERE id = $2 AND warehouse_id = $3`

		_, err := tx.Exec(ctx, query, item.Quantity, item.InventoryID, order.OriginWarehouseID)
		if err != nil {
			return "", err
		}
	}

	// 2. Update the truck status to BUSY
	truckQuery := `
	UPDATE trucks
	SET status = 'BUSY'
	WHERE id = $1`
	_, err = tx.Exec(ctx, truckQuery, order.TruckID)
	if err != nil {
		return "", err
	}

	// 5. Create the shipment master record and return its new ID
	var shipmentID string
	shipmentQuery := `
	INSERT INTO shipments (origin_warehouse_id, destination_warehouse_id, truck_id, status)
	VALUES ($1,$2,$3,'IN_TRANSIT') 
	RETURNING id`

	err = tx.QueryRow(ctx, shipmentQuery, order.OriginWarehouseID, order.DispatchWarehouseID, order.TruckID).Scan(&shipmentID)
	if err != nil {
		return "", err
	}

	// 6. Insert individual items into the shipment items junction table
	for _, item := range order.Items {
		itemQuery := `
		INSERT INTO shipment_items
		(shipment_id, inventory_id, quantity_shipped)
		VALUES ($1,$2,$3)`

		_, err = tx.Exec(ctx, itemQuery, shipmentID, item.InventoryID, item.Quantity)
		if err != nil {
			return "", err
		}
	}

	return shipmentID, tx.Commit(ctx)
}

// An endpoint to see what items are currently sitting in a specific warehouse
func (r *PostgresRepository) GetWarehouseInventory(ctx context.Context, warehouse_id string) ([]InventoryItem, error) {
	query := `
	SELECT item_name, quantity
	FROM inventory
	WHERE warehouse_id = $1`

	rows, err := r.pool.Query(ctx, query, warehouse_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // why? how can memeory leaks occur here?

	items := []InventoryItem{}

	for rows.Next() {
		var name string
		var qty int

		err = rows.Scan(&name, &qty)
		if err != nil {
			return nil, err
		}

		items = append(items, InventoryItem{ItemName: name, Quantity: qty})

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil

}

// An endpoint to fetch a single shipment's details, including the tracking truck's license plate.

func (r *PostgresRepository) GetShipment(ctx context.Context, shipmentID string) (ShipmentStatus, error) {
	query := `
		SELECT 
			s.origin_warehouse_id, 
			s.destination_warehouse_id, 
			t.license_plate
		FROM shipments s
		LEFT JOIN trucks t ON s.truck_id = t.id
		WHERE s.id = $1`

	var origin, destination, license string

	err := r.pool.QueryRow(ctx, query, shipmentID).Scan(&origin, &destination, &license)
	if err != nil {
		return ShipmentStatus{}, err
	}
	return ShipmentStatus{OriginWarehouseID: origin, DestinationWarehouseID: destination, TruckLicenseNumber: license}, nil
}

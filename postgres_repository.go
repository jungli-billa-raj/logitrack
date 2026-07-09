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

func (r *PostgresRepository) DispatchShipment(ctx context.Context, order DispatchOrder) error {
	// tx is transaction block of pgx
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	// SQL queries go here

	// 1. loop through the items and deduct quantities from the inventory
	for _, item := range order.Items {
		query := `
		UPDATE inventory 
		SET quantity = quantity = $1
		WHERE id = $2 AND warehouse_id = $3`

		_, err := tx.Exec(ctx, query, item.Quantity, item.InventoryID, order.OriginWarehouseID)
		if err != nil {
			return err
		}
	}

	// 2. Update the truck status to BUSY
	truckQuery := `
	UPDATE trucks
	SET status = 'BUSY'
	WHERE id = $1`
	_, err = tx.Exec(ctx, truckQuery, order.TruckID)
	if err != nil {
		return err
	}

	// 5. Create the shipment master record and return its new ID
	var shipmentID string
	shipmentQuery := `
	INSERT INTO shipments (origin_warehouse_id, destination_warehouse_id, truck_id, status)
	VALUES ($1,$2,$3,'IN_TRANSIT') 
	RETURNING id`

	err = tx.QueryRow(ctx, shipmentQuery, order.OriginWarehouseID, order.DispatchWarehouseID, order.TruckID).Scan(&shipmentID)
	if err != nil {
		return err
	}

	// 6. Insert individual items into the shipment items junction table
	for _, item := range order.Items {
		itemQuery := `
		INSERT INTO shipment_items
		(shipment_id, inventory_id, quantity_shipped)
		VALUES ($1,$2,$3)`

		_, err = tx.Exec(ctx, itemQuery, shipmentID, item.InventoryID, item.Quantity)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *PostgresRepository) GetWarehouseInventory(ctx context.Context, warehouse_id string) ([]InventoryItem, error) {
	// tx is transaction block of pgx
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	query := `
	SELECT item_name, quantity
	FROM inventory
	WHERE warehouse_id = $1`

	rows, err := tx.Query(ctx, query, warehouse_id)
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

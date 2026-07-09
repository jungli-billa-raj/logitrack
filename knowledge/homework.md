Implement these two additional functions inside your `PostgresRepository`.

---

### Homework Task 1: Get Warehouse Inventory

An endpoint to see what items are currently sitting in a specific warehouse. This forces you to handle returning multiple rows of data in Go.

**The Interface Method to add to `LogisticsRepository`:**

```go
GetWarehouseInventory(ctx context.Context, warehouseID string) ([]InventoryItem, error)

```

*(Note: Create a small `InventoryItem` struct that holds `ItemName` and `Quantity`).*

**Your implementation goal inside `postgres_repository.go`:**

* Use `r.pool.Query(ctx, "SELECT item_name, quantity FROM inventory WHERE warehouse_id = $1", warehouseID)` to get multiple rows back.
* Use a loop with `rows.Next()` and `rows.Scan()` to read each row into a slice of your struct.
* Remember to call `defer rows.Close()` right after checking the query error to prevent memory leaks!

---

### Homework Task 2: Get Shipment Details (The `JOIN` test)

An endpoint to fetch a single shipment's details, including the tracking truck's license plate. This forces you to use that `LEFT JOIN` we discussed earlier.

**The Interface Method to add to `LogisticsRepository`:**

```go
GetShipment(ctx context.Context, shipmentID string) (ShipmentStatus, error)

```

**Your implementation goal inside `postgres_repository.go`:**

* Write a SQL query using `QueryRow` that selects from `shipments`, but performs a `LEFT JOIN trucks ON shipments.truck_id = trucks.id`.
* Scan the shipment status *and* the truck's `license_plate` into a destination struct.

---

### How to approach this:

Give these two methods a shot in your code. Don't worry about the HTTP router or network side yet—just focus on writing the raw SQL execution and scanning rows into Go structs.

Once you have these written, we will kick off **Day 3** by writing the tests to prove all three of your repository methods work flawlessly! How does that sound for homework?
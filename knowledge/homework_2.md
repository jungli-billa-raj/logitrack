**Day 3 Homework**. Since you already wrote the logic for `GetWarehouseInventory` and `GetShipment`, your job now is to write the integration tests to verify they actually work against your seed data.

Add these two test scenarios inside your `repository_test.go` file right below your `GoodDispatchOrder` logic.

---

### Homework Test 1: Verify Warehouse Inventory Reading

Write a test case that calls your `GetWarehouseInventory` method to prove it correctly reads rows from the database.

**Your Implementation Guide:**

* Call `repo.GetWarehouseInventory(ctx, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")` (The Delhi Hub ID).
* Assert that `err` is `nil`.
* Check the slice of `InventoryItem` returned. Based on your `seed.sql`, you know exactly what should be inside.
* Write a loop or an explicit check to verify that at least one item matches `ItemName == "MacBook Pro 14"`.
* **Bonus Challenge:** Since your `GoodDispatchOrder` test case just ran right before this and deducted `50` items from stock, what should the expected quantity of MacBooks be now? Assert that the quantity matches that exact remaining value!

---

### Homework Test 2: Verify Shipment Tracking (`LEFT JOIN`)

Write a test case that calls your `GetShipment` method to prove your SQL join successfully stitches the shipment and truck data together.

**Your Implementation Guide:**

* In your `GoodDispatchOrder` test case, your code successfully inserted a brand-new shipment master record and returned a dynamic `shipmentID`.
* *Tip:* To test this, you’ll want to capture that newly generated `shipmentID` from your successful dispatch test case, rather than hardcoding a random UUID.
* Call `repo.GetShipment(ctx, shipmentID)`.
* Assert that `err` is `nil`.
* Check the returned `ShipmentStatus` struct. Verify that `TruckLicenseNumber` is exactly `"DL-01-A-1234"` (or whatever license plate matches the seeded truck you dispatched).

---

### How to execute when you get back:

Go get some coffee, stretch, and clear your head. When you come back, draft these two blocks inside `repository_test.go` and run your cache-busting command:

```bash
go test -v -count=1 .

```

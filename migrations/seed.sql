-- 1. SEED WAREHOUSES
INSERT INTO warehouses (id, name, location) VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Delhi Hub', 'Sector 5, Okhla'),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'Mumbai Port', 'JNPT, Navi Mumbai');

-- 2. SEED INVENTORY (Giving them explicit UUIDs so shipment_items can reference them)
INSERT INTO inventory (id, warehouse_id, item_name, quantity) VALUES
('11111111-1111-1111-1111-111111111111', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'MacBook Pro 14', 50),
('22222222-2222-2222-2222-222222222222', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'iPhone 15', 100),
('33333333-3333-3333-3333-333333333333', 'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'MacBook Pro 14', 10);

-- 3. SEED TRUCKS
INSERT INTO trucks (id, license_plate, status) VALUES
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 'DL-1CA-1234', 'BUSY'), -- Marked busy because it's currently delivering
('d0eebc99-9c0b-4ef8-bb6d-6bb9bd380a44', 'MH-02BD-5678', 'MAINTENANCE');

-- 4. SEED A SHIPMENT
-- An active shipment traveling from Delhi Hub to Mumbai Port using our busy truck
INSERT INTO shipments (id, origin_warehouse_id, destination_warehouse_id, truck_id, status) VALUES
('e0eebc99-9c0b-4ef8-bb6d-6bb9bd380a55', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 'IN_TRANSIT');

-- 5. SEED SHIPMENT ITEMS (The manifest details for the shipment above)
-- This shipment contains 5 MacBooks and 10 iPhones taken out of Delhi's inventory
INSERT INTO shipment_items (shipment_id, inventory_id, quantity_shipped) VALUES
('e0eebc99-9c0b-4ef8-bb6d-6bb9bd380a55', '11111111-1111-1111-1111-111111111111', 5),
('e0eebc99-9c0b-4ef8-bb6d-6bb9bd380a55', '22222222-2222-2222-2222-222222222222', 10);
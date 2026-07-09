INSERT INTO warehouses (id, name, location) VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Delhi Hub', 'Sector 5, Okhla'),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'Mumbai Port', 'JNPT, Navi Mumbai');

INSERT INTO inventory (warehouse_id, item_name, quantity) VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'MacBook Pro 14', 50),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'iPhone 15', 100),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'MacBook Pro 14', 10);

INSERT INTO trucks (id, license_plate, status) VALUES
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 'DL-1CA-1234', 'AVAILABLE'),
('d0eebc99-9c0b-4ef8-bb6d-6bb9bd380a44', 'MH-02BD-5678', 'MAINTENANCE');
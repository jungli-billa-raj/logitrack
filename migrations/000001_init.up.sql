-- Enable UUID extension (Postgres specific awesome feature)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 1. WAREHOUSES TABLE
CREATE TABLE warehouses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    location VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 2. INVENTORY TABLE
CREATE TABLE inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    warehouse_id UUID NOT NULL REFERENCES warehouses(id) ON DELETE CASCADE,
    item_name VARCHAR(100) NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0),
    UNIQUE(warehouse_id, item_name) -- Prevents duplicate rows for the same item in the same warehouse
);

-- 3. TRUCKS TABLE
CREATE TABLE trucks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    license_plate VARCHAR(20) UNIQUE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'AVAILABLE' CHECK (status IN ('AVAILABLE', 'BUSY', 'MAINTENANCE'))
);

-- 4. SHIPMENTS TABLE
CREATE TABLE shipments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    origin_warehouse_id UUID NOT NULL REFERENCES warehouses(id),
    destination_warehouse_id UUID NOT NULL REFERENCES warehouses(id),
    truck_id UUID REFERENCES trucks(id) ON DELETE SET NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'IN_TRANSIT', 'DELIVERED')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    -- Database rule: Can't ship to the exact same warehouse
    CONSTRAINT check_distinct_warehouses CHECK (origin_warehouse_id <> destination_warehouse_id)
);

-- 5. SHIPMENT ITEMS (Junction table for Many-to-Many relationship)
CREATE TABLE shipment_items (
    shipment_id UUID NOT NULL REFERENCES shipments(id) ON DELETE CASCADE,
    inventory_id UUID NOT NULL REFERENCES inventory(id) ON DELETE RESTRICT,
    quantity_shipped INT NOT NULL CHECK (quantity_shipped > 0),
    PRIMARY KEY (shipment_id, inventory_id)
);
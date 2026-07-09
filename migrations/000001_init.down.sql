-- Drop tables with foreign key dependencies first
DROP TABLE IF EXISTS shipments;
DROP TABLE IF EXISTS inventory;

-- Drop independent tables
DROP TABLE IF EXISTS trucks;
DROP TABLE IF EXISTS warehouses;

-- Optional: Drop the extension if it wasn't being used before
DROP EXTENSION IF EXISTS "uuid-ossp";
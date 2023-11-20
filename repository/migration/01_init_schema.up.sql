-- Active: 1699836715721@@127.0.0.1@5432@laundry

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE tbl_customers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    address VARCHAR(255),
    phone_number VARCHAR(255)
);


CREATE TABLE tbl_service_details (
    service_no SERIAL PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    units VARCHAR(255) NOT NULL,
    unit_price BIGINT NOT NULL
);

CREATE TABLE tbl_services (
    id SERIAL PRIMARY KEY,
    service_no INTEGER NOT NULL REFERENCES tbl_service_details (service_no),
    quantity INTEGER NOT NULL,
    subtotal BIGINT
);


CREATE TABLE tbl_bills (
    id SERIAL PRIMARY KEY,
    date_in TIMESTAMP NOT NULL,
    date_finished TIMESTAMP NOT NULL,
    customer_id UUID NOT NULL REFERENCES tbl_customers (id),
    services tbl_services[],
    bill_amount BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);




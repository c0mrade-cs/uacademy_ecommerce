CREATE TABLE "order"(
    id	CHAR(36) PRIMARY KEY,
    product_id CHAR(36),
    quantity INTEGER NOT NULL,
    customer_name VARCHAR(36) NOT NULL,
    customer_address VARCHAR(36) NOT NULL,
    customer_phone VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

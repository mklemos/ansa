CREATE TABLE customers (
    customer_id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    encrypted_payment_details TEXT,
    preloaded_balance FLOAT,
    loyalty_points INT
);

-- Create users table
CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Menu items table
CREATE TABLE menu_items
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price double precision NOT NULL,
    category VARCHAR(50),
    image_url TEXT,
    is_available BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Orders table
CREATE TABLE orders
(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) NOT NULL,
    total_amount double precision NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT valid_status CHECK (status IN ('pending', 'paid', 'preparing', 'ready', 'completed', 'cancelled'))
);

-- Order items table
CREATE TABLE order_items
(
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) NOT NULL,
    menu_item_id INTEGER REFERENCES menu_items(id) NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price double precision NOT NULL,
    subtotal double precision NOT NULL
);

-- PayPal transactions table
CREATE TABLE paypal_transactions
(
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) NOT NULL,
    paypal_transaction_id VARCHAR(100) NOT NULL UNIQUE,
    payment_status VARCHAR(50) NOT NULL,
    payment_amount double precision NOT NULL,
    currency VARCHAR(3) DEFAULT 'USD',
    payer_email VARCHAR(255),
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT valid_payment_status CHECK (payment_status IN ('pending', 'completed', 'failed', 'refunded'))
);

-- Create indexes for better query performance
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_paypal_transactions_order_id ON paypal_transactions(order_id);

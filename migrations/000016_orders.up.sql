CREATE TABLE "orders" (
    "id" SERIAL PRIMARY KEY,
    "order_id" UUID UNIQUE NOT NULL,
    "user_id" INT NOT NULL,
    "total_price" INT NOT NULL,
    "payment_status" VARCHAR(50) DEFAULT 'pending',
    "transaction_time" TIMESTAMP NOT NULL DEFAULT NOW()
);
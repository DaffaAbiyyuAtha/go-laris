CREATE TABLE "order_items" (
    "id" SERIAL PRIMARY KEY,
    "order_id" UUID NOT NULL REFERENCES "orders"("order_id"),
    "product_id" INT NOT NULL,
    "qty" INT NOT NULL,
    "price" INT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
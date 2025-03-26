CREATE TABLE "transaction"(
    "id" serial primary key,
    "date_transaction" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "total_transaction" BIGINT,
    "user_id" INT REFERENCES "user"("id"),
    "product_id" INT REFERENCES "product"("id")
);
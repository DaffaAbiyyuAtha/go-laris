CREATE TABLE "transaction"(
    "id" serial primary key,
    "date_transaction" TIMESTAMP,
    "product_id" INT REFERENCES "product"("id")
);
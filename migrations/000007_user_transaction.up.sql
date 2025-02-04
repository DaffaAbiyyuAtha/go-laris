CREATE TABLE "user_transaction"(
    "id" serial primary key,
    "user_id" INT REFERENCES "user"("id"),
    "transaction_id" INT REFERENCES "transaction"("id")
);
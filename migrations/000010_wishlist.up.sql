CREATE TABLE "wishlist"(
    "id"SERIAL PRIMARY KEY,
    "user_id" INT REFERENCES "user"("id"),
    "product_id" INT REFERENCES "product"("id")
)
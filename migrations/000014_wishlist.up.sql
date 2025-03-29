CREATE TABLE "wishlist"(
    "id"SERIAL PRIMARY KEY,
    "profile_id" INT REFERENCES "profile"("id") ON DELETE CASCADE,
    "product_id" INT REFERENCES "product"("id")
)
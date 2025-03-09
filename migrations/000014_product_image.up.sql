CREATE TABLE "product_images" (
    "id" SERIAL PRIMARY KEY,
    "product_id" INT REFERENCES "product"("id") ON DELETE CASCADE,
    "image" TEXT NOT NULL
);
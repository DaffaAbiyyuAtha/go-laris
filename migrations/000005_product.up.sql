CREATE TABLE "product"(
    "id" serial primary key,
    "name_product" VARCHAR(255),
    "price" INTEGER,
    "discount" INTEGER,
    "description" TEXT,
    "categories_id" INT REFERENCES "category"("id")
);
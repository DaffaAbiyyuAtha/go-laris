CREATE TABLE "product"(
    "id" serial primary key,
    "image" VARCHAR(255),
    "name_product" VARCHAR(255),
    "price" INTEGER,
    "discount" INTEGER,
    "total" INTEGER,
    "categories_id" INT REFERENCES "category"("id")
);


CREATE TABLE "profile"(
    "id" serial primary key,
    "picture" VARCHAR(255),
    "fullname" VARCHAR(255),
    "province" VARCHAR(255),
    "city" VARCHAR(255),
    "postal_code" INTEGER,
    "gender" INTEGER,
    "country" VARCHAR(50),
    "mobile" VARCHAR(50),
    "address" VARCHAR(255),
    "user_id" INT REFERENCES "user"("id")
);


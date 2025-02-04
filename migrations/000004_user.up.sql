CREATE TABLE "user"(
    "id" serial primary key,
    "email" VARCHAR(255),
    "password" VARCHAR(255),
    "role_id" INT REFERENCES "role"("id"),
    "profile_id" INT REFERENCES "profile"("id")
);
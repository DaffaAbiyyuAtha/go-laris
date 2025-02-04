CREATE TABLE "role"(
    "id" serial primary key,
    "owner" varchar(50),
    "admin" varchar(50),
    "user" varchar(50)
);
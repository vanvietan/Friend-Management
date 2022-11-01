CREATE TABLE "users" (
                         "id" bigint PRIMARY KEY NOT NULL,
                         "email" varchar(50) UNIQUE NOT NULL,
                         "created_at" timestamp,
                         "updated_at" timestamp
);

CREATE TABLE "relationship" (
                                "id" bigint PRIMARY KEY NOT NULL,
                                "requester_id" bigint NOT NULL,
                                "addressee_id" bigint NOT NULL,
                                "type" varchar(20),
                                "created_at" timestamp,
                                "updated_at" timestamp,

);

ALTER TABLE "relationship" ADD FOREIGN KEY ("requester_id") REFERENCES "users" ("id");

ALTER TABLE "relationship" ADD FOREIGN KEY ("addressee_id") REFERENCES "users" ("id");

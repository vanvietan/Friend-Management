CREATE TABLE "users" IF NOT EXISTS(
                         "id" bigint PRIMARY KEY NOT NULL,
                         "email" varchar(30) UNIQUE NOT NULL
);

CREATE TABLE "relationships" IF NOT EXISTS(
                                "id" bigint PRIMARY KEY NOT NULL,
                                "requester_id" bigint NOT NULL,
                                "addressee_id" bigint NOT NULL,
                                "type" varchar(20)
);

ALTER TABLE "relationships" ADD FOREIGN KEY ("requester_id") REFERENCES "users" ("id");

ALTER TABLE "relationships" ADD FOREIGN KEY ("addressee_id") REFERENCES "users" ("id");

-- +migrate Up
CREATE TABLE "users"
(
    "id"           SERIAL PRIMARY KEY,
    "user_name"    varchar(255) NOT NULL UNIQUE,
    "password"     varchar(255),
    "email"        varchar(255) UNIQUE,
    "phone_number" varchar(255) NOT NULL UNIQUE,
    "is_active"    bool      DEFAULT false,
    "active_code"  varchar(6)   NOT NULL,
    "exp_code"     timestamp    NOT NULL,
    "created_at"   timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at"   timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE vehicle_status AS ENUM ('New', 'Processing', 'Done');

CREATE TABLE "vehicle_orders"
(
    "id"         SERIAL PRIMARY KEY,
    "user_id"    bigint         NOT NULL,
    "handler_id" bigint,
    "status"     vehicle_status NOT NULL,
    "note"       text,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "history_vehicle_order_status"
(
    "id"               SERIAL PRIMARY KEY,
    "vehicle_order_id" bigint         NOT NULL,
    "status"           vehicle_status NOT NULL,
    "created_at"       timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at"       timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "items"
(
    "id"          SERIAL PRIMARY KEY,
    "category_id" bigint NOT NULL,
    "name"        text,
    "description" text,
    "price"       bigint    DEFAULT 0,
    "created_at"  timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "vehicle_order_items"
(
    "id"               SERIAL PRIMARY KEY,
    "vehicle_order_id" bigint NOT NULL,
    "item_id"          bigint,
    "note"             text,
    "price"            bigint    DEFAULT 0,
    "created_at"       timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at"       timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "categories"
(
    "id"         SERIAL PRIMARY KEY,
    "name"       varchar(255) NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "invoices"
(
    "id"               SERIAL PRIMARY KEY,
    "vehicle_order_id" bigint NOT NULL,
    "total"            bigint,
    "created_at"       timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at"       timestamp DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "vehicle_orders"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "vehicle_orders"
    ADD FOREIGN KEY ("handler_id") REFERENCES "users" ("id");

ALTER TABLE "history_vehicle_order_status"
    ADD FOREIGN KEY ("vehicle_order_id") REFERENCES "vehicle_orders" ("id");

ALTER TABLE "items"
    ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "vehicle_order_items"
    ADD FOREIGN KEY ("vehicle_order_id") REFERENCES "vehicle_orders" ("id");

ALTER TABLE "vehicle_order_items"
    ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "invoices"
    ADD FOREIGN KEY ("vehicle_order_id") REFERENCES "vehicle_orders" ("id");
-- +migrate Down

-- Drop foreign keys in reverse order of creation
ALTER TABLE "history_vehicle_order_status" DROP CONSTRAINT IF EXISTS "history_vehicle_order_status_vehicle_order_id_fkey";
ALTER TABLE "invoices" DROP CONSTRAINT IF EXISTS "invoices_vehicle_order_id_fkey";
ALTER TABLE "items" DROP CONSTRAINT IF EXISTS "items_category_id_fkey";
ALTER TABLE "vehicle_order_items" DROP CONSTRAINT IF EXISTS "vehicle_order_items_vehicle_order_id_fkey";
ALTER TABLE "vehicle_order_items" DROP CONSTRAINT IF EXISTS "vehicle_order_items_item_id_fkey";
ALTER TABLE "vehicle_orders" DROP CONSTRAINT IF EXISTS "vehicle_orders_user_id_fkey";
ALTER TABLE "vehicle_orders" DROP CONSTRAINT IF EXISTS "vehicle_orders_handler_id_fkey";

-- Drop tables in reverse order of creation
DROP TABLE "invoices";
DROP TABLE "vehicle_order_items";
DROP TABLE "items";
DROP TABLE "categories";

DROP TABLE "history_vehicle_order_status";
DROP TABLE "vehicle_orders";
DROP TYPE vehicle_status;
DROP TABLE "users";
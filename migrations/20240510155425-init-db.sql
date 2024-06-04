-- +migrate Up
CREATE TABLE "users"
(
    "id"           SERIAL PRIMARY KEY,
    "user_name"    varchar(255) NOT NULL UNIQUE,
    "password"     varchar(255),
    "email"        varchar(255) UNIQUE,
    "phone_number" varchar(255) NOT NULL UNIQUE,
    "is_active"    bool      DEFAULT false,
    "active_code"  varchar(6) NOT NULL,
    "exp_code"     timestamp NOT NULL,
    "created_at"   timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at"   timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "orders"
(
    "id"              SERIAL PRIMARY KEY,
    "order_user_id"   bigint NOT NULL,
    "handler_user_id" bigint,
    "status_id"       int       DEFAULT 1,
    "created_at"      timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at"      timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "status"
(
    "id"         SERIAL PRIMARY KEY,
    "name"       varchar(255) NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "history_order_status"
(
    "id"         SERIAL PRIMARY KEY,
    "order_id"   bigint NOT NULL,
    "status_id"  bigint NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
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

CREATE TABLE "order_items"
(
    "id"         SERIAL PRIMARY KEY,
    "order_id"   bigint NOT NULL,
    "item_id"    bigint,
    "note"       text,
    "price"      bigint    DEFAULT 0,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
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
    "id"         SERIAL PRIMARY KEY,
    "order_id"   bigint NOT NULL,
    "total"      bigint,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "orders"
    ADD FOREIGN KEY ("order_user_id") REFERENCES "users" ("id");

ALTER TABLE "orders"
    ADD FOREIGN KEY ("handler_user_id") REFERENCES "users" ("id");

ALTER TABLE "orders"
    ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");

ALTER TABLE "history_order_status"
    ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "history_order_status"
    ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");

ALTER TABLE "items"
    ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "order_items"
    ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items"
    ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "invoices"
    ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
-- +migrate Down

-- Drop foreign keys in reverse order of creation
ALTER TABLE "invoices" DROP CONSTRAINT "invoices_order_id_fkey";

ALTER TABLE "order_items" DROP CONSTRAINT "order_items_item_id_fkey";
ALTER TABLE "order_items" DROP CONSTRAINT "order_items_order_id_fkey";

ALTER TABLE "items" DROP CONSTRAINT "items_category_id_fkey";

ALTER TABLE "history_order_status" DROP CONSTRAINT "history_order_status_status_id_fkey";
ALTER TABLE "history_order_status" DROP CONSTRAINT "history_order_status_order_id_fkey";

ALTER TABLE "orders" DROP CONSTRAINT "orders_status_id_fkey";
ALTER TABLE "orders" DROP CONSTRAINT "orders_handler_user_id_fkey";
ALTER TABLE "orders" DROP CONSTRAINT "orders_order_user_id_fkey";

-- Drop tables in reverse order of creation
DROP TABLE "invoices";
DROP TABLE "order_items";
DROP TABLE "items";
DROP TABLE "categories";

DROP TABLE "status";
DROP TABLE "history_order_status";
DROP TABLE "orders";
DROP TABLE "users";
-- Database: local_dev
-- Generation Time: 2020-10-10 18:20:45.9510
-- -------------------------------------------------------------


CREATE SEQUENCE IF NOT EXISTS auctions_id_seq;

-- Table Definition
CREATE TABLE "public"."auctions" (
    "id" int8 NOT NULL DEFAULT nextval('auctions_id_seq'::regclass),
    "item_id" int8 NOT NULL,
    "increment_rate" int8 NOT NULL,
    "finished" bool NOT NULL DEFAULT false,
    "winner" int8 NOT NULL,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS bids_id_seq;

-- Table Definition
CREATE TABLE "public"."bids" (
    "id" int8 NOT NULL DEFAULT nextval('bids_id_seq'::regclass),
    "auction_id" int8 NOT NULL,
    "user_id" int8 NOT NULL,
    "initial" bool NOT NULL DEFAULT false,
    "amount" int8 NOT NULL,
    "created_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS categories_id_seq;

-- Table Definition
CREATE TABLE "public"."categories" (
    "id" int8 NOT NULL DEFAULT nextval('categories_id_seq'::regclass),
    "title" varchar NOT NULL,
    "description" text,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS items_id_seq;

-- Table Definition
CREATE TABLE "public"."items" (
    "id" int8 NOT NULL DEFAULT nextval('items_id_seq'::regclass),
    "title" varchar NOT NULL,
    "base_link" varchar NOT NULL,
    "ref" varchar,
    "description" text,
    "featured" bool NOT NULL DEFAULT false,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "ref_owner" int8,
    "featured_by" int8,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS notifications_id_seq;

-- Table Definition
CREATE TABLE "public"."notifications" (
    "id" int8 NOT NULL DEFAULT nextval('notifications_id_seq'::regclass),
    "user_id" int8 NOT NULL,
    "content" text NOT NULL,
    "created_at" timestamp NOT NULL,
    "viewed" bool NOT NULL DEFAULT false,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS ratings_id_seq;

-- Table Definition
CREATE TABLE "public"."ratings" (
    "id" int8 NOT NULL DEFAULT nextval('ratings_id_seq'::regclass),
    "item_id" int8 NOT NULL,
    "user_id" int8 NOT NULL,
    "rating" int2 NOT NULL,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS untitled_table_197_id_seq;

-- Table Definition
CREATE TABLE "public"."requests" (
    "id" int8 NOT NULL DEFAULT nextval('untitled_table_197_id_seq'::regclass),
    "type" varchar NOT NULL DEFAULT ''::character varying,
    "content" text,
    "user_id" int8 NOT NULL,
    "approved" bool NOT NULL DEFAULT false,
    "created_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "balance" int8,
    "created_at" timestamp NOT NULL,
    "activated" bool NOT NULL DEFAULT false,
    "last_ip" varchar NOT NULL,
    "admin" bool NOT NULL DEFAULT false,
    PRIMARY KEY ("id")
);


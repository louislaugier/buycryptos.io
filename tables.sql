CREATE SEQUENCE IF NOT EXISTS auctions_id_seq;

CREATE TABLE "public"."auctions" (
    "id" int8 NOT NULL DEFAULT nextval('auctions_id_seq'::regclass),
    "item_id" int8 NOT NULL,
    "increment_rate" int8 NOT NULL,
    "winner_user_id" int8 NOT NULL,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS bids_id_seq;

CREATE TABLE "public"."bids" (
    "id" int8 NOT NULL DEFAULT nextval('bids_id_seq'::regclass),
    "auction_id" int8 NOT NULL,
    "user_id" int8 NOT NULL,
    "is_initial" bool NOT NULL DEFAULT false,
    "amount" int8 NOT NULL,
    "created_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS categories_id_seq;

CREATE TABLE "public"."categories" (
    "id" int8 NOT NULL DEFAULT nextval('categories_id_seq'::regclass),
    "title" varchar NOT NULL,
    "description" text,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS ratings_id_seq;

CREATE TABLE "public"."comments" (
    "id" int8 NOT NULL DEFAULT nextval('ratings_id_seq'::regclass),
    "item_id" int8 NOT NULL,
    "user_id" int8 NOT NULL,
    "rating" int2 NOT NULL,
    "content" text,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS featurings_id_seq;

CREATE TABLE "public"."featurings" (
    "id" int8 NOT NULL DEFAULT nextval('featurings_id_seq'::regclass),
    "user_id" int8 NOT NULL,
    "item_id" int8 NOT NULL,
    "created_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS fundings_id_seq;

CREATE TABLE "public"."fundings" (
    "id" int8 NOT NULL DEFAULT nextval('fundings_id_seq'::regclass),
    "user_id" int8 NOT NULL,
    "amount" int8 NOT NULL,
    "payment_method" varchar NOT NULL,
    "created_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS items_id_seq;

CREATE TABLE "public"."items" (
    "id" int8 NOT NULL DEFAULT nextval('items_id_seq'::regclass),
    "title" varchar NOT NULL,
    "base_link" varchar NOT NULL,
    "ref_link" varchar,
    "description" text,
    "is_featured" bool NOT NULL DEFAULT false,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "ref_link_owner_user_id" int8,
    "featurer_user_id" int8,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS notifications_id_seq;

CREATE TABLE "public"."notifications" (
    "id" int8 NOT NULL DEFAULT nextval('notifications_id_seq'::regclass),
    "user_id" int8 NOT NULL,
    "content" text NOT NULL,
    "created_at" timestamp NOT NULL,
    "is_read" bool NOT NULL DEFAULT false,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS requests_id_seq;

CREATE TABLE "public"."requests" (
    "id" int8 NOT NULL DEFAULT nextval('requests_id_seq'::regclass),
    "action_type" varchar NOT NULL DEFAULT ''::character varying,
    "title" varchar,
    "base_link" varchar,
    "description" text,
    "comment" text,
    "rating" int2,
    "item_id" int8,
    "user_id" int8 NOT NULL,
    "is_approved" bool NOT NULL DEFAULT false,
    "created_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS users_id_seq;

CREATE TABLE "public"."users" (
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "balance" int8,
    "created_at" timestamp NOT NULL,
    "is_activated" bool NOT NULL DEFAULT false,
    "last_ip" varchar NOT NULL,
    "is_admin" bool NOT NULL DEFAULT false,
    "is_deleted" bool NOT NULL DEFAULT false,
    PRIMARY KEY ("id")
);


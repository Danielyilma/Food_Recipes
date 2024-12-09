alter table "public"."likes" alter column "rating" drop not null;
alter table "public"."likes" add column "rating" int4;

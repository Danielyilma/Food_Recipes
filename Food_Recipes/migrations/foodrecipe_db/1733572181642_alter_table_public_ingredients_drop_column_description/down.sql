alter table "public"."ingredients" alter column "description" drop not null;
alter table "public"."ingredients" add column "description" text;

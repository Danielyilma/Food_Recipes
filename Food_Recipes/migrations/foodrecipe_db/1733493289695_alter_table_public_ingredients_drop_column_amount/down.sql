alter table "public"."ingredients" alter column "amount" drop not null;
alter table "public"."ingredients" add column "amount" text;

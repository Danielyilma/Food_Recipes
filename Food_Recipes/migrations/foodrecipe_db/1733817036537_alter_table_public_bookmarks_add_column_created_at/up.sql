alter table "public"."bookmarks" add column "created_at" timestamptz
 null default now();

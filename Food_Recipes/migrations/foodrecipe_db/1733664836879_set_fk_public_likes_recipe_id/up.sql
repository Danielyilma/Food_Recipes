alter table "public"."likes"
  add constraint "likes_recipe_id_fkey"
  foreign key ("recipe_id")
  references "public"."recipes"
  ("id") on update restrict on delete cascade;

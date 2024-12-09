alter table "public"."recipe_ingredients" drop constraint "recipe_ingredients_recipe_id_fkey",
  add constraint "recipe_ingredients_recipe_id_fkey"
  foreign key ("recipe_id")
  references "public"."recipes"
  ("id") on update restrict on delete restrict;

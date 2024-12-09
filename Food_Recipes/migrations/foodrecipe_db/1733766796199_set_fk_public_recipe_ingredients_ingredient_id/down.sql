alter table "public"."recipe_ingredients" drop constraint "recipe_ingredients_ingredient_id_fkey",
  add constraint "recipe_ingredients_ingredient_id_fkey"
  foreign key ("ingredient_id")
  references "public"."ingredients"
  ("id") on update restrict on delete restrict;

alter table "public"."recipe_images"
  add constraint "recipe_images_step_id_fkey"
  foreign key ("step_id")
  references "public"."steps"
  ("id") on update restrict on delete cascade;

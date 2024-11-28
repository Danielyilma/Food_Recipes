CREATE TABLE "public"."recipes" ("id" serial NOT NULL, "title" text NOT NULL, "description" text NOT NULL, "prep_time" integer NOT NULL, "thumbnail" text NOT NULL, "is_paid" boolean NOT NULL, "price" numeric NOT NULL DEFAULT 0.00, "user_id" integer NOT NULL, "category_id" integer NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") , FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON UPDATE restrict ON DELETE cascade, FOREIGN KEY ("category_id") REFERENCES "public"."categories"("id") ON UPDATE restrict ON DELETE restrict);
CREATE OR REPLACE FUNCTION "public"."set_current_timestamp_updated_at"()
RETURNS TRIGGER AS $$
DECLARE
  _new record;
BEGIN
  _new := NEW;
  _new."updated_at" = NOW();
  RETURN _new;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER "set_public_recipes_updated_at"
BEFORE UPDATE ON "public"."recipes"
FOR EACH ROW
EXECUTE PROCEDURE "public"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_public_recipes_updated_at" ON "public"."recipes"
IS 'trigger to set value of column "updated_at" to current timestamp on row update';

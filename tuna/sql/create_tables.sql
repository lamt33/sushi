create table IF NOT EXISTS public.user (
user_id uuid NOT NULL DEFAULT gen_random_uuid(),
user_name varchar NOT NULL DEFAULT '',
PRIMARY KEY (user_id)
);
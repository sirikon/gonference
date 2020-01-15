CREATE TABLE "public"."user" (
    username VARCHAR (200) NOT NULL PRIMARY KEY,
    password VARCHAR (256) NOT NULL
);
INSERT INTO "public"."user" VALUES ('admin', 'jGl25bVBBBW96Qi9Te4V37Fnqchz_Eu4qB9vKrRIqRg=');

CREATE TABLE "public"."talk" (
    id serial PRIMARY KEY,
    name VARCHAR (200) NOT NULL,
    description TEXT NOT NULL,
    speaker_name VARCHAR (200) NOT NULL,
    speaker_title VARCHAR (200) NOT NULL,
    track VARCHAR (50) NOT NULL,
    when_date TIMESTAMP NOT NULL
);

CREATE TABLE "public"."rating" (
    id serial PRIMARY KEY,
    talk_id integer NOT NULL REFERENCES "public"."talk" (id) ON DELETE CASCADE,
    visitor_key uuid NOT NULL,
    stars smallint NOT NULL,
    comment varchar(600),
    UNIQUE (talk_id, visitor_key)
);

CREATE TABLE "public"."question" (
    id serial PRIMARY KEY,
    talk_id integer NOT NULL REFERENCES "public"."talk" (id) ON DELETE CASCADE,
    visitor_key uuid NOT NULL,
    question varchar(600)
);

CREATE TABLE talk (
    id serial PRIMARY KEY,
    name VARCHAR (200) NOT NULL,
    description TEXT NOT NULL,
    speaker_name VARCHAR (200) NOT NULL,
    speaker_title VARCHAR (200) NOT NULL,
    track VARCHAR (50) NOT NULL,
    when_date TIMESTAMP NOT NULL
);

CREATE TABLE "user" (
    username VARCHAR (200) NOT NULL PRIMARY KEY,
    password VARCHAR (256) NOT NULL
);

INSERT INTO "user" VALUES ('admin', 'jGl25bVBBBW96Qi9Te4V37Fnqchz_Eu4qB9vKrRIqRg=');

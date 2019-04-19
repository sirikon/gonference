CREATE TABLE talk (
    id serial PRIMARY KEY,
    name VARCHAR (200) NOT NULL,
    description TEXT NOT NULL,
    speaker_name VARCHAR (200) NOT NULL,
    speaker_title VARCHAR (200) NOT NULL,
    track VARCHAR (50) NOT NULL,
    when_date VARCHAR (50) NOT NULL
);

CREATE TABLE users (
    id serial PRIMARY KEY not null,
    username varchar not null,
    twitch_id varchar UNIQUE not null,
    image_link varchar not null,
    registration_date varchar not null
);
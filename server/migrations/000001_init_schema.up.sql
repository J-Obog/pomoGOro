CREATE TABLE users (
    id SERIAL PRIMARY KEY, 
    email VARCHAR NOT NULL, 
    password VARCHAR NOT NULL
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    title VARCHAR NOT NULL
);
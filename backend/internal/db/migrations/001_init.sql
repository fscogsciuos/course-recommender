CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    text text NOT NULL,
    done BOOLEAN NOT NULL DEFAULT false
);

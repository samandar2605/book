CREATE TABLE if not exists books(
    id serial primary key,
    title varchar NOT NULL,
    author_name varchar NOT NULL,
    price decimal(18,2) NOT NULL,
    amount int,
    created_at timestamp default current_timestamp
);
CREATE TABLE burgers (
   id BIGSERIAL PRIMARY KEY,
   name TEXT NOT NULL,
   price INT NOT NULL CHECK ( price > 0 ),
   -- don't remove any data from db
   removed BOOLEAN DEFAULT FALSE
);


CREATE TABLE IF NOT EXISTS burgers
(
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    price       INT  NOT NULL CHECK ( price > 0 ),
    description TEXT NOT NULL,
    removed     BOOLEAN DEFAULT FALSE
);

DROP TABLE burgers;

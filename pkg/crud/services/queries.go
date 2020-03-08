package services

const BurgersDDL = `CREATE TABLE IF NOT EXISTS burgers
(
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    price       INT  NOT NULL CHECK ( price > 0 ),
    description TEXT NOT NULL,
    removed     BOOLEAN DEFAULT FALSE
);`

const GetBurgers = "SELECT id, name, price/100, description FROM burgers WHERE removed = FALSE;"

const SaveBurger = "INSERT INTO burgers (name, price, description) VALUES ($1, $2, $3);"

const RemoveBurger = "UPDATE burgers SET removed = TRUE WHERE id = $1;"

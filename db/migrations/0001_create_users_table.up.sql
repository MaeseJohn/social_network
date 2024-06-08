CREATE TABLE
    users (
        id uuid PRIMARY KEY,
        name TEXT  NOT NULL,
        last_name TEXT  NOT NULL,
        email TEXT NOT NULL,
        password TEXT  NOT NULL,
        age date NOT NULL
    );
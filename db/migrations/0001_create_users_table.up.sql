CREATE TABLE users (
        user_id uuid PRIMARY KEY,
        name TEXT  NOT NULL,
        last_name TEXT  NOT NULL,
        email TEXT NOT NULL UNIQUE,
        password TEXT  NOT NULL,
        age date NOT NULL,
        private BOOLEAN NOT NULL
    );
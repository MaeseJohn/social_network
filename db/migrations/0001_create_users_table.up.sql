CREATE TABLE
    users (
        id uuid PRIMARY KEY,
        name varchar(50)  NOT NULL,
        lat_name varchar(50)  NOT NULL,
        email varchar(100) NOT NULL,
        password varchar(100)  NOT NULL,
        age date NOT NULL
    );
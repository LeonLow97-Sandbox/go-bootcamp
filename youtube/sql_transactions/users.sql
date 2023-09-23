CREATE TABLE users (
  name VARCHAR NOT NULL,
  birth_year SMALLINT NULL DEFAULT 0
);

INSERT INTO users(name, birth_year) VALUES('Leon', 1997);
-- underscore necessary as user is a reserved keyword
CREATE TABLE
  User_ (
    id serial PRIMARY KEY,
    username varchar(34) UNIQUE NOT NULL,
    password_hash_salted varchar(72) NOT NULL
  );

CREATE TABLE
  Message (
    id serial PRIMARY KEY,
    content text NOT NULL,
    sender int,
    sender_name varchar(34),
    timestamp timestamp NOT NULL,
    FOREIGN KEY (sender) REFERENCES User_ (id)
  );

INSERT INTO
  User_ (username, password_hash_salted)
VALUES
  ('glumanda', 'not_a_hash');

INSERT INTO
  Message (content, sender, sender_name, timestamp)
VALUES
  (
    'init',
    (
      SELECT
        id
      FROM
        User_
      WHERE
        username = 'glumanda'
    ),
    'glumanda',
    CURRENT_TIMESTAMP
  );

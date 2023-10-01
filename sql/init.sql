CREATE TABLE message (
  id          serial    PRIMARY KEY,
  content     text      NOT NULL,
  sender      text      NOT NULL,
  timestamp   timestamp NOT NULL
);

INSERT INTO message (content, sender, timestamp) VALUES ('init_value', 'admin', CURRENT_TIMESTAMP);

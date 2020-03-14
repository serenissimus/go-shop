DROP TABLE IF EXISTS technologies;
CREATE TABLE technologies(
  id              SERIAL PRIMARY KEY,
  name            VARCHAR(32) NOT NULL,
  details         VARCHAR(4096) NULL
);
INSERT INTO technologies(name, details) VALUES('golang', 'Go language');
INSERT INTO technologies(name, details) VALUES('react', 'React');

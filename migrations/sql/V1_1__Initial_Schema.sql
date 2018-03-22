CREATE TABLE states (
    id serial PRIMARY KEY,
    name varchar(50) NULL
);

CREATE TABLE national_parks
(
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    date_designated timestamp NOT NULL,
    state_id integer NULL
);

ALTER TABLE national_parks
ADD CONSTRAINT fk_national_parks_states
FOREIGN KEY (state_id)
REFERENCES states(id)
ON DELETE CASCADE;

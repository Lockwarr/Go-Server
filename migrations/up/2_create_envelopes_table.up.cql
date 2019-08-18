CREATE TABLE IF NOT EXISTS envelopes (
    id uuid,
    text text,
    gesmes text,
    xmlns text,
    subject text,
    sender frozen<sender>,
    cube frozen<cube>,
    PRIMARY KEY (id)
);
CREATE TYPE IF NOT EXISTS sender (
    text text,
    name text
);

CREATE TYPE IF NOT EXISTS cube3 (
    text text,
    currency text,
    rate text
);

CREATE TYPE IF NOT EXISTS cube2 (
    text text,
    time text,
    cube3 list<frozen<cube3>>
);

CREATE TYPE IF NOT EXISTS cube (
    text text,
    cube2 list<frozen<cube2>>
);





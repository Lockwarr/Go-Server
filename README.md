# Go-Server
I am using windows docker. Installed docker and then entered the following commands to create new container with ScyllaDB and enter into it's cqlsh

- >docker run -p 9042:9042 --rm --name scylladb -d scylladb/scylla
- >docker exec -it scylladb cqlsh

In order to create new keyspace, user defined types and table - copy and paste the following code into the cqlsh ( in powershell I could paste in cqlsh with right click) 

`CREATE KEYSPACE denislav WITH REPLICATION = {'class' : 'SimpleStrategy','replication_factor' : 1};`

`use denislav;`

`CREATE TYPE IF NOT EXISTS sender (
    text text,
    name text
);`


`CREATE TYPE IF NOT EXISTS cube3 (
    text text,
    currency text,
    rate text
);`

`CREATE TYPE IF NOT EXISTS cube2 (
    text text,
    time text,
    cube3 list<frozen<cube3>>
);`

`CREATE TYPE IF NOT EXISTS cube (
    text text,
    cube2 list<frozen<cube2>>
);`
`CREATE TABLE IF NOT EXISTS envelopes (
    id uuid,
    text text,
    gesmes text,
    xmlns text,
    subject text,
    sender frozen<sender>,
    cube frozen<cube>,
    PRIMARY KEY (id)
);`


Navigate to cmd/server and enter the following commands in the console
 
`go build`

`./server`



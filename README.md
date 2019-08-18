# Go-Server
docker pull cassandra
mkdir ucs-cassandra-db

docker run -p 9042:9042 --rm --name cassandra -d cassandra:3.11
docker exec -it cassandra /bin/bash
cqlsh
CREATE KEYSPACE denislav WITH REPLICATION = {'class' : 'SimpleStrategy','replication_factor' : 1};
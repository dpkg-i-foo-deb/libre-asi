#!/bin/sh
podman stop libre-asi-database

podman rm libre-asi-database

podman pull docker.io/postgres

podman run -it --network libre-asi -d -p 5432:5432 --name libre-asi-database -e POSTGRES_PASSWORD="postgres" postgres

sleep 5

podman cp ../../database/scripts/0-create-user-database.sql libre-asi-database:/0-create-user-database.sql

podman exec -it libre-asi-database psql -U postgres -f 0-create-user-database.sql

#More stuffies will be added here later

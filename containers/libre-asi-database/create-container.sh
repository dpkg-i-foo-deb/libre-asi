#!/bin/bash
podman pull docker.io/postgres

podman run -it --network libre-asi -d -p 5432:5432 --name libre-asi-database -e POSTGRES_PASSWORD="postgres" postgres

sleep 5

podman cp ../../database/scripts/0-create-user-database.sql libre-asi-database:/0-create-user-database.sql

podman exec -it libre-asi-database psql -U postgres -f 0-create-user-database.sql

podman cp ../../database/scripts/1-create-database.sql libre-asi-database:/1-create-database.sql

podman exec -it libre-asi-database psql -U libre_asi -d libre_asi -f 1-create-database.sql

podman stop libre-asi-database

podman rm libre-asi-database

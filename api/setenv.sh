#!/bin/sh
echo "Setting up environment variables"
export AUTH_KEY="owo"
export CORS_ORIGINS="*"
export API_PORT=":3000"
export DB_CONNECTION="host=192.168.122.1 user=libre_asi password=libre_asi dbname=libre_asi port=5432 sslmode=disable TimeZone=America/Bogota"
export PROD="false"
export WORLD="false"
echo "Setup complete"

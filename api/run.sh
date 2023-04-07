#!/bin/bash

# Check if the argument is dev or prod
if [ "$1" == "dev" ]; then
  # Export variables for dev environment
  export AUTH_KEY="owo"
  export CORS_ORIGINS="https://localhost:5173"
  export API_PORT=":3000"
  export PROD="false"
  export WORLD="false"
elif [ "$1" == "prod" ]; then
  # No variables needed for prod environment
  :
else
  # Invalid argument, exit with error message
  echo "Usage: $0 dev|prod"
  exit 1
fi

# Run the command
go run ./cmd/libre-asi-api/ serve
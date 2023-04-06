#!/bin/bash

# check the argument
if [ "$1" == "dev" ]; then
  # set the prerender option to false
  export VITE_PRERENDER=false
  # check if the files cert.perm and key.perm exist in the current directory
  if [ ! -f cert.pem ] || [ ! -f key.pem ]; then
    # run the openssl command to generate them
    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
  fi
elif [ "$1" == "prod" ]; then
  # set the prerender option to true
  export VITE_PRERENDER=true
else
  # print an error message and exit
  echo "Invalid argument. Please use dev or prod."
  exit 1
fi

# run the commands
npm i
npm run dev
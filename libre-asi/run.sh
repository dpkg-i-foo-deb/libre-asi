#!/bin/bash

# check the argument
if [ "$1" == "dev" ]; then
  # set the prerender option to false
  export VITE_PRERENDER=false
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
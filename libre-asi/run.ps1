# check the argument
switch ($args[0]) {
  "dev" {
    # set the prerender option to false
    $env:VITE_PRERENDER = "false"
  }
  "prod" {
    # set the prerender option to true
    $env:VITE_PRERENDER = "true"
  }
  default {
    # print an error message and exit
    Write-Host "Invalid argument. Please use dev or prod."
    exit 1
  }
}

# run the commands
npm i
npm run dev
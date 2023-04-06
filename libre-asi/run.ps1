# check the argument
if ($args[0] -eq "dev") {
  # set the prerender option to false
  $env:VITE_PRERENDER = "false"
  # check if the files cert.perm and key.perm exist in the current directory
  if (-not (Test-Path cert.pem) -or -not (Test-Path key.pem)) {
    # run the openssl command to generate them
    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
  }
}
elseif ($args[0] -eq "prod") {
  # set the prerender option to true
  $env:VITE_PRERENDER = "true"
}
else {
  # print an error message and exit
  Write-Host "Invalid argument. Please use dev or prod."
  exit 1
}

# run the commands
npm i
npm run dev
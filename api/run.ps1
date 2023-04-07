# Check if the argument is dev or prod
if ($args[0] -eq "dev") {
  # Export variables for dev environment
  $env:AUTH_KEY = "owo"
  $env:CORS_ORIGINS = "https://localhost:5173"
  $env:API_PORT = ":3000"
  $env:PROD = "false"
  $env:WORLD = "false"
}
elseif ($args[0] -eq "prod") {
  # No variables needed for prod environment
  :
}
else {
  # Invalid argument, exit with error message
  Write-Host "Usage: $PSCommandPath dev|prod"
  exit 1
}

# Run the command
go run ./cmd/libre-asi-api/ serve
Write-Host "Setting environment variables..."

$env:AUTH_KEY="owo";
$env:CORS_ORIGINS="*";
$env:API_PORT=":3000";
$env:DB_CONNECTION="";
$env:PROD="false";
$env:WORLD="false";

Write-Host "Setup complete"

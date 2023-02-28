# Libre-ASI API

# How do I run this project?
`go run . serve`

This API uses...

## GO
The language that supports everything

## GORM
An awesome ORM for GO, used to create every model and easily map them to a database for queries and insertions

## GO/JWT
To generate JSON Web Tokens for user sessions

## Fiber
An awesome framework to create APIs

# Environment Variables
Plenty of things this API uses require environment variables, the setenv.sh file can give you an example of them

## CORS_ORIGINS
Specifies the origins that should be allowed to access the API through a web browser
### Example
`"*"`

## API_PORT
Port in which the API will listen for connections
### Example
`:3000`

## DB_CONNECTION
Connection string used for the database, you can refer to the GORM documentation for more examples
### Example
`host=libre-asi-database user=libre_asi password=libre_asi dbname=libre_asi port=5432 sslmode=disable TimeZone=America/Bogota`

## ADMIN_EMAIL
The default administrator account email
### Example
`admin@example.com`

## ADMIN_USERNAME
The default administrator account username
### Example
`admin`

## ADMIN_PASSWORD
The default administrator account password (Will be stored as a hash)
### Example
`admin`

# Endpoints
Most endpoints require a user to be authenticated in the first place, here's an example on how to log in, there is a file called insomnia.json, you can use it with the respective tool to try out the endpoints and see the examples in a more interactive way

## Log in as administrator
### Route 
`http://localhost:3000/login/administrator`
### Method
POST
### Required data
- email
- password

### Example

`{
	"email":"admin@example.com",
	"password":"admin"
}`

## Log in as interviewer
### Route
`http://localhost:3000/login/interviewer`

### Method
POST

### Required data
- email
- password

### Example
`
{
	"email":"interviewer@example.com",
	"password":"interviewer"
}`

## Sign Out
### Route
`http://localhost:3000/sign-out`
### Method
POST
### Required data
- Valid Session

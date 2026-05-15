# User
POST /users/register

### In this way, run the curl command in unix base Terminal
curl -X POST http://localhost:8080/users/register \
-H "Content-Type: application/json" \
-d '{"name": "Amirhossein", "phoneNumber": "09123456789"}'


### In this way, run the curl command in PowerShell.
curl.exe -X POST http://localhost:8080/users/register `
-H "Content-Type: application/json" `
-d '{\"name\": \"Amirhossein\", \"phoneNumber\":\"09123456789\"}'

### In this way, run the curl command in Cmd.
curl -X POST http://localhost:8080/users/register -H "Content-Type: application/json" -d "{\"name\": \"Amirhossein\", \"phoneNumber\":\"09123456789\"}"

#### OR

curl -X POST http://localhost:8080/users/register ^
-H "Content-Type: application/json" ^
-d "{\"name\": \"Amirhossein\", \"phoneNumber\":\"09123456789\"}"
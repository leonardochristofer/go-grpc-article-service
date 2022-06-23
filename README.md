# Article Service
This service contains:
1. Get Article By Id API
- `http://localhost:9100/article?id=###`
2. Get List Article API
- `http://localhost:9100/articles?offset=0&limit=10&search=###&author=###`
3. Insert Article API
- `http://localhost:9100/articles`
# Prerequisite
This service uses Go, gRPC, and PostgreSQL. I assume that you have everything installed on your local machine.
# Steps
1. Clone this repository.
2. Import .sql file from database folder into your local machine PostgreSQL database.
3. Use `make` commands to run services.
4. If you are using Postman to test an API, refer to postman folder to get sample request or import .json file into your Postman collection.
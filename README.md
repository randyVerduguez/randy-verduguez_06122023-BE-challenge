# Go Weather Restful API Service

This service relies on [weatherapi.com](https://www.weatherapi.com/) to return the current forecast of a particular place given its city, region, and country. 

## Run Instructions
This service relies on docker compose to run with all its dependencies. To run this service, first you must create an .env file in the project's directory with the following fields

``DATABASE_HOST=postgres_container
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=weather
SERVER_PORT=8080
API_KEY=...``

Create a free account at [weatherapi.com](https://www.weatherapi.com/) to generate an API_KEY.

To verify service is running correctly do the following

``curl -X GET localhost:8080/weather/welcome``
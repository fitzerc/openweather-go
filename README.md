# openweather-go
REST API that takes a latitude and longitude, queries the Open Weather API, and returns a very concise summary.

An Open Weather API key is required to use openweather-go. You can register at https://openweathermap.org to get a key.

To run openweather-go:
1. Update the appid in Config -> config.json with your Open Weather API key
2. (Optional) Change the port openweather-go runs on.
3. Run from the root directory: go run main.go

Example request: http://localhost:8080/weather?lat=33.44&lon=-96.731667

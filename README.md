# Weather Web

## About
Weather Web is a web application built with Go (Golang) using the Gin Framework that displays real-time weather data.  
The app retrieves data from an external API and presents it using a modular code structure, making it easy to develop, test, and maintain.  
The purpose of this project is to provide a simple example of a backend API implementation using Go and Gin.

## Features
- Retrieve weather data by city name
- Use Gin Framework for routing and request handling
- Modular project structure (controllers, models, utils)
- Environment configuration with `.env` file
- Fast and lightweight performance

## Project Structure
```
weatherweb/
├── controllers/    
├── models/         
│   └── weather.go
├── utils/          
│   └── config.go
├── tmp/            
├── .env            
├── go.mod          
├── go.sum          
└── main.go         
```

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/afdalRafi3107/weather-web.git
   cd weather-web
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the root directory and add the following configuration:
   ```
   API_KEY=your_api_key_here
   PORT=8080
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

5. Access the app through:
   ```
   http://localhost:8080
   ```

## API Endpoints
| Method | Endpoint                | Description                             |
|--------|--------------------------|-----------------------------------------|
| GET    | `/weather?city=Jakarta` | Get weather data based on city name     |

**Example Request:**
```
GET http://localhost:8080/weather?city=Jakarta
```

**Example Response:**
```json
{
  "city": "Jakarta",
  "temperature": "31°C",
  "condition": "Partly Cloudy",
  "wind_speed": "4.5 m/s",
  "humidity": "68%"
}
```

## Tech Stack
- Go (Golang)
- Gin Web Framework
- dotenv
- OpenWeatherMap API
- JSON REST API

## Usage Example
1. Run the Go server  
2. Open your browser or use a tool like Postman  
3. Access the endpoint:
   ```
   http://localhost:8080/weather?city=Jakarta
   ```
4. View the weather data in JSON format

## Contact
**Author:** Muhammad Afdal Rafi  
**Email:** afdalrafi3107@gmail.com  
**LinkedIn:** [linkedin.com/in/afdalrafi3107](https://www.linkedin.com/in/afdalrafi3107)  
**GitHub:** [github.com/afdalRafi3107](https://github.com/afdalRafi3107)

## License
This project is licensed under the MIT License.  
You are free to use, modify, and distribute it for personal or educational purposes.

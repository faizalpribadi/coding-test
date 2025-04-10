# InterOpera Coding Test API

A RESTful API server built with Go, Fiber, and Swagger documentation.

## Overview

This project is a Go-based API server that provides endpoints for sales data and AI-powered responses. It uses the Fiber web framework for handling HTTP requests, Cobra for CLI commands, and Swagger for API documentation.

## Features

- **Sales Data API**: Endpoints to retrieve and filter sales representatives and their deals
- **AI Integration**: Integration with Google's Gemini API for AI-powered responses
- **Swagger Documentation**: Interactive API documentation
- **Configuration Management**: YAML-based configuration
- **CLI Commands**: Command-line interface for starting the server

## Prerequisites

- Go 1.23 or higher
- API key for Google Gemini (if using AI features)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/interopera-coding-test.git
   cd interopera-coding-test/backend/golang
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `config.yaml` file in the project root with the following structure:
   ```yaml
   app:
     port: 8000
   ai_provider:
     gemini:
       api_key: "your-gemini-api-key"
       model: "gemini-pro"
   ```

## Running the Application

Start the server:

```bash
go run main.go server
```

The server will start on the port specified in your config file (default: 8000).

## API Documentation

Once the server is running, you can access the Swagger documentation at:

```
http://localhost:8000/docs/
```

## API Endpoints

### Sales API

- `GET /api/sales/data` - Get all sales data
- `GET /api/sales/sales-reps` - Get all sales representatives
- `GET /api/sales/sales-reps/filter?region=&deal_status=` - Filter sales representatives by region and/or deal status

### AI API

- `POST /api/ai` - Send a question to the AI and get a response

## Project Structure

```
.
├── cmd/                # Command-line commands
│   └── server.go       # Server command implementation
├── config/             # Configuration handling
├── docs/               # Swagger documentation
├── internal/           # Internal packages
│   ├── api/            # API handlers
│   │   ├── ai/         # AI-related endpoints
│   │   └── sales/      # Sales-related endpoints
│   ├── config/         # Configuration structures
│   └── domain/         # Domain models
├── pkg/                # Reusable packages
│   └── utils/          # Utility functions
├── go.mod              # Go module file
├── go.sum              # Go module checksums
└── main.go             # Application entry point
```

## Development

### Adding New Endpoints

1. Create a new handler in the appropriate package under `internal/api/`
2. Register the routes in the handler's `RegisterRoutes` method
3. Add the handler to the server setup in `cmd/server.go`

### Generating Swagger Documentation

This project uses [swaggo/swag](https://github.com/swaggo/swag) for generating Swagger documentation.

To regenerate the documentation after making changes:

```bash
swag init
```

## Testing

Run the tests:

```bash
go test ./...
```

## License

[Specify your license here]

## Contact

[Your contact information]

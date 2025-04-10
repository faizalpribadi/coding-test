# InterOpera Coding Test API (Python)

A RESTful API server built with FastAPI, providing sales data endpoints and AI integration.

## Overview

This project is a Python-based API server that provides endpoints for sales data and AI-powered responses. It uses the FastAPI framework for handling HTTP requests and Google's Gemini API for AI capabilities.

## Features

- **Sales Data API**: Endpoints to retrieve and filter sales representatives and their deals
- **AI Integration**: Integration with Google's Gemini API for AI-powered responses
- **FastAPI Framework**: Modern, fast web framework for building APIs
- **Automatic Documentation**: Interactive API documentation with Swagger UI
- **CORS Support**: Cross-Origin Resource Sharing configuration

## Prerequisites

- Python 3.8 or higher
- API key for Google Gemini (for AI features)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/interopera-coding-test.git
   cd interopera-coding-test/backend/python
   ```

2. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```

3. Configure the Gemini API key:
   
   Edit the `config.py` file and replace the placeholder API key with your actual Gemini API key:
   ```python
   GEMINI_API_KEY = "your-gemini-api-key"
   ```

## Running the Application

Start the server:

```bash
python main.py
```

The server will start on `http://0.0.0.0:8000` by default.

## API Documentation

Once the server is running, you can access the automatic API documentation at:

```
http://localhost:8000/docs
```

## API Endpoints

### Sales API

- `GET /data` - Get all sales data
- `GET /sales-reps` - Get all sales representatives
- `GET /sales-reps/filter?region=&deal_status=` - Filter sales representatives by region and/or deal status

### AI API

- `POST /ai` - Send a question to the AI and get a response
  - Request body: `{"question": "Your question here"}`
  - Response: `{"answer": "AI-generated answer"}`

## Project Structure

```
.
├── api/                # API endpoints
│   ├── ai/             # AI-related endpoints
│   │   └── routes.py   # AI route handlers
│   └── sales/          # Sales-related endpoints
│       └── routes.py   # Sales route handlers
├── helpers/            # Helper utilities
│   └── json_utils.py   # JSON file handling utilities
├── config.py           # Configuration settings
├── main.py             # Application entry point
└── requirements.txt    # Python dependencies
```

## Development

### Adding New Endpoints

1. Create a new route file in the appropriate package under `api/`
2. Define your route handlers using FastAPI decorators
3. Import and include your router in `main.py`

### Error Handling

The application includes a global exception handler for HTTP exceptions. You can raise `HTTPException` with appropriate status codes and error messages in your route handlers.

## Testing

You can test the API endpoints using the interactive documentation or tools like curl or Postman.

Example curl command to test the AI endpoint:

```bash
curl -X POST "http://localhost:8000/ai" \
     -H "Content-Type: application/json" \
     -d '{"question": "What is FastAPI?"}'
```

## License

[Specify your license here]

## Contact

[Your contact information]

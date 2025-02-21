# API Gateway

## Overview

This project implements an API Gateway in Go using the Gin framework. The API Gateway serves as the single entry point for routing requests from the Next.js frontend to various backend microservices. It handles request routing, logging, error recovery, and can enforce security policies such as authentication and rate limiting.

**Key Features:**
- **Routing:** Routes incoming HTTP requests to the appropriate backend services.
- **Middleware:** Uses logging, recovery, and (optionally) authentication middleware.
- **Environment Configurations:** Supports separate configurations for development and production environments.
- **HTTP Client:** Contains HTTP clients for forwarding requests to backend microservices.
- **Microservices Integration:** Designed to work within a microservices architecture (e.g., Trip Planning, User Management, AI Trip Planning, etc.).
- **Structured Logging:** Uses Logrus for JSON-formatted, structured logging.

## Project Structure
## Project Structure

```
./
├── README.md
├── cmd
│   └── server
│       └── main.go
├── config
│   └── config.go
├── go.mod
├── go.sum
├── handlers
│   └── trip_handler.go
├── middleware
│   └── logging.go
├── pkg
│   └── client
│       └── trip_client.go
└── routes
    └── routes.go
```

## Getting Started

### Prerequisites

- **Go:** Version 1.18 or later.
- **Git:** For source control.
- Optionally, **Docker** if you plan on containerizing the application.

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourorg/api-gateway.git
   cd api-gateway
Install dependencies:

Go modules will handle your dependencies. Run:

bash
Copy
go mod tidy
Set Up Configuration:

Create a .env file in the project root for development:

dotenv
Copy
PORT=8080
ENV=development
TRIP_SERVICE_ENDPOINT=http://localhost:9000/trip
READ_TIMEOUT=5
WRITE_TIMEOUT=10
Running the API Gateway
Development Mode
To run the server in development mode, use:

bash
Copy
go run cmd/server/main.go
You should see log messages indicating that the server is running (e.g., on port 8080). Access the API Gateway at http://localhost:8080.

Production Mode
Build the Binary:
```bash
go build -o api-gateway cmd/server/main.go
```

Configure Environment Variables for Production:

Ensure you set the appropriate environment variables (or use a production .env file). For example:

```bash
export ENV=production
export PORT=8080
export TRIP_SERVICE_ENDPOINT=https://your-trip-service.example.com/trip
```
Run the Binary:
```bash
./api-gateway
```
For production deployments, consider using a process manager (like systemd or Docker) and a load balancer to manage the service.

Testing the API Gateway
You can test the API endpoints using curl, Postman, or any REST client. For example, to test the trip creation endpoint:
# Project: CEP Lookup Service

This project is a simple HTTP server written in Go that allows users to fetch postal address information by CEP (Brazilian postal code) using the [ViaCEP API](https://viacep.com.br/). The server implements caching to reduce the number of API calls and improve response time for frequently accessed CEPs.

## Features
- Fetch address information using a CEP.
- Cache results to minimize external API calls.
- Health check endpoint to verify server availability.
- Modular and extensible codebase.

## Project Structure
```
/cmd
    /server
        main.go
/internal
    /handlers
        cep_handler.go
    /services
        cep_service.go
    /cache
        cache.go
/pkg
    /models
        cep.go
```
### Components
- **`/cmd/server/main.go`**: Initializes the server and sets up the routes.
- **`/internal/handlers`**: Contains HTTP route handlers.
- **`/internal/services`**: Handles the core logic, including calling the ViaCEP API and caching results.
- **`/internal/cache`**: Manages caching functionality.
- **`/pkg/models`**: Defines the data models (e.g., `Cep`).

## Endpoints

### Health Check
- **GET /**
  - Response: `ping`
  - Purpose: Check server availability.

### Fetch CEP Information
- **GET /cep/{id}**
  - Path Parameter:
    - `id`: CEP (e.g., `01001000`)
  - Response:
    ```json
    {
      "cep": "01001-000",
      "logradouro": "Praça da Sé",
      "complemento": "lado ímpar",
      "bairro": "Sé",
      "localidade": "São Paulo",
      "uf": "SP",
      "unidade": "",
      "ibge": "3550308",
      "gia": "1004"
    }
    ```
  - Purpose: Retrieve postal address information by CEP.

## How to Run

### Prerequisites
- Go 1.19+

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/eduardonakaidev/via-cep-golang.git
   cd cep-lookup-service
   ```
2. Build the project:
   ```bash
   go build -o server ./cmd/server
   ```
3. Run the server:
   ```bash
   ./server
   ```
4. Access the server:
   - Health Check: `http://localhost:3000/`
   - Fetch CEP: `http://localhost:3000/cep/{id}`

## Environment Variables
- `CACHE_TIME`: Duration (in seconds) for which the CEP data is cached. Default: `500`.
- `SERVER_PORT`: Port on which the server runs. Default: `3000`.

## Example Usage
1. Start the server.
2. Make a GET request to the endpoint:
   ```bash
   curl http://localhost:3000/cep/01001000
   ```
3. Receive the address information in JSON format.

## Improvements
- Add authentication for restricted access.
- Support additional APIs for international postal codes.
- Add more advanced caching mechanisms (e.g., Redis).
- Implement rate limiting to prevent abuse.

## License
This project is open-source and available under the MIT License.

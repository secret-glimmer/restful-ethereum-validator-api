# RESTful Ethereum Validator API



## Objective:

Develop a small RESTful API application using Go. This application will interact with Ethereum data to provide information about block rewards and validator sync committee duties per slot.

## Task Requirements:
- Language and Frameworks
    - The application must be written in Go (Golang).
    - Use any Go framework/libraries you deem necessary for REST API development and interaction with Ethereum.

## Endpoints:
1. ### GET /blockreward/{slot}
- **Purpose**: 
    - Retrieves information about the block reward for a given slot.
- **Parameters**:
    - **slot** (integer): The slot number in the Ethereum blockchain.
- **Response** (JSON):
    - **status**: Indicates if the slot contains a block produced by a MEV relay or a vanilla block.
    - **reward**: The reward amount for the validator (in GWEI).
- **Error Handling**:
    - **404**: Slot does not exist / was missed.
    - **400**: Requested slot is in the future.
    - **500**: Internal server error.
2. ### GET /syncduties/{slot}
- **Purpose**:
    - Retrieves a list of validators with sync committee duties for a given slot.
- **Parameters**:
    - slot (integer): The slot number in the Ethereum blockchain.
- **Response**:
    - A list of public keys of validators with sync committee duties for the specified slot.
- **Error Handling**:
    - **404**: Slot does not exist.
    - **400**: Requested slot is too far in the future.
    - **500**: Internal server error.
    
## Data Sources:
- Use the following Quicknode endpoint for RPC:
    - **HTTP**: <https://radial-misty-butterfly.quiknode.pro/d71f751e03f2b6466202f2561941b6c1c0defd13>
    - **WebSocket**: <wss://radial-misty-butterfly.quiknode.pro/d71f751e03f2b6466202f2561941b6c1c0defd13>

## Provide a README file with:
- Instructions on how to build and run your application.
- A brief explanation of your design choices and frameworks used.
- Examples of how to call the API endpoints using curl or any other HTTP client.

## Run Application on Local:
- **Copy `.env.example` and rename it to `.env`.**

- **Install packages.**
    ``` bash
    go mod download
    ```

- **Create swagger documentation.**
    ``` bash
    swag init -g cmd/main.go
    ```
- **Run application.**
    ``` bash
    go run ./cmd
    ```
- **You can see swagger documentation at this url.**
    
    <http://localhost:8080/docs/index.html>.
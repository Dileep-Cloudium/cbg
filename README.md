# Serverless Framework Node HTTP API on AWS

This repository is for the Nexus Gate PBM API. This solution is intended to act as a collection of CBG services that support the CBG PBM APIs.
The CBG PBM API contains endpoint collections (proxies) for member portal v2, Cervey API wrapper, and GuiltySpark Hasura API wrapper, all maintained
in one repo based on the Go Gin Proxy approach similar to how we use in the NeonArc Utility API. All our teams collaborate on the same API when we
need new endpoints. The goal with this approach is to minimize the number of repos and APIs we have to maintain for our lean team.

Within the CBG PBM API, we hit different "wrappers" or APIs with URL path segments:

- `/member-api/such-and-such` (member portal)
- `/claims-api/such-and-such` (Cervey API)
- `/guiltyspark-api/such-and-such` (GuiltySpark Hasura API)

The APIs are both RESTful and RPC in nature. This solution deploys a HTTP API to AWS. The API Gateway v2 is backed by the lambda function
with a Go runtime running the Gin engine. Persistence is provided by interactions with Amazon DynamoDB, Hasura, and other services, as needed.

## Prerequisites

- Go installed
- AWS credentials configured
- Make utility
- AWS CLI
- Serverless Framework v4 installed

## Project structure

Here's a recommended project structure that follows Go best practices, and is used in this project.

```plaintext
project-root/
├── cmd/                     # Application entry points
│   ├── ginapi/              # Entrypoint for lambda function
│   │   └── main.go
├── internal/                # Private application packages
│   ├── db/                  # Database operations
│   │   └── dynamodb.go
│   ├── handler/             # HTTP handlers
│   │   ├── portal_api/           # Member portal handlers
│   │   │   ├── member_handler.go
│   │   │   └── auth_handler.go
│   │   ├── cervey_api/          # Cervey API handlers
│   │   │   ├── claims_handler.go
│   │   │   └── provider_handler.go
│   │   ├── guiltyspark_api/     # GuiltySpark handlers
│   │   │   └── graphql_handler.go
│   │   └── common/              # Shared handlers
│   │       └── health_handler.go
│   ├── middleware/          # Custom middleware
│   │   └── auth.go
│   ├── router/              # API Routing and Management
│   │   ├── portal_api/          # Member portal routes
│   │   │   └── routes.go
│   │   ├── cervey_api/         # Cervey API routes
│   │   │   └── routes.go
│   │   ├── guiltyspark_api/    # GuiltySpark routes
│   │   │   └── routes.go
│   │   └── router.go           # Main router that combines all sub-routers
│   ├── models/              # Data models
│   │   └── user.go
│   ├── repository/          # Database operations
│   │   └── user_repo.go
│   ├── service/             # Business logic
│   │   ├── portal_api/         # Member portal services
│   │   │   ├── member_service.go
│   │   │   └── auth_service.go
│   │   ├── cervey_api/        # Cervey API services
│   │   │   ├── claims_service.go
│   │   │   └── provider_service.go
│   │   ├── guiltyspark_api/   # GuiltySpark services
│   │   │   └── graphql_service.go
│   │   └── common/            # Shared services
│   │       └── health_service.go
│   └── types/               # Type definitions
│       ├── request          # Request types
│       │   └── user_request.go
│       └── response         # Response types
│           └── user_response.go
├── pkg/                     # Public packages that can be used by other projects
│   ├── utils/
│   └── config/
├── api/                     # API documentation and schemas
│   └── swagger/
├── config/                  # Configuration files
│   └── config.yaml
├── scripts/                 # Build and deployment scripts
│   ├── deploy-serverless.sh
│   └── remove-serverless.sh
├── bin/                     # Build artifacts
├── .env                     # Environment variables
├── .gitignore
├── go.mod
├── go.sum
├── Makefile                 # Makefile for building and deploying the application
└── README.md
```

The key directories and their purposes:

1. `cmd/ginapi/main.go`: Contains the application entry point. This is where you initialize your Gin router, setup middleware, and start the server.

2. `internal/`: Contains packages that are private to your application:
    - `handler/`: HTTP handlers that process requests and responses
    - `middleware/`: Custom middleware functions
    - `router/`: API Routing and Management
    - `models/`: Data structures and domain models
    - `repository/`: Database operations and data access layer
    - `service/`: Business logic implementation
    - `types/`: Defines the structure for requests and responses

3. `pkg/`: Contains packages that could be used by external projects. Keep utility functions and shared code here.

4. `api/`: API documentation and OpenAPI/Swagger specifications.

5. `config/`: Configuration files and settings.

6. `migrations/`: Database migrations

7. `scripts/`: Build and deployment scripts

## Building and developing

Install and upgrade dependencies.

```zsh
make upgrade
```

Build all Lambda functions.

```zsh
make build
```

Run all unit tests (requires AWS credentials).

```zsh
AWS_PROFILE="arcadia-develop" make test
```

Update the SSM parameters for the serverless application.

```zsh
CBG_ACCOUNT_SET="arcadia"
CBG_ENVIRONMENT="develop"

aws sso login --profile "$CBG_ACCOUNT_SET-$CBG_ENVIRONMENT"

export AWS_DEFAULT_REGION="us-west-2"
export AWS_PROFILE="$CBG_ACCOUNT_SET-$CBG_ENVIRONMENT"

(
  GRAPHQL_API_KEY="your-api-key-here" \
  ./scripts/update-params.sh -p "$CBG_ACCOUNT_SET-$CBG_ENVIRONMENT" -r "us-west-2" -s "$CBG_ENVIRONMENT"
)
```

Build and deploy the serverless application to AWS (requires AWS credentials).

```zsh
make build

CBG_ACCOUNT_SET="arcadia"
CBG_ENVIRONMENT="develop"

export DATADOG_API_KEY="954e911dcb46190af7f124780f2517bc"
./scripts/deploy-serverless.sh -p "$CBG_ACCOUNT_SET" -e "$CBG_ENVIRONMENT" \
    -c "CBG" -r "us-west-2"
```

Remove the serverless application from AWS (requires AWS credentials).

```zsh
CBG_ACCOUNT_SET="arcadia"
CBG_ENVIRONMENT="develop"

./scripts/remove-serverless.sh -p "$CBG_ACCOUNT_SET" -e "$CBG_ENVIRONMENT" \
    -c "CBG" -r "us-west-2"
```

## Example API calls

Create a new user.

```zsh
curl -X POST "https://hdf5sio670.execute-api.us-west-2.amazonaws.com/api/v1/user" -H 'accept: application/json' -H 'Authorization: {token}' -H "Content-Type: application/json" -d '{"email": "test@gmail.com","first_name": "test", "last_name": "last", "password": "Pass0wrd"}'

# OR with the friendly domain name

curl -X POST "https://nexusgate-pbmapi.cbgapps-develop.io/api/v1/user/list" -H 'accept: application/json' -H 'Authorization: {token}' -H "Content-Type: application/json" -d '{"email": "test@gmail.com","first_name": "test", "last_name": "last", "password": "Pass0wrd"}'
```

Read the user resource.

```zsh
curl -X GET "https://hdf5sio670.execute-api.us-west-2.amazonaws.com/api/v1/user/{user_id}" -H 'accept: application/json' -H 'Authorization: {token}'

# OR with the friendly domain name

curl -X GET "https://nexusgate-pbmapi.cbgapps-develop.io/api/v1/user/{user_id}" -H 'accept: application/json' -H 'Authorization: {token}'
```

Update the user.

```zsh
curl -X PUT "https://hdf5sio670.execute-api.us-west-2.amazonaws.com/api/v1/user/{user_id}" -H 'accept: application/json' -H 'Authorization: {token}' -H "Content-Type: application/json" -d '{"first_name": "test1", "last_name": "last1"}'

# OR with the friendly domain name

curl -X PUT "https://nexusgate-pbmapi.cbgapps-develop.io/api/v1/user/{user_id}" -H 'accept: application/json' -H 'Authorization: {token}' -H "Content-Type: application/json" -d '{"first_name": "test1", "last_name": "last1"}'
```

Delete the user.

```zsh
curl -X DELETE "https://hdf5sio670.execute-api.us-west-2.amazonaws.com/api/v1/user/{user_id}" -H 'accept: application/json' -H 'Authorization: {token}'

# OR with the friendly domain name

curl -X DELETE "https://nexusgate-pbmapi.cbgapps-develop.io/api/v1/user/{user_id}" -H 'accept: application/json' -H 'Authorization: {token}'
```

List all users.

```zsh
curl -X GET "https://hdf5sio670.execute-api.us-west-2.amazonaws.com/api/v1/user/list" -H 'accept: application/json' -H 'Authorization: {token}'

# OR with the friendly domain name

curl -X GET "https://nexusgate-pbmapi.cbgapps-develop.io/api/v1/user/list" -H 'accept: application/json' -H 'Authorization: {token}'
```

List all states.

```zsh
curl -X GET "https://hdf5sio670.execute-api.us-west-2.amazonaws.com/api/v1/public/states"

# OR with the friendly domain name

curl -X GET "https://nexusgate-pbmapi.cbgapps-develop.io/api/v1/public/states"
```

## Documentation

The API documentation will soon be available in the `api/swagger` directory. The `index.html` file will be the main entry point for the Swagger UI.

For now, there is documentation for specific services listed below.

- Redirect service: [Redirect Service Documentation](./docs/redirect-service.md)

## Function Naming Conventions

Our APIs combine both CRUD(L) operations (REST-style) and RPC capabilities. We follow these naming conventions to maintain consistency and clarity across the codebase.

### CRUD(L) Operations

For standard database/resource operations, use the following pattern:

```go
{Action}{Resource}

// Examples:
CreateUser(...)
ReadUser(...)
UpdateUser(...)
DeleteUser(...)
ListUsers(...) // Note: plural for List operations
```

### RPC-Style Operations

For RPC-style operations, use clear, action-based names that describe the operation:

```go
// Good examples:
ActivateAccount(...)
ValidateCredentials(...)
ProcessPayment(...)
GenerateReport(...)

// Avoid ambiguous names like:
Handle(...)
Process(...)
Do(...)
```

### Why These Conventions?

1. **Clarity**: Including the resource name (e.g., `CreateUser` vs just `Create`) makes the code self-documenting and prevents ambiguity when working with multiple services
2. **Consistency**: Standard CRUD(L) prefixes (`Create`, `Read`, `Update`, `Delete`, `List`) align with REST conventions while remaining explicit
3. **Discoverability**: Resource-suffixed names make IDE code completion more useful
4. **Mixed-Pattern Support**: These conventions work well for both REST and RPC patterns, which our API needs to support
5. **Code Review**: Clear naming patterns make code reviews easier and help maintain consistency across the team

### Anti-patterns to Avoid

```go
// Avoid generic names
Process() // Too vague
HandleUserStuff() // Too informal
DoUserCreation() // Redundant "Do" prefix

// Instead, be specific and direct
ProcessPayment()
ValidateUserCredentials()
CreateUser()
```

### Tips

- Use active verbs for RPC operations
- Keep names concise but descriptive
- If you're creating a new pattern, discuss with the team first
- When in doubt, favor clarity over brevity

## Contribution

The quality of the code is very important. Please make sure to follow the project structure and best practices. To contribute to this project, please follow these steps:

1. Clone the repository
2. Create a new feature branch tied to the story number
3. Make your changes
4. Add tests for your changes (this is a MUST)
5. Create a pull request

Thank you for your contribution!

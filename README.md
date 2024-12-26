<h1>Golibaba Backend Microservice</h1>
The Golibaba is a modular and scalable application inspired by travel platforms like Alibaba. It provides a comprehensive system for managing travel plans, bookings, and payments for different modes of transportation (e.g., buses, trains, planes, ships) and accommodations (e.g., hotels). The project aims to:

- Simplify travel planning and booking.
- Offer dynamic pricing and operator management.
- Ensure secure and efficient handling of user data and payments.

## Menu
<!-- TOC -->
  * [Menu](#menu)
  * [How to run the project](#how-to-run-the-project)
    * [Up & Down](#up--down-)
    * [Build](#build)
    * [Logs](#logs)
  * [Microservice Structure](#microservice-structure)
  * [Service Structure Sample](#service-structure-sample)
  * [Features](#features)
    * [1. User Management](#1-user-management)
    * [2. Transportation Services](#2-transportation-services)
    * [3. Hotel Management](#3-hotel-management)
    * [4. Tour Packages](#4-tour-packages)
    * [5. Notifications](#5-notifications)
    * [6. Payments and Wallet](#6-payments-and-wallet)
  * [Architecture](#architecture)
  * [Tools and Technologies](#tools-and-technologies)
    * [Programming Language](#programming-language)
    * [Database](#database)
    * [Communication](#communication)
    * [Containerization](#containerization)
    * [Monitoring and Logging](#monitoring-and-logging)
    * [Additional Tools](#additional-tools)
  * [Security](#security)
  * [Project Logic](#project-logic)
    * [User Journey](#user-journey)
    * [Admin Features](#admin-features)
    * [Operators](#operators)
    * [Business Rules](#business-rules)
  * [Deployment](#deployment)
  * [Development Guidelines](#development-guidelines)
  * [Future Enhancements](#future-enhancements)
  * [Contribution](#contribution)
<!-- TOC -->

## How to run the project
You can use Makefile to up, down, build project.

### Up & Down 
```bash
make up
```
This command runs `up` commands in `api_gateway` service and all the services in `services` directory. Every service should also have a Makefile in their root directory that handles up command.

Also, you can use this command:
```bash
make down
```
This command runs `down` command in `api_gateway` and all the services.
Some services maybe remains `still_in_use` state. You can run down command twice to make sure they are down.
### Build
When you run the project for the first time, docker build all the containers automatically. But after changing the code base, docker only runs the cached binary. To avoid this, you can run a command like this to build changed service before start that.
```bash
make -C services/user build
make up
```

### Logs
You can see every container log just by using docker commands:
```bash
# You can use -f flag to continue watching logs
docker logs golibaba-gateway-app
```

## Microservice Structure
```bash
.
│─── api_gateway
│─── compose
│    │─── nginx
│    │    │─── nginx.conf
│    │    │─── docker-compsoe.yaml
│    │    └─── ssl
│    └─── rabbitmq
│         └─── docker-compsoe.yaml
│─── common
│    │─── service_domain.go
│    │─── other_service_domain.go
│    │─── other_service_port.go
│    │─── ...
│    └─── go.mod
│─── modules
│    │─── cache
│    │─── gateway_client
│    │    │─── client.go
│    │    │─── contract.go
│    │    └─── go.mod
│    │─── user_service_client
│    │─── other_services_client
│    └─── ...
│─── services
│    │─── user
│    │─── hotels
│    │─── notifications
│    │─── payment
│    │─── travel
│    │─── hotel
│    └─── other services
│─── proto
│    │─── pb
│    │─── proto
│    └─── Makefile
└─── Makefile
```

## Service Structure Sample
```bash
.
│─── cmd
│    └─── main.go
│─── app
│    │─── contact.go  
│    └─── app.go
│─── api
│    │─── http
│    │    │─── handlers
│    │    │    │─── helpers
│    │    │    │─── presenters
│    │    │    └─── sample_api_handler.go...
│    │    │─── middlewares
│    │    │─── services
│    │    │─── types
│    │    └─── server.go
│    └─── service
│         └─── route services..
│─── config
│    │─── read.go   
│    └─── type.go  
│─── build   
│    │─── redis
│    │    ├─── docker-compose.yaml
│    │    └─── .env
│    └─── project
│         ├─── Dockerfile
│         ├─── docker-compose.yaml
│         └─── .env
│─── internal
│    │─── common # sub domains
│    │    ├─── identifiers.go
│    │    └─── interfaces.go
│    └─── sample_logic
│    │    ├─── port
│    │    │    ├─── sample_logic.go
│    │    │    └─── service.go
│    │    ├─── domain
│    │    │    └─── sample_logic.go
│    │    └─── service.go
│─── pkg
│    ├─── adapters
│    │    ├─── cache
│    │    │    └─── redis.go
│    │    ├─── email
│    │    │    └─── email.go
│    │    ├─── rbac
│    │    │    └─── casbin.go
│    │    └─── storage
│    │         │─── helpers
│    │         │─── mapper
│    │         │─── migrations
│    │         │─── types
│    │         └─── repository files...
│    ├─── postgres
│    │    └─── gorm.go
│    ├─── cache
│    │    ├─── provider.go
│    │    └─── serialization.go
│    ├─── logger
│    ├─── hash
│    ├─── context
│    └─── jwt
│         ├─── claims.go
│         └─── auth.go
│─── tests
│─── config.json
│─── sample-config.json
│─── go.mod
│─── go.sum
│─── README.md
│─── LICENSE
│─── .gitignore
└─── Makefile
```

## Features

### 1. User Management
- **Registration and Authentication**: Users can register with an email and password, and their identity is verified during login.
- **Role-Based Access Control (RBAC)**: Roles include Admin, Travel Agencies, Operators, and Regular Users. Users can have multiple roles, and permissions vary by role.

### 2. Transportation Services
- Support for multiple transportation types (e.g., buses, trains, planes, ships).
- Companies can create dynamic travel options with customizable schedules and prices.
- Centralized management of routes, operators, and tickets.

### 3. Hotel Management
- Hotels can list rooms with different features and dynamic pricing.
- Users can book rooms, and payments are securely processed.
- Admins monitor transactions and manage hotel registrations.

### 4. Tour Packages
- Travel agencies can create comprehensive tour packages, including transportation and accommodations.
- Dynamic pricing based on availability and user demand.

### 5. Notifications
- Users receive notifications for key events, such as booking confirmations, cancellations, and reminders.
- Customizable notification settings based on user preferences.

### 6. Payments and Wallet
- Integrated wallet system for secure transactions.
- Supports deposits, withdrawals, and refunds.
- Payments are linked to invoices generated by the platform.

## Architecture
The project is built using the **Hexagonal Architecture** to ensure modularity, scalability, and ease of testing. Core components include:

- **Domain Layer**: Encapsulates business logic and rules.
- **Ports and Adapters**: Provides interfaces for external systems like databases and APIs.
- **Microservices**: Decoupled services for User Management, Transportation, Hotel Management, Payment, and Notification.

## Tools and Technologies

### Programming Language
- **Go (Golang)**: The primary language for building all microservices.

### Database
- **PostgreSQL**: Used for data persistence and storage.

### Communication
- **gRPC**: For service-to-service communication ensuring high performance.
- **Messaging System**: Includes support for tools like Kafka or NATS for event-driven processing.

### Containerization
- **Docker**: Each microservice runs in its own container, managed by Docker Compose.

### Monitoring and Logging
- **Prometheus** and **Grafana**: For resource monitoring and visualization.
- **Logstash**: Centralized logging to track logs by scope (e.g., authentication, transactions).

### Additional Tools
- **Nginx**: Reverse proxy for service routing.
- **MinIO**: Object storage for media and document files.
- **OpenTelemetry**: For distributed tracing and telemetry.

## Security
- **Authentication**: Secure login with hashed passwords using tools like bcrypt.
- **Authorization**: Role-based access control.
- **Encryption**: Sensitive data is encrypted using symmetric and asymmetric cryptography.
- **Anti-DOS Protection**: Rate limiting and blocking suspicious requests.

## Project Logic

### User Journey
1. **Registration**: Users sign up and are assigned roles.
2. **Service Access**: Depending on roles, users access booking systems, manage tours, or administer the platform.
3. **Booking and Payment**: Users search for services, book, and complete payments.
4. **Notifications**: Users are informed of booking updates and reminders.

### Admin Features
- Manage roles, access, and content across the platform.
- Monitor transactions and resolve disputes.
- Block/unblock entities (users, companies, or services).

### Operators
- Create and manage travel schedules, routes, and pricing.
- Approve or cancel trips based on user demand and availability.

### Business Rules
- Dynamic pricing adjusts based on demand and booking times.
- Revenue distribution among stakeholders is automated upon service completion.

## Deployment
- All microservices are deployed using Docker.
- A shared network connects all services.
- Reverse proxying is configured using Nginx.

## Development Guidelines

1. **Modularity**: Ensure components are loosely coupled and highly cohesive.
2. **CRUD Operations**: Fully implemented for all entities.
3. **Caching**: Used to improve system performance.
4. **Queueing**: Implemented to handle asynchronous tasks.
5. **Git Workflow**: Proper branching and commits for team collaboration.

## Future Enhancements
- Kubernetes for container orchestration.
- Advanced analytics for user behavior and system optimization.
- Support for additional payment gateways.

---
## Contribution
Feel free to fork this repository and submit pull requests. Ensure your contributions follow the coding standards and include necessary documentation.


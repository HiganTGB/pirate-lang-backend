#  TOEIC Test and Practice API

## Overview

## Features
- **Modular Architecture**
    - Controller for HTTP request handling
    - Service  for business logic
    - Repository for data persistence

- **Database Integration**
    - PostgreSQL integration with SQLC
    - Database migrations support
    - Repository pattern implementation
    - Transaction management

- **API Features**
    - RESTful endpoints
    - Request validation
    - Error handling
    - Response formatting

- **Security & Performance**
    - Authentication middleware
    - Request rate limiting
    - Graceful shutdown
    - Logging system

- **Development Tools**
    - Environment configuration
    - Docker containerization

## Technical Stack

### Core
- Go 1.21+
- Echo Framework - HTTP router and middleware
- Viper - Configuration management

### Database & Storage
- PostgreSQL 15+ - Primary database
- SQLC - SQL toolkit
- golang-migrate - Database migrations
- Minio - S3 storage
- Redis

### Development & Tools
- Docker & Docker Compose
- Make - Build automation


### Monitoring & Logging
- Log - Structured logging

# Temporal Demo with PostgreSQL and ScyllaDB

This project showcases a temporal workflow using PostgreSQL and ScyllaDB as backend databases. It includes Docker Compose for setting up the infrastructure, a worker component to handle tasks, and a starter component to trigger the cron workflow.

## Getting Started

### Prerequisites
- Go programming language
- Docker
- Docker Compose
- PostgreSQL
- ScyllaDB

### Installation

1. Clone the repository:
   ```bash
2. docker-compose up -d

3. go run ./worker 
4. go run ./starter

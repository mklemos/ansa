
# Ansa Project

This project is a payment processing system using Go, Echo framework, and PostgreSQL.

## Setup Instructions

### Prerequisites

- Go (version 1.16+)
- PostgreSQL (version 12+)
- Migration tool: [migrate](https://github.com/golang-migrate/migrate)

### Database Setup

1. **Start PostgreSQL Server**

   ```bash
   pg_ctl -D /usr/local/var/postgres start
   ```

2. **Initialize Database**

   ```bash
   createdb mydatabase
   ```

3. **Create User**

   ```bash
   createuser myuser -P
   ```

4. **Run Migrations**

   ```bash
   migrate -path db/migrations -database postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable up
   ```

### Running the Application

1. **Start the Server**

   ```bash
   go run main.go
   ```

2. **API Endpoints**

   - **Create Customer**
     ```bash
     curl -X POST http://localhost:8080/customers -H "Content-Type: application/json" -d '{"CustomerID": "2", "Name": "Jack Doe", "Email": "jack@example.com"}'
     ```

### Stopping the Database Server

```bash
pg_ctl -D /usr/local/var/postgres stop
```

### Starting the Database Server

```bash
pg_ctl -D /usr/local/var/postgres start
```

## Troubleshooting

- Ensure the PostgreSQL server is running before starting the application.
- Check for any connection errors and validate database credentials.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

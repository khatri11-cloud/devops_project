## Step 1 - Setting up the Go Application

Today I initialized the Go module and created the initial project structure. Instead of placing everything in a single file, I chose a structure that separates the application entry point (`cmd`) from the internal business logic (`internal`). Although this project is small, this layout is commonly used in production Go applications because it scales well as the project grows.

### Learned

- How Go modules work
- Why production Go projects use `cmd/` and `internal/`
- How to create a basic HTTP server using Go's standard library
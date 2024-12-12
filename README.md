# gRPC Blog Management Project

This project is a gRPC-based blog management application implemented using **Go**, **MongoDB**, and **Docker**. The application provides the following core functionalities:

- **Create Blog**
- **Read Blog**
- **Update Blog**
- **Delete Blog**
- **List Blogs**

The application is containerized using Docker for easy deployment and scalability.

## Features

- **gRPC for Communication**: Provides a fast and efficient RPC framework for communication between the client and the server.
- **MongoDB for Data Storage**: Stores blog data persistently.
- **Dockerized Setup**: Simplifies deployment with Docker and Docker Compose.
- **CRUD Operations**: Full support for Create, Read, Update, and Delete operations on blog posts.

## Technologies Used

- **Go**: Programming language used for implementing the application.
- **gRPC**: RPC framework for client-server communication.
- **MongoDB**: Database used for storing blog information.
- **Docker**: Used for containerizing the application.
- **Mongo Express**: Web-based MongoDB admin interface.

## Installation

1. **Clone the Repository**:
 ```bash
 git clone <repository-url>
 cd grpc-project
 ```

Ensure you have Docker and Docker Compose installed on your system.

Build and start the services using Docker Compose:
```bash
docker-compose up --build
```

The MongoDB service will be available on port 27017, and Mongo Express will be accessible at http://localhost:8081.

## Usage

### Running the Application

The application is structured with separate client and server implementations.
The server handles gRPC requests and interacts with the MongoDB database.
The client sends gRPC requests to perform CRUD operations.

Run the server:
```bash
go run server/main.go
```

Run the client:
```bash
go run client/main.go
```

### Available gRPC Functions

- **CreateBlog**: Adds a new blog entry.
- **ReadBlog**: Fetches the details of a specific blog by ID.
- **UpdateBlog**: Updates an existing blog entry by ID.
- **DeleteBlog**: Deletes a blog entry by ID.
- **ListBlogs**: Lists all available blogs.

## Environment Variables

The `docker-compose.yml` file defines the necessary environment variables:

### MongoDB
- `MONGO_INITDB_ROOT_USERNAME`: Root username for MongoDB.
- `MONGO_INITDB_ROOT_PASSWORD`: Root password for MongoDB.

### Mongo Express
- `ME_CONFIG_MONGODB_ADMINUSERNAME`: Admin username for Mongo Express.
- `ME_CONFIG_MONGODB_ADMINPASSWORD`: Admin password for Mongo Express.
- `ME_CONFIG_MONGODB_URL`: MongoDB connection URL.

## Project Structure

```
grpc-project/
├── bin/                # Compiled binaries
├── client/             # Client implementation
│   ├── main.go         # Client entry point
├── server/             # Server implementation
│   ├── main.go         # Server entry point
│   ├── create.go       # CreateBlog logic
│   ├── read.go         # ReadBlog logic
│   ├── update.go       # UpdateBlog logic
│   ├── delete.go       # DeleteBlog logic
│   ├── list.go         # ListBlogs logic
├── proto/              # Protocol Buffers definitions
│   ├── blog.proto      # Proto file
│   ├── blog.pb.go      # Generated Go code from proto
├── docker-compose.yml  # Docker Compose configuration
├── go.mod              # Go module dependencies
├── go.sum              # Dependency checksum file
├── README.md           # Project documentation
```

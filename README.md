# Micro-Panel: An Ecommerce Microservice Platform for Comic Book Stores  

**Micro-Panel** is a Golang-based microservice project designed to provide a seamless ecommerce experience for comic book stores. This project consists of multiple microservices working together to manage CRUD operations, notifications, and real-time updates.

## Features  

- **RESTful API** for CRUD operations with robust authentication and middleware routing.  
- **gRPC service** for efficient and real-time updates using both unary and streaming calls.  
- **Notification service** with a stateful database-backed queue to notify users of updates.  
- **MySQL** as the database for persistence.  
- **Docker/Docker Compose** for containerized deployments.  

---

## Microservices Overview  

### 1. `micropanel-api`  
This microservice provides the core HTTP RESTful API functionality to manage resources in the system.  

**Key Features:**  
- **CRUD operations**: Handles database queries for comic books, users, orders, and more.  
- **Routing with middleware**: Includes middlewares for logging, request validation, and error handling.  
- **Authentication & Authorization**: Secured with JWT tokens to manage user sessions and permissions.  

---

### 2. `micropanel-grpc`  
This microservice is responsible for handling updates using gRPC for better performance and real-time communication.  

**Key Features:**  
- **Protocol Buffers**: Defines the protobuf schemas for efficient serialization.  
- **Unary gRPC Calls**: Handles single request-response updates.  
- **Streaming gRPC**: Supports server-side streaming for broadcasting updates.  

---

### 3. `notification-svc`  
This service is designed to notify users when updates are added to the queue by `micropanel-grpc`.  

**Key Features:**  
- **Stateful database queue**: Uses a MySQL-backed queue for managing notifications reliably.  
- **User notifications**: Notifies users via email, SMS, or other integrations (configurable).  

---

## Project Architecture  

The project is organized as follows:  

```
micro-panel/  
├── micropanel-api/        # RESTful API service  
├── micropanel-grpc/       # gRPC service  
├── notification-svc/      # Notification service  
├── docker-compose.yml     # Docker Compose setup  
└── README.md              # Project documentation  
```  

---

## Tech Stack  

- **Programming Language**: Golang  
- **Database**: MySQL  
- **Communication**: REST API, gRPC  
- **Containerization**: Docker & Docker Compose  

---

## Getting Started  

### Prerequisites  

- Golang 1.23+  
- Docker & Docker Compose  
- MySQL  

---

### Installation  

1. Clone the repository:  
   ```bash  
   git clone https://github.com/Turtel216/micro-panel.git  
   cd micro-panel  
   ```  

2. Start the services with Docker Compose:  
   ```bash  
   docker-compose up --build  
   ```  

3. The services will be available at:  
   - `micropanel-api`: http://localhost:8080  
   - `micropanel-grpc`: grpc://localhost:50051  

4. Access the MySQL database:  
   ```bash  
   docker exec -it micro-panel_mysql_1 mysql -u root -p  
   ```  

---

## Configuration  

- **Environment Variables**:  
  Each service has its own `.env` file to configure parameters like database connection strings, JWT secrets, etc.  

- **Protobuf Definitions**:  
  Protobuf files are stored under `micropanel-grpc/proto/`. Run the following command to regenerate Go bindings after editing:  
  ```bash  
  protoc --go_out=. --go-grpc_out=. proto/*.proto  
  ```  

---

## Contribution  

1. Fork the repository.  
2. Create a new branch:  
   ```bash  
   git checkout -b feature-name  
   ```  
3. Commit your changes:  
   ```bash  
   git commit -m "Add feature-name"  
   ```  
4. Push to the branch:  
   ```bash  
   git push origin feature-name  
   ```  
5. Submit a pull request.  

---

## License  

This project is licensed under the MIT License. See the `LICENSE` file for details.  

---

## Acknowledgments  

- Inspired by modern microservice design principles.  
- Special thanks to the open-source community for providing tools like Docker, gRPC, and MySQL.  
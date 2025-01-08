# ProductApp - REST API with PostgreSQL and Docker

## 📄 Project Description
ProductApp provides a RESTful API for managing product information. The project is built on a PostgreSQL database and can be easily launched using Docker containers. It supports basic CRUD operations such as adding, updating, deleting, and listing products.

## 🛠 Technologies Used
- *Go (Golang):* Used for API development.
- *PostgreSQL:* Chosen for database management.
- *Docker:* Used to run PostgreSQL in a container.
- *pgx (PostgreSQL driver):* For interacting with PostgreSQL.
- *LabStack Echo:* For HTTP server and routing.

## 🚀 How to Run the Project

### 1. Clone the Repository
```bash
git clone https://github.com/gkhancobanoglu/productapp.git 
cd productapp
```

### 2. Set Up PostgreSQL with Docker
Follow these steps to run the PostgreSQL database and set up the required tables:

#### a. Run the test_db.sh Script
```bash
chmod +x test_db.sh
./test_db.sh
```

This script will create a Docker container and set up the database and table in PostgreSQL.

### 3. Set Up the Environment Configuration

Define the database connection details in the config.go file. Example:
go
postgresql.Config{
    Host: "localhost",
    Port: "6432",
    UserName: "postgres",
    Password: "postgres",
    DbName: "productapp",
    MaxConnections: "10",
    MaxConnectionIdleTime: "30s",
}


### 4. Run the Project
Use the following command to start the API:
```bash
go run main.go
```

### 5. Test the API
Once the project is running, you can test the endpoints using Postman or a similar tool:

#### a. Add Product
- *Endpoint:* POST /products
- *Body:*
json
{
  "name": "Example Product",
  "price": 100.0,
  "discount": 10.0,
  "store": "Example Store"
}


#### b. Get All Products
- *Endpoint:* GET /products

#### c. Get Product by ID
- *Endpoint:* GET /products/:id

#### d. Update Product Price
- *Endpoint:* PUT /products/:id/price
- *Body:*
json
{
  "newPrice": 120.0
}


#### e. Delete Product by ID
- *Endpoint:* DELETE /products/:id

## 📂 Project Structure
```bash
projectapp
├── main.go              # Entry point of the application
├── service              # Business logic layer
├── domain               # Domain models
├── persistence          # Database interaction logic
├── postgresql           # Database connection and configuration
├── handlers             # API request handlers
├── test_db.sh           # Script to set up PostgreSQL with Docker
└── README.md            # Documentation
```

# Desksetup Link API

REST API สำหรับจัดการข้อมูล links และ products (Go implementation)

## Requirements

- Go 1.21 or higher
- PostgreSQL 12+ (Hostinger or local)

## Installation

```bash
# Download dependencies
go mod download

# Or use tidy to clean up
go mod tidy
```

## Configuration

### Environment Variables

Create a `.env` file in the project root:

```env
DATABASE_URL=postgresql://username:password@host:port/dbname?sslmode=disable
PORT=8080
```

**For Hostinger:**
```env
DATABASE_URL=postgresql://user:password@72.62.251.164:5433/desksetup?sslmode=disable
PORT=8080
```

## Database Setup

### Create Tables

Run the seed script to create tables:

```bash
# From JSON samples (recommended)
psql "$DATABASE_URL" -f db/seed_from_json.sql

# Or minimal schema
psql "$DATABASE_URL" -f db/seed.sql
```

### Seed Data

Sample data for products, sub_items (variants), and highlights are included in `db/seed_from_json.sql`.

## Running the Server

### Development
```bash
go run main.go
```

### Production (build binary)
```bash
# Build
go build -o desksetup-api

# Run
./desksetup-api
```

### With hot reload (using air)
```bash
# Install air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

Server will run on `http://localhost:8080` (or configured PORT)

## API Endpoints

### Products

#### 1. Get All Products
```
GET /api/products
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "Reptilian KX78 HE",
      "category": ["keyboard"],
      "brand": "Saru Space",
      "img": "assets/products/saru-kx78he.jpg",
      "tag": "new",
      "description": "Gaming Magnetic Keyboard...",
      "code": "reptilian-kx78-he",
      "links": {
        "shopee": "https://s.shopee.co.th/806LteAKoG",
        "lazada": "",
        "tiktok": "",
        "other": ""
      },
      "created_at": "2026-01-27T10:00:00Z",
      "updated_at": "2026-01-27T10:00:00Z"
    }
  ],
  "count": 1
}
```

#### 2. Get Product by ID
```
GET /api/products/:id
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Reptilian KX78 HE",
    ...
  }
}
```

#### 3. Create Product
```
POST /api/products
```

**Request Body:**
```json
{
  "title": "Reptilian KX78 HE",
  "category": ["keyboard"],
  "brand": "Saru Space",
  "img": "assets/products/saru-kx78he.jpg",
  "tag": "new",
  "description": "Gaming Magnetic Keyboard...",
  "code": "reptilian-kx78-he",
  "links": {
    "shopee": "https://s.shopee.co.th/806LteAKoG",
    "lazada": "",
    "tiktok": "",
    "other": ""
  }
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "message": "Product created successfully",
  "data": { ... }
}
```

#### 4. Update Product
```
PUT /api/products/:id
```

**Response:**
```json
{
  "success": true,
  "message": "Product updated successfully",
  "data": { ... }
}
```

#### 5. Delete Product
```
DELETE /api/products/:id
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Product deleted successfully"
}
```

## Error Responses

### 400 Bad Request
```json
{
  "success": false,
  "error": "Invalid request body"
}
```

### 404 Not Found
```json
{
  "success": false,
  "error": "Product not found"
}
```

### 500 Internal Server Error
```json
{
  "success": false,
  "error": "Internal server error"
}
```

## Database Schema (PostgreSQL)

### products
```sql
CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    brand VARCHAR(100),
    img VARCHAR(255),
    category JSONB,
    description TEXT,
    code VARCHAR(50),
    tag VARCHAR(50),
    links JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

### sub_items (product variants)
```sql
CREATE TABLE sub_items (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    subtitle VARCHAR(100),
    brand VARCHAR(100),
    img VARCHAR(255),
    category JSONB,
    description TEXT,
    code VARCHAR(50),
    shopee_link VARCHAR(255),
    tiktok_link VARCHAR(255),
    lazada_link VARCHAR(255),
    other_link VARCHAR(255),
    display_order SMALLINT DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

### highlights (featured products)
```sql
CREATE TABLE highlights (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    priority SMALLINT DEFAULT 0,
    end_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT unique_product UNIQUE (product_id)
);
```

## Project Structure

```
.
├── main.go                         # Entry point
├── db/
│   ├── db.go                       # Database connection
│   ├── seed.sql                    # Schema only
│   └── seed_from_json.sql          # Schema + sample data
├── domain/
│   ├── product.go                  # Domain models
│   └── errors.go                   # Domain errors
├── repository/
│   └── product_repository.go       # Data access layer
├── service/
│   └── product.go                  # Business logic
├── handler/
│   ├── product_handler.go          # HTTP handlers
│   ├── response.go                 # Response helpers
│   └── router.go                   # Routes
├── go.mod
├── go.sum
├── .env
├── .env.example
└── README.md
```

## Notes

- Nullable fields use pointers (`*string`, `*time.Time`)
- JSONB columns for flexible category and links storage
- Product code as unique identifier alongside ID
- Supports collections via parent products with sub_items
- Timestamps with timezone support

## Future Enhancements

- [ ] Database migrations tooling (goose, golang-migrate)
- [ ] User authentication (JWT)
- [ ] Input validation with validator package
- [ ] Pagination & filtering
- [ ] Rate limiting middleware
- [ ] Caching (Redis)
- [ ] API documentation (Swagger)
- [ ] Logging (zerolog, zap)
- [ ] Configuration management (viper)
- [ ] Graceful shutdown

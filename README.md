# Go Heroes API

A simple REST API that provides CRUD operations on a `hero` object, it was built using Go and PostgreSQL.

## Dependencies

This project makes use of gorilla/mux, gorm, and godotenv.

## Getting Started

**1. Clone the application**

```bash
git clone https://github.com/nicolaspearson/go.heroes.api.git
```

**2. Start the database**

```bash
docker-compose up
```

**3. Build and run the app**

#### Run the app in development mode:

```bash
go run main.go
```

The app will start running at <http://localhost:8000>

#### Run the app in release mode:

```bash
go build
./go-hero
```

The app will start running at <http://localhost:8000>

## Endpoints

The following endpoints are available:

```
POST /user/register
```

```
POST /user/login
```

```
GET /hero?id={heroId}
```

```
GET /heroes
```

```
POST /hero
```

```
PUT /hero/{heroId}
```

```
DELETE /hero/{heroId}
```

## Benchmarking

Run this command to benchmark request performance:

```
wrk -d1m http://localhost:8000/heroes
```

![benchmark](/img/benchmark.png)

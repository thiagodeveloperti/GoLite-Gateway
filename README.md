# GoLite Gateway

**GoLite Gateway** is a lightweight, extensible, and developer-friendly API Gateway written in Go.  
It focuses on simplicity, high performance, and a clean architecture—making it ideal for learning, small projects, and open-source experimentation.

## 🚀 Features

- **Dynamic route management** via YAML configuration  
- **Reverse proxy** powered by Go’s native `httputil.ReverseProxy`  
- Optional **JWT authentication middleware**  
- **Hot-reload configuration watcher** (auto-reload on file changes)  
- Clean and modular project layout (Go best practices)  
- Admin API (future) for managing routes through a UI panel  
- Frontend panel using React or Svelte (future)

## 📁 Project Structure

```text
golite-gateway/
├── cmd/
│ ├── gateway/ # Main API Gateway server
│ └── admin-api/ # Admin API for UI management (future)
├── internal/
│ ├── config/ # YAML config loader + watcher
│ ├── router/ # HTTP routing and dynamic registration
│ ├── proxy/ # Reverse proxy logic
│ ├── auth/ # JWT middleware
│ ├── store/ # In-memory route store
├── web/ # Frontend (React/Svelte)
├── config.yaml # Main configuration file
└── README.md
```

This layout follows common Go conventions and matches patterns used by
projects like Kubernetes, Caddy, Prometheus, and Traefik.

## 🛠 How to Run

1. Install Go 1.25+  
2. Clone the repository  
3. Run the gateway:
```bash
go run ./cmd/gateway
```
4. Gateway will start on:
```bash
http://localhost:8080
```

## ⚙️ Example config.yaml

```yaml
routes:
  - path: /users
    upstream: http://localhost:3001
    methods: [GET, POST]
    auth: false

  - path: /catalog
    upstream: http://localhost:3002
    methods: [GET]
    auth: true

jwt:
  secret: "secret123"
  issuer: "golite"
```


## 📌 Roadmap

**v0.1 (MVP)**
- Dynamic routes
- Reverse proxy
- JWT middleware
- Config watcher

**v0.2**
- Rate limiting
- Import/export routes
- Admin API

**v0.3**
- Plugin system
- Metrics endpoint
- Health checks for upstreams

**v1.0**
- gRPC proxying
- Service discovery (etcd / Consul)
- Clustering support


## 🤝 Contributing

Contributions are welcome!
Feel free to open issues, submit ideas, or send pull requests.


## 📜 License

This project is licensed under the **MIT License** (see below).

MIT is recommended for open-source tools, libraries, and community projects because it:
- allows commercial use
- allows modification
- allows redistribution
- is extremely permissive
- is the most widely used license on GitHub
- encourages contributions without legal friction

You’re free to change it later if needed.

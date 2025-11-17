# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [v0.1.0] - 2025-11-17
### Added
- Dynamic route loading from `config.yaml`.
- Reverse proxy system using `httputil.NewSingleHostReverseProxy`.
- JWT authentication middleware with configurable secret.
- File watcher to auto-reload configuration when `config.yaml` changes.
- Internal health check support using `upstream: self`.
- Structured and extensible project layout (`internal/config`, `internal/router`, `internal/auth`, `internal/proxy`, `internal/store`).
- Initial documentation and release notes.


## [Unreleased]
### Planned
- Rate limiting middleware.
- Admin API for managing routes at runtime.
- Route import/export.
- Plugin system for extending gateway behavior.
- Metrics and observability endpoints.
- Upstream health checks.
- gRPC proxying and service discovery integration (etcd / Consul).
- Clustering support.

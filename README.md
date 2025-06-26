# Distributed Log Server Service

A **distributed, secure, observable commit log system** built in Go. This system offers leader-based replication, gRPC APIs, service discovery, and optional HTTP/JSON interface. Designed for running in local Kubernetes clusters with robust observability, security, and modular extensibility.


## Features

- **gRPC API** for high-performance communication between clients and servers
- **Leader-follower replication** using Raft consensus
- **Serf-based cluster discovery** and membership monitoring
- **Local Kubernetes cluster** support for development and deployment
- **Routing logic** for:
  - Forwarding `Append` requests to the leader
  - Load-balancing `Consume` requests across follower nodes
- **Custom service multiplexing** over a single TCP port
- **Mutual and One-way TLS** for secure communication
- **ACL-based authorization** for fine-grained access control
- **Health checks** via gRPC health service
- **Distributed tracing and metrics** for observability

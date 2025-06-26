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

## Running Tests

```bash
    # initialize the service directory
    make init

    # generate certs
    make gencert

    # run standard test-suite
    make test

    # run metric-specific tests
    make test-metric
```


## Running Cluster Locally

### Prerequisites

Before you begin, make sure you have the following installed:

- `docker`
- `kind` (Kubernetes in Docker)
- `helm`
- `kubectl`

```bash
    # generate docker image
    docker build -t <image-name>:<tag> <path-to-Dockerfile-or-context>

    # spin-up kind cluster
    kind create cluster
    
    # load the created image into the cluster
    kind load docker-image <image-name>
    
    # install helm chart in the cluster
    helm install dlss deploy/dlss

    # check pods status
    kubectl get pods
```
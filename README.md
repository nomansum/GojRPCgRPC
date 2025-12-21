# GojRPCgRPC - Multi-Protocol Order Service in Go (gRPC, JSON-RPC, REST)

This repository demonstrates a sample order management microservice implemented in Go, exposing the **same core business logic** through three different API protocols:

- **gRPC** (high-performance binary protocol)
- **JSON-RPC 2.0** over HTTP
- **RESTful APIs** via gRPC-Gateway

The services are fully containerized with Docker and orchestrated on a local Kubernetes cluster using **Kind**. A complete observability stack is included with **Prometheus** (metrics), **Loki** (logs), **Promtail** (log shipping), and **Grafana** (dashboards and exploration).

Repository: https://github.com/nomansum/GojRPCgRPC/tree/main

## Features

- Shared business logic for order creation and cancellation
- gRPC server implementation
- JSON-RPC 2.0 server over HTTP
- REST endpoints automatically generated from gRPC using gRPC-Gateway
- Separate Docker images for each service (grpc-server, jrpc-server, grpc-gateway)
- Kubernetes manifests for deployments, services, and namespaces
- Full monitoring stack:
  - Prometheus for metrics
  - Loki + Promtail for centralized logging
  - Grafana for visualization of metrics and logs

## Architecture

Client Requests
├── JSON-RPC → jrpc-server (port 30080 via NodePort)
├── REST API  → grpc-gateway (port 30081 via NodePort)
└── gRPC      → grpc-server (internal, accessible via grpc-gateway)
All services → Shared Go business logic (CreateOrder / CancelOrder)



The `grpc-gateway` acts as a reverse proxy that translates incoming REST/JSON-RPC requests into gRPC calls to the backend `grpc-server`.

## Prerequisites

- Git
- Docker
- Go 1.21+ (only needed if modifying code)
- [Kind](https://kind.sigs.k8s.io/) (Kubernetes in Docker)
- kubectl
- curl or any HTTP client for testing

## Installation and Deployment (Local Kind Cluster)

### 1. Clone the Repository

```bash
git clone https://github.com/nomansum/GojRPCgRPC.git
cd GojRPCgRPC
```

### 2. Create a Kind Cluster



```bash
kind create cluster --config ./k8s/kind-config.yaml

# Or use an existing cluster: kind create cluster --name your-cluster-name
```
### 3. Build Docker Images
```bash
go mod tidy

docker build -f Dockerfile.grpc -t grpc-server:latest .
docker build -f Dockerfile.jrpc -t jrpc-server:latest .
docker build -f Dockerfile.grpcgateway -t grpc-gateway:latest .

```

### 4. Load Images into Kind


```bash
kind load docker-image grpc-server:latest
kind load docker-image jrpc-server:latest
kind load docker-image grpc-gateway:latest
```

### 5. Deploy the Application Services
```bash
kubectl apply -f ./k8s/services/app-namespace.yaml

kubectl apply -f ./k8s/services/grpc-deployment.yaml
kubectl apply -f ./k8s/services/grpc-service.yaml

kubectl apply -f ./k8s/services/grpc-gateway-deployment.yaml
kubectl apply -f ./k8s/services/grpc-gateway-service.yaml

kubectl apply -f ./k8s/services/jrpc-deployment.yaml
kubectl apply -f ./k8s/services/jrpc-service.yaml
```

wait for pods to be ready

```bash
kubectl get pods -n app-ns
```

### 6. Test the APIs

#### JSON-RPC Endpoint (jrpc-server via NodePort 30080)

Create Order

```bash
curl -X POST http://127.0.0.1:30080/rpc \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "method": "CreateOrder",
    "params": {"id": 1022}
  }'
```
Cancel Order 

```bash
curl -X POST http://127.0.0.1:30080/rpc \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "method": "CancelOrder",
    "params": {"id": 1022}
  }'
```

### REST Endpoint (grpc-gateway via NodePort 30081)
Create order (POST):

```bash
curl -X POST http://127.0.0.1:30081/v1/orders \
  -H "Content-Type: application/json" \
  -d '{"id": 1012222}'

```

Cancel Order :

```bash
curl -X DELETE http://127.0.0.1:30081/v1/orders/1012222
```

### 7. Deploy Monitoring Stack

```bash
kubectl apply -f ./k8s/Monitoring/monitoring-namespace.yaml
kubectl apply -f ./k8s/Monitoring/loki-config.yaml
kubectl apply -f ./k8s/Monitoring/loki-deploy-service.yaml

kubectl apply -f ./k8s/Monitoring/prometheus-config.yaml
kubectl apply -f ./k8s/Monitoring/prome-deploy-service.yaml

kubectl apply -f ./k8s/Monitoring/promtail-config.yaml
kubectl apply -f ./k8s/Monitoring/promtail-daemonset.yaml

kubectl apply -f ./k8s/Monitoring/grafana-deploy-service.yaml
```

wait for pods to ready

```bash
kubectl get pods -n monitoring
```


Open browser: http://localhost:3000
Default credentials: admin / admin (you'll be prompted to change the password).
Add Loki Data Source

Configuration → Data Sources → Add data source
Select Loki
URL: http://loki.monitoring.svc.cluster.local:3100
Save & Test

Sample Log Queries in Grafana Explore (Loki)

All logs from app pods: {job="kind-pods"}
Successful order creation: {job="kind-pods"} |= "CreateOrder success"
Cancel order logs: {job="kind-pods"} |= "CancelOrder"

Prometheus is available at http://prometheus.monitoring.svc.cluster.local:9090 (port-forward if needed).
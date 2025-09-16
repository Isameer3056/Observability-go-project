# 🚀 Observability-Go-Project — by Sameer Imtiaz

An end-to-end observability project showcasing **metrics, logs, and traces** with **Go (and Python)**.  
Built and extended by me to learn and demonstrate modern **Site Reliability Engineering (SRE)** practices.

---

## 🌟 Why I Built This

As a **DevOps & SRE enthusiast**, I wanted more than just theory — I wanted to *feel* observability in action.  
So, I extended an open-source Go example into a **full observability playground**, adding:

- 📊 Metrics with **Prometheus & Grafana**  
- 📝 Structured logging with **correlation IDs**  
- 🔎 Distributed tracing across services  
- ⚡ Alerts for **real-time issue detection**  
- 🐍 A Python microservice to prove observability in a **polyglot stack**  
- ☸️ Deployment to **Kubernetes** for production-like realism  


---

## 🏗️ Architecture Overview

```mermaid
graph TD
  A[Go Service] -->|HTTP/gRPC| B[Python FastAPI Service]
  A -->|/metrics| C[Prometheus]
  B -->|/metrics| C
  C --> D[Grafana Dashboards]
  A --> E[Structured Logs -> Loki/Console]
  B --> E
  A --> F[Tracing -> Jaeger]
  B --> F
  ```
-   **Go Service**: Core backend, instrumented for metrics, logs, traces
    
-   **Python Service**: Adds multi-language observability
    
-   **Prometheus**: Scrapes metrics
    
-   **Grafana**: Visualizes dashboards & alerts
    
-   **Jaeger**: Shows request traces
    
-   **Docker Compose / K8s**: Infrastructure

## ⚙️ Environment Setup

### 1. Prerequisites

-   Go ≥ 1.20
    
-   Python ≥ 3.10 (for FastAPI service)
    
-   Docker & Docker Compose
    
-   (Optional) Kubernetes (minikube, kind, or any cluster)
    

### 2. Clone the Repo

`git clone https://github.com/Isameer3056/Observability-Go-Project.git cd Observability-Go-Project` 

### 3. Start the Stack

`docker-compose up -d` 

Services exposed:

-   Prometheus → [http://localhost:9090](http://localhost:9090)
    
-   Grafana → [http://localhost:3000](http://localhost:3000)
    
-   Jaeger → [http://localhost:16686](http://localhost:16686)
    
-   Go app → [http://localhost:8080](http://localhost:8080)
    
-   Python app → [http://localhost:8000](http://localhost:8000)
    

----------

## 🔍 Features I Added

### ✅ Metrics

-   Request counters & error counters
    
-   Latency histograms
    
-   Custom Grafana dashboards
    

### ✅ Logging

-   Structured JSON logging with correlation IDs
    
-   Log enrichment for tracing context
    

### ✅ Tracing

-   End-to-end request traces via Jaeger
    
-   Service-to-service (Go → Python) trace propagation
    

### ✅ Alerts

-   Error rate > 5%
    
-   High latency (> 1s for 95th percentile)
    
-   Alerts routed via Grafana/Alertmanager
    

### ✅ Python Service

-   Simple FastAPI app with observability hooks
    
-   Demonstrates **polyglot observability**
    

### ✅ Kubernetes Deployment

`k8s/` folder with manifests for:

-   Go app
    
-   Python app
    
-   Prometheus + Grafana + Jaeger
    

----------

## 📊 Dashboards

-   Request rate per endpoint
    
-   Error % over time
    
-   Latency heatmap
    
-   Tracing waterfall from Go → Python
    

----------

## 🧩 How to Extend Further

-   🔄 Add database (Postgres/Redis) & observe queries
    
-   ☁️ Deploy to cloud (EKS/GKE/AKS)
    
-   🔔 Integrate Slack/PagerDuty alerts
    
-   💡 Apply chaos testing to see observability in action
    

----------

## 🎯 What I Learned

-   How to instrument **Go & Python services** for observability
    
-   How **metrics/logs/traces complement each other**
    
-   How to build meaningful dashboards for **SRE work**
    
-   How to think in **SLOs & error budgets**

### Stopping the Services
To stop all services, run:

```bash
docker-compose down
```





# kubewatch-api
kubewatch-api 📡 – Watching for deprecated API usage in Kubernetes

## 📖 Overview
kubewatch-api is a lightweight tool that periodically scans your Kubernetes cluster for deprecated API versions and exposes them as Prometheus metrics.

## 🚀 Features
- Detects **deprecated Kubernetes API versions**.
- Exposes **Prometheus-compatible metrics** at `/metrics`.
- Includes **Helm chart for easy deployment**.
- Provides **ServiceMonitor** for integration with Prometheus Operator.

## 📂 Project Structure
```
kubewatch-api/
│── cmd/
│   └── main.go         # Main entry point
│── pkg/
│   ├── monitor/
│   │   ├── monitor.go  # API monitoring logic
│── config/
│── deployment/
│   ├── helm/
│   │   ├── templates/
│   │   │   ├── deployment.yaml
│   │   │   ├── service.yaml
│   │   │   ├── servicemonitor.yaml
│   │   ├── values.yaml
│   │   ├── Chart.yaml
│── Dockerfile          # Docker build file
│── .gitignore          # Ignore files for Git
│── README.md           # Documentation
│── LICENSE             # Open-source license
```

## 🛠 Prerequisites
- **Kubernetes Cluster** (v1.22+ recommended)
- **Prometheus Operator** (for ServiceMonitor integration)
- **Helm** (for deployment)
- **Docker** (for containerization)

## 📦 Installation & Usage

### 1️⃣ Clone Repository
```sh
git clone https://github.com/Binyamse/kubewatch-api.git
cd kubewatch-api
```

### 2️⃣ Build & Run Locally
```sh
export KUBECONFIG=~/.kube/config  # Adjust path if needed
go mod tidy  # Install dependencies
go run cmd/main.go  # Start the application
```
Test it:
```sh
curl http://localhost:8080/metrics
```

### 3️⃣ Build & Run with Docker
```sh
docker build -t kubewatch-api .
docker run -p 8080:8080 -e KUBECONFIG=/root/.kube/config kubewatch-api
```

### 4️⃣ Deploy Using Helm
First, package the Helm chart:
```sh
helm package deployment/helm
```
Deploy it:
```sh
helm install kubewatch-api deployment/helm
```
Check deployment:
```sh
kubectl get pods -l app=kubewatch-api
kubectl port-forward svc/kubewatch-api 8080:8080
curl http://localhost:8080/metrics
```

### 5️⃣ Enable Prometheus Monitoring
The included **ServiceMonitor** assumes **Prometheus Operator** is installed. If not, install it first:
```sh
helm install prometheus prometheus-community/kube-prometheus-stack
```
Then apply the **ServiceMonitor**:
```sh
kubectl apply -f deployment/helm/templates/servicemonitor.yaml
```

### 6️⃣ Setup Prometheus Alerts
Create a **PrometheusRule** for alerting:
```yaml
apiVersion: monitoring.coreos.com/v1
kind: PrometheusRule
metadata:
    name: kubewatch-api-rule
spec:
  groups:
  - name: kubewatch-api.rules
    rules:
    - alert: DeprecatedAPIDetected
      expr: kubewatch_deprecated_api > 0
      for: 5m
      labels:
        severity: warning
      annotations:
        summary: "Deprecated Kubernetes API Detected"
        description: "Some deprecated Kubernetes APIs are still in use. Check /metrics endpoint."
```
Apply it:
```sh
kubectl apply -f prometheus-rule.yaml
```

## ✅ Next Steps
- **[ ]** Add a Grafana dashboard for visualization.
- **[ ]** Add Slack/Teams notifications for alerts.
- **[ ]** Extend API checks for feature gates.

## 📜 License
This project is licensed under the **MIT License**.

---

🚀 **Contributions Welcome!** Fork the repo and submit a PR. Happy monitoring! 🎯

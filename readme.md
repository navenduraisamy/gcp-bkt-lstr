# 📦 GCP Bucket Lister
A simple service that lists all Google Cloud Storage buckets in a given project. Built for learning GKE + Workload Identity + GCP IAM integration.

## 🚀 Features
- Lists GCS buckets using Google Cloud Client Libraries
- Deployed on Google Kubernetes Engine (GKE)
- Uses Workload Identity for secure, keyless auth to GCP APIs


## 🌐 API Endpoints

| Method | Endpoint       | Description                                |
|--------|----------------|--------------------------------------------|
| GET    | `api/gcp/buckets`     | Lists all GCS buckets in the project       |
| GET    | `/api/healthz` | Returns 200 if `PROJECT_ID` env is present |

# Kubernetes Whack-A-Mole

## Backend

> cd backend

Build
> docker build -t gcr.io/whack-a-mole/whack-a-mole-backend:v1.0 .
>
Run locally
> docker run --rm -p 127.0.0.1:8080:8080 gcr.io/whack-a-mole/whack-a-mole-backend:v1.0

## Moles

> cd moles

Build
> docker build -t gcr.io/whack-a-mole/whack-a-mole-mole:v1.0 .
>
Run locally
> docker run --rm -p 127.0.0.1:8080:8080 gcr.io/whack-a-mole/whack-a-mole-mole:v1.0

### Frontend

> cd frontend

Build
> docker build -t gcr.io/whack-a-mole/whack-a-mole-frontend:v1.0 .
>
Run locally
> docker run --rm -p 127.0.0.1:80:8081 gcr.io/whack-a-mole/whack-a-mole-frontend:v1.0

## Deployment

Create whack-a-mole deployment
> kubectl create -f deployments/whack-a-mole.yaml

Create test-moles deployment
> kubectl create -f deployments/test-moles.yaml

To check pods is running
> kubectl get pod

Expose service with port 80 for frontend
> kubectl expose deployment whack-a-mole --type "LoadBalancer" --port 8080

Check service is working
> kubectl get service whack-a-mole
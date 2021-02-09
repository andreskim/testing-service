# testing-service

Example Golang API project that uses OpenAPI spec-first approach and default CI.

## Local development

### Prerequisites

1. Docker Desktop
1. Kubernetes context pointing to current QA cluster acquired using `az aks get-credentials` command
1. Helm3 and access to `ceresacr` repository:

   ```bash
   az acr helm repo add --name ceresacr
   ```

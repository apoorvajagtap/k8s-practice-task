The following tasks are achieved in this project (For Practice):
- Created a simple web app in go that displays a simple text.
- Created Dockerfile for web app.
- Built it's image & pushed to image repo (Quay: quay.io/apoorvaj/helloworldapp:v1)
- Created deployment & service for the web-app (to host application on minikube).

###### Note: One can create the deployment & service either by using the individual yamls (`deployment.yaml` & `service.yaml`), OR using the `deployment_service_template.yaml` (all available in k8s_manifests) in one go. 
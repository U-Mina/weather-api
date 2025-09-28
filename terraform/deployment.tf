resource "kubernetes_deployment" "weather_api" {
  metadata {
    name = "weather-api-deployment"
    labels = {
      app = "weather-api"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "weather-api"
      }
    }

    template {
      metadata {
        labels = {
          app = "weather-api"
        }
      }
      spec {
        container {
          name = "weather-api-container"
          image = "weather-api:v1"
          port {
            container_port = 8080
          }
          env_from {
            secret_ref {
              name = kubernetes_secret.openweathermap.metadata[0].name
            }
          }
        }
      }
    }
  }
}
resource "kubernetes_service" "weather_api" {
  metadata {
    name = "weather-api-service"
  }
  spec {
    selector = {
        app = kubernetes_deployment.weather_api.spec[0].template[0].metadata[0].labels.app
    }
    port {
      port = 80
      target_port = 8080
    }
  }
}
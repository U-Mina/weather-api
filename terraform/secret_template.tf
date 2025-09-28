# resource "kubernetes_secret" "openweathermap" {
#   metadata {
#     name = "openweathermap-secret"
#   }
#   data = {
#     API_KEY = "apiKey"
#   }
#   type = "Opaque"
# }
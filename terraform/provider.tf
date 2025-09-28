terraform {
    required_providers {
        kubernetes = {
            source = "hashicorp/kubernetes"
            version = "~>2.0"
        }
    }
}

provider "kubernetes" {
#   terraform can automatically figure out the kubectl config
    config_path = "~/.kube/config"
}
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.80.0"
    }
  }

  required_version = "1.6.3"
}

provider "azurerm" {
  features {}
}


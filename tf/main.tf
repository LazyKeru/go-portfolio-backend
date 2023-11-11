resource "azurerm_resource_group" "rg" {
  location = var.ressource_group_location
  name     = var.resource_group_name
}

resource "azurerm_service_plan" "backend" {
  name                = "lazykeru-backend"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  os_type             = "Linux"
  sku_name            = "P1v2"
}

resource "azurerm_app_service_custom_hostname_binding" "domain" {
  hostname = "backend.lazykeru.fr"
  app_service_name = azurerm_linux_web_app.backend.name
  resource_group_name = azurerm_resource_group.rg.name
}

resource "azurerm_linux_web_app" "backend" {
  name                = "lazykeru-backend"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  service_plan_id     = azurerm_service_plan.backend.id

  site_config {
    application_stack {
      docker_image_name   = "lazykeru/go-portfolio-backend:v1.0.0"
      docker_registry_url = "https://index.docker.io"
    }
  }
}
variable "resource_group_name" {
    type = string
    default = "go-portfolio-backend"
    description = "Name of teh ressource group."
}

variable "resource_group_location" {
    type = string
    default = "East US"
    description = "Location of the resource group."
}

variable "resource_group_name_prefix" {
  type        = string
  default     = "rg"
  description = "Prefix of the resource group name that's combined with a random ID so name is unique in your Azure subscription."
}
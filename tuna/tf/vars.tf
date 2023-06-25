# Defining Public Key
variable "public_key" {
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDKIMI9J02ZyBGBOpgQrPuIodhEQ5Uf0Snb3xngtYP8+64EV9F5C5IxYvudbFQvr9hShBzF+Q89YvahOketpCJle4Vd2QhtfP1MTP2YFPECZ+tWUOuMWv9Rzu5l3AVk99pJ/HhILHnapawjq6e3g3YQaAsTLP1W2nHdtT1V2YHVoGB0AmuEFc+XjLS1icv9l7rBfpo9cBx89W7uDq7ZEwQSUe1hX4cyZUIyNxP38cxEVxibpkyDa8Pvw/miRcJ+MQ9Ut8898QIujMJniT8720rZQfU3LOsVsLHg0vMUr8aG8z4qeUrnr9DcmAfg6jtqYYvFU7f86WlOAbCucV9IhU9r"
}
# Defining Private Key
variable "private_key" {
  default = "/Users/tomokilam/Docs/Git_Projects/keys/aws"
}
# Definign Key Name for connection
variable "key_name" {
  default = "tests"
  description = "Name of AWS key pair"
}
# Defining CIDR Block for VPC
variable "vpc_cidr" {
  default = "10.0.0.0/16"
}
# Defining CIDR Block for Subnet
variable "subnet_cidr" {
  default = "10.0.1.0/24"
}
# Defining CIDR Block for 2d Subnet
variable "subnet1_cidr" {
  default = "10.0.2.0/24"
}
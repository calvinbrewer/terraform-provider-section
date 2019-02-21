provider "section" {
  username    = "${var.section_username}"
  password  =  "${var.section_password}"
}

resource "section_account" "example" {
  name            = "Example"
  hostname        = "www.example.com"
  origin          = "example.com"
  stack_name      = "varnish-5.1-basic"

}


resource "section_application" "example" {
  hostname        = "www.example.com"
  origin          = "example.com"
  stack_name      = "varnish-5.1-basic"

}


resource "section_environment" "example" {
  name                        = "Example"
  source_environment_name     = "Production"
  domain_name                 = "example.com"

}
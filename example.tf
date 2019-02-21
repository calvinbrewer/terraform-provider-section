provider "section" {
  username    = "${var.section_username}"
  password  =  "${var.section_password}"
}
# Username and password can alternatively be set using the SECTION_USERNAME & SECTION_PASSWORD environment variables

resource "section_account" "myaccount" {
  name       = "My Company"
  hostname   = "www.mysite.example"
  origin     = "origin.mysite.example"
  stack_name = "varnish:6.1.1-basic"
}

resource "section_application" "othersite" {
  hostname   = "www.othersite.example"
  origin     = "origin.othersite.example"
  stack_name = "openresty:1.13.6.1"
  account_id = "${data.section_account.myaccount.id}"
}

resource "section_environment" "staging" {
  name                    = "Staging"
  source_environment_name = "Production"
  domain_name             = "staging.mysite.example"
  account_id              = "${data.section_account.myaccount.id}"
  application_id          = "${data.section_application.othersite.id}"
}
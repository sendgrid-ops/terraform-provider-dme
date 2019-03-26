Terraform Provider
==================

This is a SendGrid hack/fork of the original provider. It does two things:

 * Makes resource `updates` actually work
 * Hard-coded to do absolutely nothing unless `BUILDKITE=true`

Installation
------------

 * Download from [Releases](https://github.com/sendgrid-ops/terraform-provider-dme/releases)
 * Place in `~/.terraform.d/plugins/terraform-provider-dme` - make executable
 * Go to your terraform dir, remove the state `rm -rf .terraform`
 * Run `terraform init` again - you're now on your custom provider

Questions
---------

Hit us up in #autobots-public

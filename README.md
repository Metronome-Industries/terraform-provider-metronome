# Metronome Terraform Provider

The [Metronome Terraform provider](https://registry.terraform.io/providers/Metronome-Industries/metronome/latest/docs) provides convenient access to
the [Metronome REST API](https://docs.metronome.com) from Terraform.

It is generated with [Stainless](https://www.stainless.com/).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Add the following to your `main.tf` file:

<!-- x-release-please-start-version -->

```hcl
# Declare the provider and version
terraform {
  required_providers {
    metronome = {
      source  = "Metronome-Industries/metronome"
      version = "~> 0.1.0-alpha.2"
    }
  }
}

# Initialize the provider
provider "metronome" {
  bearer_token = "My Bearer Token" # or set METRONOME_BEARER_TOKEN env variable
  webhook_secret = "My Webhook Secret" # or set METRONOME_WEBHOOK_SECRET env variable
}

# Configure a resource

```

<!-- x-release-please-end -->

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/Metronome-Industries/metronome/latest/docs).

### Provider Options

When you initialize the provider, the following options are supported. It is recommended to use environment variables for sensitive values like access tokens.
If an environment variable is provided, then the option does not need to be set in the terraform source.

| Property       | Environment variable       | Required | Default value |
| -------------- | -------------------------- | -------- | ------------- |
| bearer_token   | `METRONOME_BEARER_TOKEN`   | true     | —             |
| webhook_secret | `METRONOME_WEBHOOK_SECRET` | false    | —             |

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/Metronome-Industries/terraform-provider-metronome/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).

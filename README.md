# tf-azurerm-module_primitive-application_insights

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This Terraform module creates an Azure Application Insights resource.

## Usage

See [examples/complete](examples/complete) for a complete workspace-based Application Insights example.

## Module Development

### Pre-Requisites

The following commands should be available on your system:

- `asdf` or `mise`
- `make`
- `python3` (for pre-commit)
- Azure CLI

Additionally, configure your Git user and email. Run `make configure` from the repository root to install and verify the required tooling.

### Pre-Commit hooks

[.pre-commit-config.yaml](.pre-commit-config.yaml) defines hooks for Terraform formatting, validation, documentation generation, readonly-test wiring, and secret detection. Hooks are installed by `make configure`. Go linting runs through `make lint` locally and in CI.

### Terratest example

Functional and readonly tests explicitly target `examples/complete`. The readonly suite uses the non-destructive test runner.

### Local Validation

Validate changes locally before pushing them to GitHub:

1. Run `make configure` successfully.
2. Sign in to the appropriate Azure subscription and ensure the test identity can create and delete the example resources.
3. Run the Terraform and Go linters:

```shell
make lint
```

4. Run integration tests, which apply, test, and destroy real Azure resources:

```shell
make test
```

Pre-commit validation, `make lint`, and `make test` also run in CI.

### Review & Merge Process

Open a pull request to the default (`main`) branch. The PR title must follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/#specification) format to merge and drive semantic versioning.

Ensure CI workflows pass, address review feedback, and obtain the approvals required by `CODEOWNERS`.

### Automatic Updates

Shared configuration and workflow files are largely managed through [launch-terraform-skeleton](https://github.com/launchbynttdata/launch-terraform-skeleton). Avoid one-off edits to copied skeleton files unless necessary. Use `copier check-update` or `copier update` when refreshing from the skeleton.
<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.2 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | >= 3.77, < 5.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azurerm_application_insights.app_insights](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/application_insights) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application_type"></a> [application\_type](#input\_application\_type) | Specifies the type of the application - web, ios, java etc. | `string` | `"web"` | no |
| <a name="input_daily_data_cap_in_gb"></a> [daily\_data\_cap\_in\_gb](#input\_daily\_data\_cap\_in\_gb) | Specifies the Application Insights component daily data volume cap in GB. | `number` | `null` | no |
| <a name="input_daily_data_cap_notifications_disabled"></a> [daily\_data\_cap\_notifications\_disabled](#input\_daily\_data\_cap\_notifications\_disabled) | Specifies if a notification email will be send when the daily data volume cap is met. | `bool` | `null` | no |
| <a name="input_disable_ip_masking"></a> [disable\_ip\_masking](#input\_disable\_ip\_masking) | By default the real client IP is masked as 0.0.0.0 in the logs. Use this argument to disable masking and<br/>    log the real client IP. Defaults to false | `bool` | `false` | no |
| <a name="input_force_customer_storage_for_profiler"></a> [force\_customer\_storage\_for\_profiler](#input\_force\_customer\_storage\_for\_profiler) | Should the Application Insights component force users to create their own storage account for profiling?<br/>    Defaults to false. | `bool` | `false` | no |
| <a name="input_internet_ingestion_enabled"></a> [internet\_ingestion\_enabled](#input\_internet\_ingestion\_enabled) | Should the Application Insights component support ingestion over the Public Internet? Defaults to true. | `bool` | `true` | no |
| <a name="input_internet_query_enabled"></a> [internet\_query\_enabled](#input\_internet\_query\_enabled) | Should the Application Insights component support querying over the Public Internet? Defaults to true. | `bool` | `true` | no |
| <a name="input_local_authentication_disabled"></a> [local\_authentication\_disabled](#input\_local\_authentication\_disabled) | Disable Non-Azure AD based Auth. Defaults to false. | `bool` | `false` | no |
| <a name="input_location"></a> [location](#input\_location) | The location/region where the resource will be created | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | The name of the resource | `string` | n/a | yes |
| <a name="input_resource_group_name"></a> [resource\_group\_name](#input\_resource\_group\_name) | The name of the resource group in which the resource will be created | `string` | n/a | yes |
| <a name="input_retention_in_days"></a> [retention\_in\_days](#input\_retention\_in\_days) | Specifies the retention period in days. Possible values are 30, 60, 90, 120, 180, 270, 365, 550 or 730.<br/>    Defaults to 90 | `number` | `30` | no |
| <a name="input_sampling_percentage"></a> [sampling\_percentage](#input\_sampling\_percentage) | Specifies the percentage of the data produced by the application which is sampled for Application Insights.<br/>    Defaults to 100 | `number` | `100` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Tags for the resource | `map(string)` | `{}` | no |
| <a name="input_workspace_id"></a> [workspace\_id](#input\_workspace\_id) | Specifies the id of a log analytics workspace resource. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_app_id"></a> [app\_id](#output\_app\_id) | The application ID associated with the Application Insights component. |
| <a name="output_connection_string"></a> [connection\_string](#output\_connection\_string) | The connection string of the Application Insights component. |
| <a name="output_id"></a> [id](#output\_id) | The resource ID of the Application Insights component. |
| <a name="output_instrumentation_key"></a> [instrumentation\_key](#output\_instrumentation\_key) | The instrumentation key of the Application Insights component. |
| <a name="output_name"></a> [name](#output\_name) | The name of the Application Insights component. |
<!-- END_TF_DOCS -->

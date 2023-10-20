<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.22.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_codeartifact_repository"></a> [codeartifact\_repository](#module\_codeartifact\_repository) | ../.. | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_codeartifact_domain.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/codeartifact_domain) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_domain"></a> [domain](#input\_domain) | The name of the  repository domain. | `string` | n/a | yes |
| <a name="input_repository"></a> [repository](#input\_repository) | The name of the repository to create. | `string` | n/a | yes |
| <a name="input_domain_owner"></a> [domain\_owner](#input\_domain\_owner) | The account number of the AWS account that owns the domain. | `string` | `null` | no |
| <a name="input_description"></a> [description](#input\_description) | The description of the repository. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Key-value map of resource tags. If configured with a provider default\_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | The ARN of the repository. |
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the repository. |
| <a name="output_administrator_account"></a> [administrator\_account](#output\_administrator\_account) | The account number of the AWS account that manages the repository. |
| <a name="output_tags_all"></a> [tags\_all](#output\_tags\_all) | A map of tags assigned to the resource, including those inherited from the provider default\_tags configuration block. |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

output "id" {
  description = "The ARN of the repository."
  value       = aws_codeartifact_repository.this.id
}

output "arn" {
  description = "The ARN of the repository."
  value       = aws_codeartifact_repository.this.arn
}

output "domain" {
  description = "The domain of the repository."
  value       = aws_codeartifact_repository.this.domain
}


output "repository" {
  description = "The name of the repository."
  value       = aws_codeartifact_repository.this.repository
}


output "administrator_account" {
  description = "The account number of the AWS account that manages the repository."
  value       = aws_codeartifact_repository.this.administrator_account
}

output "tags_all" {
  description = "A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block."
  value       = aws_codeartifact_repository.this.tags_all
}

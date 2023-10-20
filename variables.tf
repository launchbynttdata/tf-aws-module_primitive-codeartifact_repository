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

variable "repository" {
  description = "The name of the repository to create."
  type        = string
}

variable "domain" {
  description = "The domain that contains the created repository."
  type        = string
}

variable "domain_owner" {
  description = "The account number of the AWS account that owns the domain."
  type        = string
  default     = null
}

variable "description" {
  description = "The description of the repository."
  type        = string
  default     = null
}

variable "upstream_repository_name" {
  description = "The name of an upstream repository."
  type        = string
  default     = null
}

variable "external_connection_name" {
  description = "The name of the external connection associated with a repository."
  type        = string
  default     = null
}

variable "tags" {
  description = "Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

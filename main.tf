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

resource "aws_codeartifact_repository" "this" {
  repository   = var.repository
  domain       = var.domain
  domain_owner = var.domain_owner
  description  = var.description

  dynamic "upstream" {
    for_each = var.upstream_repository_name != null ? [1] : []
    content {
      repository_name = var.upstream_repository_name
    }
  }
  dynamic "external_connections" {
    for_each = var.external_connection_name != null ? [1] : []
    content {
      external_connection_name = var.external_connection_name
    }
  }

  tags = local.tags
}

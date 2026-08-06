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

output "app_insights_id" {
  description = "The resource ID of the Application Insights component."
  value       = module.app_insights.id
}

output "app_insights_name" {
  description = "The name of the Application Insights component."
  value       = module.app_insights.name
}

output "log_analytics_id" {
  description = "The resource ID of the Log Analytics workspace."
  value       = module.log_analytics_workspace.id
}

output "resource_group_id" {
  description = "The resource ID of the example resource group."
  value       = module.resource_group.id
}

output "resource_group_name" {
  description = "The name of the example resource group."
  value       = module.resource_group.name
}

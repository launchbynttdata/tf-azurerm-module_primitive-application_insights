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
  description = "The resource ID of the Application Insights component."
  value       = azurerm_application_insights.app_insights.id
}

output "name" {
  description = "The name of the Application Insights component."
  value       = azurerm_application_insights.app_insights.name
}

output "app_id" {
  description = "The application ID associated with the Application Insights component."
  value       = azurerm_application_insights.app_insights.app_id
}

output "instrumentation_key" {
  description = "The instrumentation key of the Application Insights component."
  value       = azurerm_application_insights.app_insights.instrumentation_key
  sensitive   = true
}

output "connection_string" {
  description = "The connection string of the Application Insights component."
  value       = azurerm_application_insights.app_insights.connection_string
  sensitive   = true
}

/* 
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * OpenAPI spec version: 0.1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

// Command describes a step performed as part of the build pipeline.
type Command struct {

	// Name of the command, as presented on the command line, or if the command is packaged as a Docker container, as presented to `docker pull`.
	Name string `json:"name,omitempty"`

	// Environment variables set before running this Command.
	Env []string `json:"env,omitempty"`

	// Command-line arguments used when executing this Command.
	Args []string `json:"args,omitempty"`

	// Working directory (relative to project source root) used when running this Command.
	Dir string `json:"dir,omitempty"`

	// Optional unique identifier for this Command, used in wait_for to reference this Command as a dependency.
	Id string `json:"id,omitempty"`

	// The ID(s) of the Command(s) that this Command depends on.
	WaitFor []string `json:"waitFor,omitempty"`
}
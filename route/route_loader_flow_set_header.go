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

package route

import (
	"github.com/lburgazzoli/camel-go/logger"
	"github.com/rs/zerolog"
)

// SetHeaderStepHandler --
func SetHeaderStepHandler(step Step, route *RouteDefinition) (*RouteDefinition, error) {
	impl := struct {
		TypedStep

		Key string `yaml:"key"`
		Val string `yaml:"val"`
	}{}

	err := decodeStep("setheader", step, &impl)
	if err != nil {
		return nil, err
	}

	return route.SetHeader(impl.Key, impl.Val), nil
}

// SetHeadersStepHandler --
func SetHeadersStepHandler(step Step, route *RouteDefinition) (*RouteDefinition, error) {
	headers := make(map[string]interface{})
	for k, v := range step {
		if k != "type" {
			headers[k] = v
		}
	}

	logger.Log(zerolog.DebugLevel, "handle setheaders: step=<%v>", step)
	return route.SetHeaders(headers), nil
}

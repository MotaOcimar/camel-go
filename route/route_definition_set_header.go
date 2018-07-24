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
	"github.com/lburgazzoli/camel-go/api"
	"github.com/lburgazzoli/camel-go/processor"

	zlog "github.com/rs/zerolog/log"
)

// ==========================
//
// Extend RouteDefinition DSL
//
// ==========================

// SetHeader --
func (definition *RouteDefinition) SetHeader(key string, val interface{}) *RouteDefinition {
	d := SetHeaderDefinition{
		parent:   definition,
		children: nil,
		key:      key,
		val:      val,
	}

	definition.AddChild(&d)

	return definition
}

// ==========================
//
// FilterDefinition
//
// ==========================

// SetHeaderDefinition --
type SetHeaderDefinition struct {
	api.ContextAware
	ProcessingNode

	context  api.Context
	parent   *RouteDefinition
	children []Definition

	key string
	val interface{}
}

// SetContext --
func (definition *SetHeaderDefinition) SetContext(context api.Context) {
	definition.context = context
}

// Context --
func (definition *SetHeaderDefinition) Context() api.Context {
	return definition.context
}

// Parent --
func (definition *SetHeaderDefinition) Parent() Definition {
	return definition.parent
}

// Children --
func (definition *SetHeaderDefinition) Children() []Definition {
	return definition.children
}

// Processor ---
func (definition *SetHeaderDefinition) Processor() (api.Processor, error) {
	if definition.key != "" && definition.val != nil {
		p := processor.NewProcessingPipeline(func(exchange api.Exchange) {
			zlog.Info().Msgf("SetHeader: %s=%v", definition.key, definition.val)
			exchange.Headers().Bind(definition.key, definition.val)
		})

		return p, nil
	}

	return nil, nil
}

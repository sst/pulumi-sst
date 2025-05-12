package provider

import (
	"context"
)

type Dynamic struct{}

type DynamicArgs = map[string]interface{}

type DynamicState = map[string]interface{}

// All resources must implement Create at a minimum.
func (Dynamic) Create(ctx context.Context, name string, input DynamicArgs, preview bool) (string, DynamicState, error) {
	state := DynamicState{}
	if preview {
		return name, state, nil
	}
	state["created"] = true
	return name, state, nil
}

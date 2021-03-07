package directives

import "github.com/mrverdant13/dash_buttons/backend/graph/generated"

// NewDirectives creates a GraphQL directives holder.
func NewDirectives() generated.DirectiveRoot {
	return generated.DirectiveRoot{
		AdminAction: adminActionDirective,
	}
}

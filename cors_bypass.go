package api

import (
	alert "github.com/appscode/api/alert/v0.1"
	ci "github.com/appscode/api/ci/v0.1"
	credential "github.com/appscode/api/credential/v0.1"
	db "github.com/appscode/api/db/v0.1"
	kubernetes "github.com/appscode/api/kubernetes/v0.1"
	namespace "github.com/appscode/api/namespace/v0.1"
	"github.com/gengo/grpc-gateway/runtime"
)

// This is a hackish method to add support javascript
// CORS request to the grpc-gateway.
// TODO: Make those auto generated.
func Patterens() []runtime.Pattern {
	ps := make([]runtime.Pattern, 0)

	ps = append(ps, alert.Patterns()...)
	ps = append(ps, credential.Patterns()...)
	ps = append(ps, namespace.Patterns()...)
	ps = append(ps, kubernetes.Patterns()...)
	ps = append(ps, db.Patterns()...)
	ps = append(ps, ci.Patterns()...)

	return ps
}

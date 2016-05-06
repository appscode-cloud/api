package loadbalancer

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Loadbalancers_Create_0)
	patterns = append(patterns, pattern_Loadbalancers_Delete_0)
	patterns = append(patterns, pattern_Loadbalancers_Describe_0)
	patterns = append(patterns, pattern_Loadbalancers_List_0)
	patterns = append(patterns, pattern_Loadbalancers_Update_0)
	return patterns
}

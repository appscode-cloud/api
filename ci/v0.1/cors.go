package ci

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Slaves_List_0)
	patterns = append(patterns, pattern_Slaves_Create_0)
	patterns = append(patterns, pattern_Slaves_Delete_0)
	patterns = append(patterns, pattern_Slaves_Describe_0)
	patterns = append(patterns, pattern_Slaves_Restart_0)
	return patterns
}

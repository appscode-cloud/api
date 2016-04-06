package db

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Databases_Create_0)
	patterns = append(patterns, pattern_Databases_Delete_0)
	patterns = append(patterns, pattern_Databases_Backup_0)
	patterns = append(patterns, pattern_Databases_Restore_0)
	patterns = append(patterns, pattern_Databases_Describe_0)
	patterns = append(patterns, pattern_Databases_List_0)
	patterns = append(patterns, pattern_Slaves_Add_0)
	return patterns
}
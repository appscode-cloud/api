package pv

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)

	patterns = append(patterns, pattern_Disks_List_0)
	patterns = append(patterns, pattern_Disks_Create_0)
	patterns = append(patterns, pattern_Disks_Delete_0)

//	patterns = append(patterns, pattern_PersistentVolumes_List_0)
	patterns = append(patterns, pattern_PersistentVolumes_Register_0)
	patterns = append(patterns, pattern_PersistentVolumes_Unregister_0)

//	patterns = append(patterns, pattern_PersistentVolumeClaims_List_0)
	patterns = append(patterns, pattern_PersistentVolumeClaims_Register_0)
	patterns = append(patterns, pattern_PersistentVolumeClaims_Unregister_0)

	return patterns
}

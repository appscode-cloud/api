package certificate

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_TLSCertificates_Create_0)
	patterns = append(patterns, pattern_TLSCertificates_Delete_0)
	patterns = append(patterns, pattern_TLSCertificates_Deploy_0)
	patterns = append(patterns, pattern_TLSCertificates_Describe_0)
	patterns = append(patterns, pattern_TLSCertificates_List_0)
	return patterns
}

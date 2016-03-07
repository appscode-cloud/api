package dtypes

import (
	"errors"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	google_protobuf "github.com/golang/protobuf/ptypes/any"
	"strconv"
)

// Helper methods for status object.
const (
	StatusCodeOK string = "0"
)

// returns the status code string of the response status.
func (s *Status) Status() string {
	code, err := strconv.Atoi(s.Code)
	if err != nil {
		code = int(StatusCode_BADREQUEST)
	}
	return proto.EnumName(StatusCode_name, int32(code))
}

func (s *Status) IsOK() bool {
	if s.Code == StatusCodeOK {
		return true
	}
	return false
}

func (s *Status) IsErr() bool {
	if s.Code != StatusCodeOK {
		return true
	}
	return false
}

func (s *Status) Error() error {
	return errors.New(s.Message)
}

func (s *Status) ErrorMessage() string {
	return s.Message
}

// Adds any proto message in the details field of the Status message.
// This uses google.protobuf.any to to hold and retried data.
func (a *Status) AddDetails(v ...proto.Message) {
	if len(a.Details) == 0 {
		a.Details = make([]*google_protobuf.Any, 0)
	}
	for _, val := range v {
		value, err := proto.Marshal(val)
		if err != nil {
			glog.V(1).Infoln("Marshaling any failed.")
			continue
		}
		anyValue := &google_protobuf.Any{
			TypeUrl: proto.MessageName(val),
			Value:   value,
		}
		a.Details = append(a.Details, anyValue)
	}
}

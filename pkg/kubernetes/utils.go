package kubernetes

import (
	"encoding/json"
	"fmt"

	"github.com/mlycore/log"

	"k8s.io/api/admission/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewAdmissionReview will make any AdmissionReview
func NewAdmissionReview(allow bool, message string) *v1.AdmissionReview {
	return &v1.AdmissionReview{
		Response: &v1.AdmissionResponse{
			Allowed: allow,
			Result: &meta_v1.Status{
				Message: message,
			},
		},
	}
}

// NewAdmissionReview will make any AdmissionReview with patch data
func NewAdmissionReviewWithPatch(allow bool, message string, patch []byte) *v1.AdmissionReview {
	return &v1.AdmissionReview{
		Response: &v1.AdmissionResponse{
			Allowed: allow,
			Patch:   patch,
			Result: &meta_v1.Status{
				Message: message,
			},
		},
	}
}

// EncodeAdmissionReview is used for json marshal the admission review response
func EncodeAdmissionReview(ar *v1.AdmissionReview) []byte {
	resp, err := json.Marshal(ar)
	if err != nil {
		log.Errorf("response marshal error: %s", err)
		return nil
	}
	return resp
}

// NewAdmissionReviewError used for make error AdmissionReview quickly
func NewAdmissionReviewError(err error) []byte {
	ar := NewAdmissionReview(false, fmt.Sprintf("%s", err))
	return EncodeAdmissionReview(ar)
}

// NewAdmissionReviewError used for make ok AdmissionReview with a message
func NewAdmissionReviewOkWithMessage(err error) []byte {
	ar := NewAdmissionReview(true, fmt.Sprintf("%s", err))
	return EncodeAdmissionReview(ar)
}

// NewAdmissionReviewOk used for make ok AdmissionReview
func NewAdmissionReviewOk() []byte {
	ar := NewAdmissionReview(true, "")
	return EncodeAdmissionReview(ar)
}

package main

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"

	"github.com/emicklei/go-restful"
	"github.com/mlycore/log"
	admission "k8s.io/api/admission/v1"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/SphereEx/database-mesh/pkg/kubernetes"
)

// Injector defines a Sidecar injector
type Injector struct {
	//TODO: Config
}

// Do sidecar injection
func (inj *Injector) Do(req *restful.Request, resp *restful.Response) {
	ar := &admission.AdmissionReview{}
	err := req.ReadEntity(&ar)
	if err != nil {
		log.Errorf("read entity error: %s", err)
		data := kubernetes.NewAdmissionReviewError(err)
		resp.Write(data)
		return
	}

	if isPodAdmissionReviewRequest(ar) {
		pod := &v1.Pod{}
		if err := json.Unmarshal(ar.Request.Object.Raw, &pod); err != nil {
			log.Errorf("unmarshal error: %s", err)
			data := kubernetes.NewAdmissionReviewError(err)
			resp.Write(data)
			return
		}

		patchBytes, err := inj.do(pod)
		if err != nil {
			log.Errorf("mutation error: %s", err)
			data := kubernetes.NewAdmissionReviewOkWithMessage(err)
			resp.Write(data)
			return
		}

		ar.Response = &admission.AdmissionResponse{
			Allowed: true,
			Patch:   patchBytes,
			PatchType: func() *admission.PatchType {
				pt := admission.PatchTypeJSONPatch
				return &pt
			}(),
		}

		data, err := json.Marshal(ar)
		if err != nil {
			log.Errorf("mutation marshal error: %s", err)
			data := kubernetes.NewAdmissionReviewOkWithMessage(err)
			resp.Write(data)
			return
		}

		resp.Write(data)
		log.Infof("admission review succeed")
		return
	}

	data := kubernetes.NewAdmissionReviewOk()
	resp.Write(data)
	log.Tracef("It is not a Pod admission request: %s/%s", ar.Request.Namespace, ar.Request.Name)
	return
}

func (inj *Injector) do(pod *v1.Pod) ([]byte, error) {
	var (
		injected bool
		patch    []Operation
	)
	for _, container := range pod.Spec.Containers {
		if container.Name == "database-mesh" {
			injected = true
		}
	}

	if !injected {
		patch = append(patch, addContainers(pod.Spec.Containers, []v1.Container{{
			Name:  "database-mesh",
			Image: os.Getenv("IMAGE"),
		}}, "/spec/containers")...)
		return json.Marshal(patch)
	}

	return nil, errors.New("already injected")
}

func addContainers(current, containers []v1.Container, base string) []Operation {
	var (
		patches = []Operation{}
		value   interface{}
		path    string
		first   = len(current) == 0
	)
	for _, c := range containers {
		if first {
			first = false
			path = base
			value = []v1.Container{c}
		} else {
			path = path + "/-"
			value = c
		}
		patches = append(patches, Operation{
			Op:    "add",
			Path:  path,
			Value: value,
		})
	}
	return patches
}

// Verify if it is a AdmissionReviewRequest about Pods
func isPodAdmissionReviewRequest(request *admission.AdmissionReview) bool {
	podResource := meta_v1.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	return reflect.DeepEqual(request.Request.Resource, podResource)
}

// Operation refers to patch operation
type Operation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

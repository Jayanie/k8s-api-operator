// Copyright (c)  WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// WSO2 Inc. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package errors

import "fmt"

type Reason string

const (
	AnnotationNotExists Reason = "AnnotationNotExists"
	InvalidContent      Reason = "InvalidContent"
)

// Ingress defines errors with a reason
type Ingress interface {
	Reason() Reason
}

// IngressError represents errors for ingresses
type IngressError struct {
	ErrReason Reason
	Message   string
}

func (e IngressError) Error() string {
	return e.Message
}

func (e IngressError) Reason() Reason {
	return e.ErrReason
}

// NewAnnotationNotExists returns a new IngressError with error Reason AnnotationNotExists
func NewAnnotationNotExists(name string) IngressError {
	return IngressError{
		ErrReason: AnnotationNotExists,
		Message:   fmt.Sprintf("Annotation '%s' is not provided", name),
	}
}
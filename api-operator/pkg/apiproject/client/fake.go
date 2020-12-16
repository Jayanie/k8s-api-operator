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

package client

import (
	"context"
	"github.com/wso2/k8s-api-operator/api-operator/pkg/apiproject/build"
	"github.com/wso2/k8s-api-operator/api-operator/pkg/controller/common"
	"math/rand"
	"time"
)

// Fake Adapter client for testing
type Fake struct {
	ProjectMap     *build.ProjectsMap
	Response       Response
	responseMethod func(projects *build.ProjectsMap) Response
}

// NewFake returns a Fake client which returns given response
func NewFake(response Response) *Fake {
	return &Fake{
		responseMethod: func(projects *build.ProjectsMap) Response {
			return response
		},
	}
}

// NewFakeAllSucceeded returns a Fake client which returns success of all projects
func NewFakeAllSucceeded() *Fake {
	return &Fake{
		responseMethod: func(projects *build.ProjectsMap) Response {
			r := Response{}

			for name, project := range *projects {
				switch project.Action {
				case build.ForceUpdate:
					r[name] = Updated
				case build.Delete:
					r[name] = Deleted
				}
			}
			return r
		},
	}
}

// NewFakeAllFailed returns a Fake client which returns of all projects updating failure
func NewFakeAllFailed() *Fake {
	return &Fake{
		responseMethod: func(projects *build.ProjectsMap) Response {
			r := Response{}

			for name := range *projects {
				r[name] = Failed
			}
			return r
		},
	}
}

// NewFakeWithRandomResponse returns a Fake client which returns random failure and success of updating projects
func NewFakeWithRandomResponse() *Fake {
	return &Fake{
		responseMethod: func(projects *build.ProjectsMap) Response {
			r := Response{}
			rand.Seed(time.Now().UnixNano())

			for name, project := range *projects {
				if rand.Intn(2) == 0 {
					r[name] = Failed
				} else {
					switch project.Action {
					case build.ForceUpdate:
						r[name] = Updated
					case build.Delete:
						r[name] = Deleted
					}
				}
			}
			return r
		},
	}
}

func (c *Fake) Update(ctx context.Context, reqInfo *common.RequestInfo, projects *build.ProjectsMap) (Response, error) {
	// TODO: (renuka) delete following comments after testing
	//for s, project := range *projects {
	//	fmt.Println("")
	//	fmt.Println("******* PRINT PROJECT ******")
	//	fmt.Printf("Project name: %s\n", s)
	//	fmt.Printf("Action: %s\n", project.Action)
	//
	//	if project.Action != build.ForceUpdate {
	//		continue
	//	}
	//	err := project.OAS.Validate(ctx)
	//	fmt.Printf("Swagger validation: %v\n", err == nil)
	//	fmt.Printf("Tls certs: %v\n", project.TlsCertificate)
	//	fmt.Println(swagger.PrettyString(project.OAS))
	//}

	c.ProjectMap = projects
	c.Response = c.responseMethod(projects)
	return c.Response, nil
}
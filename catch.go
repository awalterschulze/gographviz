//Copyright 2017 GoGraphviz Authors
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package gographviz

import (
	"fmt"
	"strings"
)

type errInterface interface {
	SetStrict(strict bool)
	SetDir(directed bool)
	SetName(name string)
	AddPortEdge(src, srcPort, dst, dstPort string, directed bool, attrs map[string]string)
	AddEdge(src, dst string, directed bool, attrs map[string]string)
	AddNode(parentGraph string, name string, attrs map[string]string)
	AddAttr(parentGraph string, field, value string)
	AddSubGraph(parentGraph string, name string, attrs map[string]string)
	String() string
	getError() error
}

func newErrCatcher(g Interface) errInterface {
	return &errCatcher{g, nil}
}

type errCatcher struct {
	Interface
	errs []error
}

func (this *errCatcher) AddAttr(parentGraph string, field, value string) {
	if err := this.Interface.AddAttr(parentGraph, field, value); err != nil {
		this.errs = append(this.errs, err)
	}
}

func (this *errCatcher) AddSubGraph(parentGraph string, name string, attrs map[string]string) {
	if err := this.Interface.AddSubGraph(parentGraph, name, attrs); err != nil {
		this.errs = append(this.errs, err)
	}
}

func (this *errCatcher) AddNode(parentGraph string, name string, attrs map[string]string) {
	if err := this.Interface.AddNode(parentGraph, name, attrs); err != nil {
		this.errs = append(this.errs, err)
	}
}

func (this *errCatcher) getError() error {
	if len(this.errs) == 0 {
		return nil
	}
	ss := make([]string, len(this.errs))
	for i, err := range this.errs {
		ss[i] = err.Error()
	}
	return fmt.Errorf("errors: [%s]", strings.Join(ss, ","))
}

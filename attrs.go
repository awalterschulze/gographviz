//Copyright 2013 GoGraphviz Authors
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
	"sort"
)

//Represents attributes for an Edge, Node or Graph.
type Attrs map[Attr]string

//Creates an empty Attributes type.
func NewAttrs(m map[string]string) (Attrs, error) {
	as := make(Attrs)
	for k, v := range m {
		if err := as.Add(k, v); err != nil {
			return nil, err
		}
	}
	return as, nil
}

func NewAttr(key string) (Attr, error) {
	a, ok := validAttrs[key]
	if !ok {
		return Attr(""), fmt.Errorf("%s is not a valid attribute", key)
	}
	return a, nil
}

//Adds an attribute name and value.
func (this Attrs) Add(field string, value string) error {
	a, err := NewAttr(field)
	if err != nil {
		return err
	}
	this.add(a, value)
	return nil
}

func (this Attrs) add(field Attr, value string) {
	this[field] = value
}

//Adds the attributes into this Attrs type overwriting duplicates.
func (this Attrs) Extend(more Attrs) {
	for key, value := range more {
		this.add(key, value)
	}
}

//Only adds the missing attributes to this Attrs type.
func (this Attrs) Ammend(more Attrs) {
	for key, value := range more {
		if _, ok := this[key]; !ok {
			this.add(key, value)
		}
	}
}

func (this Attrs) ToMap() map[string]string {
	m := make(map[string]string)
	for k, v := range this {
		m[string(k)] = v
	}
	return m
}

type attrList []Attr

func (this attrList) Len() int { return len(this) }
func (this attrList) Less(i, j int) bool {
	return this[i] < this[j]
}
func (this attrList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Attrs) SortedNames() []Attr {
	keys := make(attrList, 0)
	for key := range this {
		keys = append(keys, key)
	}
	sort.Sort(keys)
	return []Attr(keys)
}

func (this Attrs) Copy() Attrs {
	mm := make(Attrs)
	for k, v := range this {
		mm[k] = v
	}
	return mm
}

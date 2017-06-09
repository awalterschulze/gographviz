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
	"testing"
)

// https://github.com/awalterschulze/gographviz/issues/32
func TestIssue32DefaultAttrs(t *testing.T) {
	inputString := `
	digraph G {
		node [shape=record, fillcolor=red, style=filled];
		A [shape=circle];
		A -> B;
		B [fillcolor=blue];
	}
	`

	g, err := Read([]byte(inputString))
	if err != nil {
		t.Fatal(err)
	}

	type nodeCase struct {
		name  string
		attr  Attr
		value string
	}
	for _, c := range []nodeCase{
		{"A", "shape", "circle"},
		{"A", "fillcolor", "red"},
		{"B", "shape", "record"},
		// Node declarations after the nodes' edges
		// still override the default attributes.
		{"B", "fillcolor", "blue"},
	} {
		node := g.Nodes.Lookup[c.name]
		if value := node.Attrs[c.attr]; value != c.value {
			t.Errorf("expected %s=%s in %s; got %s", c.attr, c.value, c.name, value)
		}
	}
}

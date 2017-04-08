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
	"reflect"
	"testing"
)

func TestEdges_Sorted(t *testing.T) {
	var tts = map[string]struct {
		edges    []*Edge
		expected []*Edge
	}{
		"empty": {
			edges:    []*Edge{},
			expected: []*Edge{},
		},
		"one edge": {
			edges: []*Edge{
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
			},
			expected: []*Edge{
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
			},
		},
		"two non parallel edges": {
			edges: []*Edge{
				{Src: "0", Dst: "2", Attrs: Attrs{LABEL: "hello"}},
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
			},
			expected: []*Edge{
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
				{Src: "0", Dst: "2", Attrs: Attrs{LABEL: "hello"}},
			},
		},
		"two parallel edges": {
			edges: []*Edge{
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "hello"}},
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
			},
			expected: []*Edge{
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "hello"}},
			},
		},
		"two parallel edges - one without label": {
			edges: []*Edge{
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
				{Src: "0", Dst: "1"},
			},
			expected: []*Edge{
				{Src: "0", Dst: "1"},
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
			},
		},
		"several non parallel edges": {
			edges: []*Edge{
				{Src: "0", Dst: "2", Attrs: Attrs{LABEL: "hello"}},
				{Src: "1", Dst: "1", Attrs: Attrs{LABEL: "world"}},
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
				{Src: "1", Dst: "0", Attrs: Attrs{LABEL: "golang"}},
			},
			expected: []*Edge{
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
				{Src: "0", Dst: "2", Attrs: Attrs{LABEL: "hello"}},
				{Src: "1", Dst: "0", Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "1", Attrs: Attrs{LABEL: "world"}},
			},
		},
		"several with parallel edges": {
			edges: []*Edge{
				{Src: "0", Dst: "2", Attrs: Attrs{LABEL: "hello"}},
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "cba"}},
				{Src: "1", Dst: "1", Attrs: Attrs{LABEL: "world"}},
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
				{Src: "1", Dst: "0", Attrs: Attrs{LABEL: "gopher"}},
				{Src: "1", Dst: "0", Attrs: Attrs{LABEL: "golang"}},
			},
			expected: []*Edge{
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "cba"}},
				{Src: "0", Dst: "2", Attrs: Attrs{LABEL: "hello"}},
				{Src: "1", Dst: "0", Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", Attrs: Attrs{LABEL: "gopher"}},
				{Src: "1", Dst: "1", Attrs: Attrs{LABEL: "world"}},
			},
		},
		"edges with ports": {
			edges: []*Edge{
				{Src: "0", Dst: "1", SrcPort: "a", DstPort: "b"},
				{Src: "0", Dst: "1", SrcPort: "a", DstPort: "a"},
				{Src: "0", Dst: "1", SrcPort: "b", DstPort: "a"},
			},
			expected: []*Edge{
				{Src: "0", Dst: "1", SrcPort: "a", DstPort: "a"},
				{Src: "0", Dst: "1", SrcPort: "a", DstPort: "b"},
				{Src: "0", Dst: "1", SrcPort: "b", DstPort: "a"},
			},
		},
		"directed edges before non directed edges": {
			edges: []*Edge{
				{Src: "0", Dst: "1", Dir: false},
				{Src: "0", Dst: "1", Dir: true},
			},
			expected: []*Edge{
				{Src: "0", Dst: "1", Dir: true},
				{Src: "0", Dst: "1", Dir: false},
			},
		},
		"the theory of everything": {
			edges: []*Edge{
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "cba"}},
				{Src: "1", Dst: "0", SrcPort: "a", Dir: false, Attrs: Attrs{LABEL: "gopher"}},
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
				{Src: "0", Dst: "2", Attrs: Attrs{LABEL: "hello"}},
				{Src: "1", Dst: "0", Attrs: Attrs{LABEL: "gopher"}},
				{Src: "1", Dst: "0", Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", SrcPort: "b", Attrs: Attrs{LABEL: "gopher"}},
				{Src: "1", Dst: "0", SrcPort: "a", DstPort: "b", Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "1", Attrs: Attrs{"comment": "test", LABEL: "world"}},
				{Src: "1", Dst: "0", SrcPort: "a", DstPort: "a", Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", SrcPort: "a", Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", SrcPort: "b", Dir: false, Attrs: Attrs{LABEL: "gopher"}},
				{Src: "1", Dst: "1", Attrs: Attrs{LABEL: "world"}},
				{Src: "1", Dst: "0", SrcPort: "a", DstPort: "b", Dir: true, Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "1", Attrs: Attrs{"comment": "test", LABEL: "hello"}},
				{Src: "1", Dst: "0", SrcPort: "a", Dir: true, Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", SrcPort: "a", DstPort: "b", Dir: true, Attrs: Attrs{LABEL: "graphviz"}},
			},
			expected: []*Edge{
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "abc"}},
				{Src: "0", Dst: "1", Attrs: Attrs{LABEL: "cba"}},
				{Src: "0", Dst: "2", Attrs: Attrs{LABEL: "hello"}},
				{Src: "1", Dst: "0", Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", Attrs: Attrs{LABEL: "gopher"}},
				{Src: "1", Dst: "0", SrcPort: "a", Dir: true, Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", SrcPort: "a", Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", SrcPort: "a", Dir: false, Attrs: Attrs{LABEL: "gopher"}},
				{Src: "1", Dst: "0", SrcPort: "a", DstPort: "a", Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", SrcPort: "a", DstPort: "b", Dir: true, Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", SrcPort: "a", DstPort: "b", Dir: true, Attrs: Attrs{LABEL: "graphviz"}},
				{Src: "1", Dst: "0", SrcPort: "a", DstPort: "b", Attrs: Attrs{LABEL: "golang"}},
				{Src: "1", Dst: "0", SrcPort: "b", Dir: false, Attrs: Attrs{LABEL: "gopher"}},
				{Src: "1", Dst: "0", SrcPort: "b", Attrs: Attrs{LABEL: "gopher"}},
				{Src: "1", Dst: "1", Attrs: Attrs{LABEL: "world"}},
				{Src: "1", Dst: "1", Attrs: Attrs{"comment": "test", LABEL: "hello"}},
				{Src: "1", Dst: "1", Attrs: Attrs{"comment": "test", LABEL: "world"}},
			},
		},
	}

	for name, tt := range tts {
		edges := NewEdges()
		for _, e := range tt.edges {
			edges.Add(e)
		}
		s := edges.Sorted()
		if !reflect.DeepEqual(tt.expected, s) {
			t.Errorf("%s - Sorted invalid: expected %v got %v", name, tt.expected, s)
		} else if !reflect.DeepEqual(edges.Edges, tt.edges) {
			t.Errorf("%s - Sorted should not have changed original order: expected %v got %v", name, tt.edges, edges.Edges)
		}
	}
}

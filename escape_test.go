//Copyright 2013 Vastech SA (PTY) LTD
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
	"strings"
	"testing"
)

func TestEscape(t *testing.T) {
	g := NewEscape()
	g.SetName("asdf adsf")
	g.SetDir(true)
	g.AddNode("asdf asdf", "kasdf99 99", map[string]string{
		"label": "a << b",
	})
	g.AddNode("asdf asdf", "7", map[string]string{
		"label": "&red",
	})
	g.AddNode("asdf asdf", "a << b", nil)
	g.AddEdge("kasdf99 99", "7", true, nil)
	g.AddNode("asdf asdf", "Weird Test Case", map[string]string{
		"label": "< starts with < symbol. This should be fixed in the future",
	})
	s := g.String()
	if !strings.HasPrefix(s, `digraph "asdf adsf" {
	"kasdf99 99"->7;
	"Weird Test Case" [ label="< starts with < symbol. This should be fixed in the future" ];
	"a &lt;&lt; b";
	"kasdf99 99" [ label="a &lt;&lt; b" ];
	7 [ label="&amp;red" ];

}`) {
		t.Fatalf("%s", s)
	}
	if !g.IsNode("a << b") {
		t.Fatalf("should be a node")
	}
}

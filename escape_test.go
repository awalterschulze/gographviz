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

	attrs := NewAttrs()
	attrs.Add(URL, "<asfd") // note not actually a URL
	g.AddNode("asdf asdf", "kasdf99 99", attrs)
	g.AddNode("asdf asdf", "7", attrs)
	g.AddNode("asdf asdf", "a << b", nil)
	g.AddEdge("kasdf99 99", "7", true, nil)
	s := g.String()
	if !strings.HasPrefix(s, `digraph "asdf adsf" {
	"kasdf99 99"->7;
	"a &lt;&lt; b";
	"kasdf99 99" [ URL="<asfd" ];
	7 [ URL="<asfd" ];

}`) {
		t.Fatalf("%s", s)
	}
	if !g.IsNode("a << b") {
		t.Fatalf("should be a node")
	}
}

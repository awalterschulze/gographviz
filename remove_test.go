//Copyright 2018 GoGraphviz Authors
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

func TestRemoveNode(t *testing.T) {
	g := NewGraph()
	if err := g.SetName("G"); err != nil {
		t.Fatal(err)
	}
	if err := g.SetDir(true); err != nil {
		t.Fatal(err)
	}
	if err := g.AddNode("G", "Hello", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddNode("G", "World", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddEdge("Hello", "World", true, nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddSubGraph("G", "cluster_Core", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddNode("cluster_Core", "CoreFunction", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddSubGraph("G", "cluster_Supporting", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddNode("cluster_Supporting", "SupportingFunction", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddSubGraph("G", "cluster_Development", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddNode("cluster_Development", "DevelopmentFunction", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddPortEdge("DevelopmentFunction", "cluster_Development", "CoreFunction", "cluster_Core", true, nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddPortEdge("SupportingFunction", "cluster_Supporting", "CoreFunction", "cluster_Core", true, nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddPortEdge("Hello", "", "CoreFunction", "cluster_Core", true, nil); err != nil {
		t.Fatal(err)
	}

	expected := `digraph G {
	DevelopmentFunction:cluster_Development->CoreFunction:cluster_Core;
	Hello->CoreFunction:cluster_Core;
	Hello->World;
	SupportingFunction:cluster_Supporting->CoreFunction:cluster_Core;
	subgraph cluster_Core {
		CoreFunction;
	};
	subgraph cluster_Development {
		DevelopmentFunction;
	};
	subgraph cluster_Supporting {
		SupportingFunction;
	};
	Hello;
	World;
}
`

	if got := g.String(); got != expected {
		t.Fatalf("output is not expected")
	}

	if err := g.RemoveNode("cluster_Development", "DevelopmentFunction"); err != nil {
		t.Fatal(err)
	}

	expected = `digraph G {
	Hello->CoreFunction:cluster_Core;
	Hello->World;
	SupportingFunction:cluster_Supporting->CoreFunction:cluster_Core;
	subgraph cluster_Core {
		CoreFunction;
	};
	subgraph cluster_Development {

	};
	subgraph cluster_Supporting {
		SupportingFunction;
	};
	Hello;
	World;
}
`
	if got := g.String(); got != expected {
		t.Fatalf("output is not expected")
	}

	if err := g.RemoveNode("G", "Hello"); err != nil {
		t.Fatal(err)
	}
	if err := g.RemoveNode("G", "World"); err != nil {
		t.Fatal(err)
	}

	expected = `digraph G {
	SupportingFunction:cluster_Supporting->CoreFunction:cluster_Core;
	subgraph cluster_Core {
		CoreFunction;
	};
	subgraph cluster_Development {

	};
	subgraph cluster_Supporting {
		SupportingFunction;
	};
}
`
	if g.String() != expected {
		t.Fatalf("output is not expected")
	}

	if err := g.RemoveNode("cluster_Supporting", "SupportingFunction"); err != nil {
		t.Fatal(err)
	}

	expected = `digraph G {
	subgraph cluster_Core {
		CoreFunction;
	};
	subgraph cluster_Development {

	};
	subgraph cluster_Supporting {

	};
}
`

	if g.String() != expected {
		t.Fatalf("output is not expected")
	}

	if err := g.RemoveNode("cluster_Core", "CoreFunction"); err != nil {
		t.Fatal(err)
	}

	expected = `digraph G {
	subgraph cluster_Core {

	};
	subgraph cluster_Development {

	};
	subgraph cluster_Supporting {

	};
}
`

	if g.String() != expected {
		t.Fatalf("output is not expected")
	}
}

func TestRemoveSubGraph(t *testing.T) {
	g := NewGraph()
	if err := g.SetName("G"); err != nil {
		t.Fatal(err)
	}
	if err := g.SetDir(true); err != nil {
		t.Fatal(err)
	}
	if err := g.AddNode("G", "Hello", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddNode("G", "World", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddEdge("Hello", "World", true, nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddSubGraph("G", "cluster_Core", map[string]string{
		"label": "Core",
	}); err != nil {
		t.Fatal(err)
	}
	if err := g.AddNode("cluster_Core", "CoreFunction", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddSubGraph("G", "cluster_Supporting", map[string]string{
		"label": "Supporting",
	}); err != nil {
		t.Fatal(err)
	}
	if err := g.AddNode("cluster_Supporting", "SupportingFunction", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddSubGraph("G", "cluster_Development", map[string]string{
		"label": "Development",
	}); err != nil {
		t.Fatal(err)
	}
	if err := g.AddNode("cluster_Development", "DevelopmentFunction", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddPortEdge("DevelopmentFunction", "cluster_Development", "CoreFunction", "cluster_Core", true, nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddPortEdge("SupportingFunction", "cluster_Supporting", "CoreFunction", "cluster_Core", true, nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddPortEdge("Hello", "", "CoreFunction", "cluster_Core", true, nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddSubGraph("G", "cluster_FooBar", nil); err != nil {
		t.Fatal(err)
	}
	if err := g.AddEdge("Hello", "cluster_FooBar", true, nil); err != nil {
		t.Fatal(err)
	}

	expected := `digraph G {
	DevelopmentFunction:cluster_Development->CoreFunction:cluster_Core;
	Hello->CoreFunction:cluster_Core;
	Hello->World;
	Hello->cluster_FooBar;
	SupportingFunction:cluster_Supporting->CoreFunction:cluster_Core;
	subgraph cluster_Core {
		label=Core;
		CoreFunction;
	};
	subgraph cluster_Development {
		label=Development;
		DevelopmentFunction;
	};
	subgraph cluster_FooBar {

	};
	subgraph cluster_Supporting {
		label=Supporting;
		SupportingFunction;
	};
	Hello;
	World;
}
`
	if got := g.String(); got != expected {
		t.Fatalf("output is not expected")
	}

	if err := g.RemoveSubGraph("G", "cluster_Development"); err != nil {
		t.Fatal(err)
	}

	expected = `digraph G {
	Hello->CoreFunction:cluster_Core;
	Hello->World;
	Hello->cluster_FooBar;
	SupportingFunction:cluster_Supporting->CoreFunction:cluster_Core;
	subgraph cluster_Core {
		label=Core;
		CoreFunction;
	};
	subgraph cluster_FooBar {

	};
	subgraph cluster_Supporting {
		label=Supporting;
		SupportingFunction;
	};
	Hello;
	World;
}
`
	if got := g.String(); got != expected {
		t.Fatalf("output is not expected")
	}

	if err := g.RemoveSubGraph("G", "cluster_Supporting"); err != nil {
		t.Fatal(err)
	}

	expected = `digraph G {
	Hello->CoreFunction:cluster_Core;
	Hello->World;
	Hello->cluster_FooBar;
	subgraph cluster_Core {
		label=Core;
		CoreFunction;
	};
	subgraph cluster_FooBar {

	};
	Hello;
	World;
}
`
	if got := g.String(); got != expected {
		t.Fatalf("output is not expected")
	}

	if err := g.RemoveSubGraph("G", "cluster_Core"); err != nil {
		t.Fatal(err)
	}

	expected = `digraph G {
	Hello->World;
	Hello->cluster_FooBar;
	subgraph cluster_FooBar {

	};
	Hello;
	World;
}
`
	if g.String() != expected {
		t.Fatalf("output is not expected")
	}

	if err := g.RemoveSubGraph("G", "cluster_FooBar"); err != nil {
		t.Fatal(err)
	}

	expected = `digraph G {
	Hello->World;
	Hello;
	World;
}
`
	if g.String() != expected {
		t.Fatalf("output is not expected")
	}
}

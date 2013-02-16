package gographviz

import (
	"code.google.com/p/gographviz/ast"
	"code.google.com/p/gographviz/parser"
	"testing"
)

type bugSubGraphWorldVisitor struct {
	g *Graph
	t *testing.T
}

func (this *bugSubGraphWorldVisitor) Visit(v ast.Elem) ast.Visitor {
	edge, ok := v.(*ast.EdgeStmt)
	if !ok {
		return this
	}
	if edge.Source.GetId() != "2" {
		return this
	}
	dst := edge.EdgeRHS[0].Destination
	if _, ok := dst.(*ast.SubGraph); !ok {
		this.t.Fatalf("2 -> Not SubGraph")
	}
	return this
}

func TestBugSubGraphWorld(t *testing.T) {
	g := analtest(t, "world.gv.txt")
	st, err := parser.ParseString(g.String())
	check(t, err)
	s := &bugSubGraphWorldVisitor{
		g: g,
		t: t,
	}
	st.Walk(s)
}

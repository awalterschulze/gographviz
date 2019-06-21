package gographviz

import (
	"testing"
)

func TestInvalidAttr(t *testing.T) {
	inputString := `
	digraph G {
    Ga->Gb;
    sA->sB [invalid];
	ssA->ssB;
	state [invalid2=ok]
}
`
	graphAst, err := ParseString(inputString)
	if err != nil {
		t.Fatal(err)
	}

	graph := NewGraph()

	if err := Analyse(graphAst, graph); err == nil {
		t.Fatal(err)
	}

	if err := AnalyseWithExtraAttrs(graphAst, graph, []string{"invalid", "invalid2"}); err != nil {
		t.Fatal(err)
	}

}

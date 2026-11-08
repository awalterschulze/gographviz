**This repo has moved to https://git.awalterschulze.org/gographviz**

**The go module has been renamed to `go.awalterschulze.org/gographviz`**

**This new module name and new repo location allowed us to fix the tags to have go compatible versioning. The previous newest version v2.0.3 was renamed to v0.2.3 at the new repo location. Now at the new location and with the new module path, a new version was tagged v0.3.0.**

**You can upgrade to the new version at the new location with following require statement in your module**

```
require go.awalterschulze.org/gographviz v0.3.0
```

**You will also need to update your imports to this new module path**

Parses the Graphviz DOT language and creates an interface, in golang, with which to easily create new and manipulate existing graphs which can be written back to the DOT format.

This parser has been created using [gocc](https://github.com/goccmack/gocc).

### Example (Parse and Edit) ###

```
package main

import "go.awalterschulze.org/gographviz"

func main() {
	graphAst, err := gographviz.ParseString(`digraph G {}`)
	if err != nil {
		panic(err)
	}
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}
	graph.AddNode("G", "a", nil)
	graph.AddNode("G", "b", nil)
	graph.AddEdge("a", "b", true, nil)
	output := graph.String()
	println(output)
}
```

### Documentation ###

  - The [godoc](https://godoc.org/go.awalterschulze.org/gographviz) includes some more examples.
  - [How to implement an anonymous subgraph](https://github.com/awalterschulze/gographviz/issues/59)

### Installation ###

```
go get go.awalterschulze.org/gographviz
```

### Users ###

  - [aptly](https://github.com/smira/aptly) - Debian repository management tool
  - [gorgonia](https://github.com/chewxy/gorgonia) - A Library that helps facilitate machine learning in Go
  - [imagemonkey](https://imagemonkey.io/graph?editor=true) - Let's create our own image dataset
  - [depviz](https://github.com/moul/depviz) - GitHub dependency visualizer (auto-roadmap)
  - [kustomize-graph](https://github.com/jpreese/kustomize-graph) - A tool to visualize Kustomize dependencies
  - [inframap](https://github.com/cycloidio/inframap) - Read your tfstate or HCL to generate a graph specific for each Terraform provider
  - [Antrea Traceflow](https://github.com/vmware-tanzu/antrea/blob/master/docs/traceflow-guide.md) supports using Traceflow for network diagnosis for Antrea, a Kubernetes networking solution intended to be Kubernetes native

### Mentions ###

[Using Golang and GraphViz to Visualize Complex Grails Applications](http://ilikeorangutans.github.io/2014/05/03/using-golang-and-graphviz-to-visualize-complex-grails-applications/)

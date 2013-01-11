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

package dot

import (
	"fmt"
)

type Interface interface {
	SetStrict(strict bool)
	SetDir(directed bool)
	SetName(name string)
	AddEdge(src, srcPort, dst, dstPort string, directed bool, attrs map[string]string)
	AddNode(parentGraph string, name string, attrs map[string]string)
	AddAttr(parentGraph string, field, value string)
	AddSubGraph(parentGraph string, name string, attrs map[string]string)
}

type Graph struct {
	Attrs     Attrs
	Name      string
	Directed  bool
	Strict    bool
	Nodes     *Nodes
	Edges     *Edges
	SubGraphs *SubGraphs
	Relations *Relations
}

func NewGraph() *Graph {
	return &Graph{
		Attrs:     make(Attrs),
		Name:      "",
		Directed:  false,
		Strict:    false,
		Nodes:     NewNodes(),
		Edges:     NewEdges(),
		SubGraphs: NewSubGraphs(),
		Relations: NewRelations(),
	}
}

func (this *Graph) SetStrict(strict bool) {
	this.Strict = strict
}

func (this *Graph) SetDir(dir bool) {
	this.Directed = dir
}

func (this *Graph) SetName(name string) {
	this.Name = name
}

func (this *Graph) AddEdge(src, srcPort, dst, dstPort string, directed bool, attrs map[string]string) {
	this.Edges.Add(&Edge{src, srcPort, dst, dstPort, directed, attrs})
}

func (this *Graph) AddNode(parentGraph string, name string, attrs map[string]string) {
	this.Nodes.Add(&Node{name, attrs})
	this.Relations.Add(parentGraph, name)
}

func (this *Graph) getAttrs(graphName string) Attrs {
	if this.Name == graphName {
		return this.Attrs
	}
	g, ok := this.SubGraphs.SubGraphs[graphName]
	if !ok {
		panic("graph or subgraph " + graphName + " does not exist")
	}
	return g.Attrs
}

func (this *Graph) AddAttr(parentGraph string, field string, value string) {
	this.getAttrs(parentGraph).Add(field, value)
}

func (this *Graph) AddSubGraph(parentGraph string, name string, attrs map[string]string) {
	fmt.Printf("AddSubGraph %#v\n", name)
	this.SubGraphs.Add(name)
	for key, value := range attrs {
		this.AddAttr(name, key, value)
	}
}

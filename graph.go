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
	"github.com/awalterschulze/gographviz/common"
)

//The analysed representation of the Graph parsed from the DOT format.
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

//Creates a new empty graph, ready to be populated.
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

//If the graph is strict then multiple edges are not allowed between the same pairs of nodes,
//see dot man page.
func (this *Graph) SetStrict(strict bool) {
	this.Strict = strict
}

//Sets whether the graph is directed (true) or undirected (false).
func (this *Graph) SetDir(dir bool) {
	this.Directed = dir
}

//Sets the graph name.
func (this *Graph) SetName(name string) {
	this.Name = name
}

//Adds an edge to the graph from node src to node dst.
//srcPort and dstPort are the port the node ports, leave as empty strings if it is not required.
//This does not imply the adding of missing nodes.
func (this *Graph) AddPortEdge(src, srcPort, dst, dstPort string, directed bool, inputAttrs map[string]string) {
	attrs, err := fromStringMap(inputAttrs)
	if err != nil {
		panic(err) // TODO : proper handling?
	}
	this.addPortEdge(src, srcPort, dst, dstPort, directed, attrs)
}

func (this *Graph) addPortEdge(src, srcPort, dst, dstPort string, directed bool, attrs Attrs) {
	this.Edges.Add(&Edge{src, srcPort, dst, dstPort, directed, attrs})
}

//Adds an edge to the graph from node src to node dst.
//This does not imply the adding of missing nodes.
func (this *Graph) AddEdge(src, dst string, directed bool, inputAttrs map[string]string) {
	attrs, err := fromStringMap(inputAttrs)
	if err != nil {
		panic(err) // TODO : proper handling?
	}
	this.addPortEdge(src, "", dst, "", directed, attrs)
}

//Adds a node to a graph/subgraph.
//If not subgraph exists use the name of the main graph.
//This does not imply the adding of a missing subgraph.
func (this *Graph) AddNode(parentGraph, name string, inputAttrs map[string]string) {
	attrs, err := fromStringMap(inputAttrs)
	if err != nil {
		panic(err) // TODO : proper handling?
	}
	this.addNode(parentGraph, name, attrs)
}

func (this *Graph) addNode(parentGraph, name string, attrs Attrs) {
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

//Adds an attribute to a graph/subgraph.
func (this *Graph) AddAttr(parentGraph string, field, value string) {
	this.getAttrs(parentGraph).Add(field, value)
}

func (this *Graph) addAttr(parentGraph string, field common.Attribute, value string) {
	this.getAttrs(parentGraph).add(field, value)
}

//Adds a subgraph to a graph/subgraph.
func (this *Graph) AddSubGraph(parentGraph, name string, inputAttrs map[string]string) {
	attrs, err := fromStringMap(inputAttrs)
	if err != nil {
		panic(err) // TODO : proper handling?
	}
	this.addSubGraph(parentGraph, name, attrs)
}

func (this *Graph) addSubGraph(parentGraph, name string, attrs Attrs) {
	this.SubGraphs.Add(name)
	for key, value := range attrs {
		this.addAttr(name, key, value)
	}
}

func (this *Graph) IsNode(name string) bool {
	_, ok := this.Nodes.Lookup[name]
	return ok
}

func (this *Graph) IsSubGraph(name string) bool {
	_, ok := this.SubGraphs.SubGraphs[name]
	return ok
}

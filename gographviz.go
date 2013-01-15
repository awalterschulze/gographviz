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

//Package gographviz provides parsing for the DOT grammar into
//an abstract syntax tree representing a graph,
//analysis of the abstract syntax tree into a more usable structure, 
//and writing back of this structure into the DOT format.
package gographviz

import (
	"gographviz/parser"
	"gographviz/ast"
)

var _ Interface = NewGraph()

//Implementing this interface allows you to parse the graph into your own structure.
type Interface interface {
	//If the graph is strict then multiple edges are not allowed between the same pairs of nodes, 
	//see dot man page.
	SetStrict(strict bool)
	//Sets whether the graph is directed (true) or undirected (false).
	SetDir(directed bool)
	//Sets the graph name.
	SetName(name string)
	//Adds an edge to the graph from node src to node dst.
	//srcPort and dstPort are the port the node ports, leave as empty strings if it is not required.
	//This does not imply the adding of missing nodes.
	AddEdge(src, srcPort, dst, dstPort string, directed bool, attrs map[string]string)
	//Adds a node to a graph/subgraph.
	//If not subgraph exists use the name of the main graph.
	//This does not imply the adding of a missing subgraph.
	AddNode(parentGraph string, name string, attrs map[string]string)
	//Adds an attribute to a graph/subgraph.
	AddAttr(parentGraph string, field, value string)
	//Adds a subgraph to a graph/subgraph.
	AddSubGraph(parentGraph string, name string, attrs map[string]string)
}

//Parses the buffer into a abstract syntax tree representing the graph.
func Parse(buf []byte) (*ast.Graph, error) {
	return parser.ParseBytes(buf)
}

//Parses and creates a new Graph from the data.
func Read(buf []byte) (*Graph, error) {
	st, err := Parse(buf)
	if err != nil {
		return nil, err
	}
	return NewAnalysedGraph(st), nil
}

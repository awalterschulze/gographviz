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
	"fmt"

	"github.com/awalterschulze/gographviz/ast"
)

type writer struct {
	*Graph
	writtenLocations map[string]bool
}

func newWriter(g *Graph) *writer {
	return &writer{g, make(map[string]bool)}
}

func appendAttrs(list ast.StmtList, attrs Attrs) ast.StmtList {
	for _, name := range attrs.SortedNames() {
		stmt := &ast.Attr{
			Field: ast.Id(name),
			Value: ast.Id(attrs[name]),
		}
		list = append(list, stmt)
	}
	return list
}

func (this *writer) newSubGraph(name string) (*ast.SubGraph, error) {
	sub := this.SubGraphs.SubGraphs[name]
	this.writtenLocations[sub.Name] = true
	s := &ast.SubGraph{}
	s.Id = ast.Id(sub.Name)
	s.StmtList = appendAttrs(s.StmtList, sub.Attrs)
	children := this.Relations.SortedChildren(name)
	for _, child := range children {
		if this.IsNode(child) {
			s.StmtList = append(s.StmtList, this.newNodeStmt(child))
		} else if this.IsSubGraph(child) {
			subgraph, err := this.newSubGraph(child)
			if err != nil {
				return nil, err
			}
			s.StmtList = append(s.StmtList, subgraph)
		} else {
			return nil, fmt.Errorf("%v is not a node or a subgraph", child)
		}
	}
	return s, nil
}

func (this *writer) newNodeId(name string, port string) *ast.NodeId {
	node := this.Nodes.Lookup[name]
	return ast.MakeNodeId(node.Name, port)
}

func (this *writer) newNodeStmt(name string) *ast.NodeStmt {
	node := this.Nodes.Lookup[name]
	id := ast.MakeNodeId(node.Name, "")
	this.writtenLocations[node.Name] = true
	return &ast.NodeStmt{
		id,
		ast.PutMap(node.Attrs.ToMap()),
	}
}

func (this *writer) newLocation(name string, port string) (ast.Location, error) {
	if this.IsNode(name) {
		return this.newNodeId(name, port), nil
	} else if this.IsClusterSubGraph(name) {
		if len(port) != 0 {
			return nil, fmt.Errorf("subgraph cannot have a port: %v", port)
		}
		return ast.MakeNodeId(name, port), nil
	} else if this.IsSubGraph(name) {
		if len(port) != 0 {
			return nil, fmt.Errorf("subgraph cannot have a port: %v", port)
		}
		return this.newSubGraph(name)
	}
	return nil, fmt.Errorf("%v is not a node or a subgraph", name)
}

func (this *writer) newEdgeStmt(edge *Edge) (*ast.EdgeStmt, error) {
	src, err := this.newLocation(edge.Src, edge.SrcPort)
	if err != nil {
		return nil, err
	}
	dst, err := this.newLocation(edge.Dst, edge.DstPort)
	if err != nil {
		return nil, err
	}
	stmt := &ast.EdgeStmt{
		Source: src,
		EdgeRHS: ast.EdgeRHS{
			&ast.EdgeRH{
				ast.EdgeOp(edge.Dir),
				dst,
			},
		},
		Attrs: ast.PutMap(edge.Attrs.ToMap()),
	}
	return stmt, nil
}

func (this *writer) Write() (*ast.Graph, error) {
	t := &ast.Graph{}
	t.Strict = this.Strict
	t.Type = ast.GraphType(this.Directed)
	t.Id = ast.Id(this.Name)

	t.StmtList = appendAttrs(t.StmtList, this.Attrs)

	for _, edge := range this.Edges.Edges {
		e, err := this.newEdgeStmt(edge)
		if err != nil {
			return nil, err
		}
		t.StmtList = append(t.StmtList, e)
	}

	subGraphs := this.SubGraphs.Sorted()
	for _, s := range subGraphs {
		if _, ok := this.writtenLocations[s.Name]; !ok {
			if _, ok := this.Relations.ParentToChildren[this.Name][s.Name]; ok {
				s, err := this.newSubGraph(s.Name)
				if err != nil {
					return nil, err
				}
				t.StmtList = append(t.StmtList, s)
			}
		}
	}

	nodes := this.Nodes.Sorted()
	for _, n := range nodes {
		if _, ok := this.writtenLocations[n.Name]; !ok {
			t.StmtList = append(t.StmtList, this.newNodeStmt(n.Name))
		}
	}

	return t, nil
}

//Creates an Abstract Syntrax Tree from the Graph.
func (g *Graph) WriteAst() (*ast.Graph, error) {
	w := newWriter(g)
	return w.Write()
}

//Returns a DOT string representing the Graph.
func (g *Graph) String() string {
	w, err := g.WriteAst()
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return w.String()
}

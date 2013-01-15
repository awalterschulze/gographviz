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
	"code.google.com/p/gographviz/ast"
	"sort"
)

func sortedSubGraphs(keyvalue map[string]*SubGraph) []string {
	keys := make([]string, 0)
	for key, _ := range keyvalue {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func sortedKeys3(keyvalue map[string]bool) []string {
	keys := make([]string, 0)
	for key, _ := range keyvalue {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func sortedKeys(keyvalue map[string]string) []string {
	keys := make([]string, 0)
	for key, _ := range keyvalue {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

//Creates an Abstract Syntrax Tree from the Graph.
func (g *Graph) Write() *ast.Graph {
	nodes := make(map[string][]*ast.NodeStmt)
	addedNodes := make(map[string]bool)
	for parent, children := range g.Relations.ParentToChildren {
		if parent == g.Name {
			continue
		}
		nodes[parent] = make([]*ast.NodeStmt, 0)
		sortedChildren := sortedKeys3(children)
		for _, child := range sortedChildren {
			node := g.Nodes.Lookup[child]
			stmt := &ast.NodeStmt{
				ast.MakeNodeId(node.Name, ""),
				ast.PutMap(node.Attrs),
			}
			addedNodes[child] = true
			nodes[parent] = append(nodes[parent], stmt)
		}
	}
	t := &ast.Graph{}
	t.Strict = g.Strict
	t.Type = ast.GraphType(g.Directed)
	t.Id = ast.Id(g.Name)
	attrs := g.Attrs
	keys := sortedKeys(attrs)
	for _, key := range keys {
		value := attrs[key]
		stmt := &ast.Attr{
			Field: ast.Id(key),
			Value: ast.Id(value),
		}
		t.StmtList = append(t.StmtList, stmt)
	}
	sortedChildren := sortedKeys3(g.Relations.ParentToChildren[g.Name])
	for _, child := range sortedChildren {
		if _, ok := addedNodes[child]; ok {
			continue
		}
		node := g.Nodes.Lookup[child]
		stmt := &ast.NodeStmt{
			ast.MakeNodeId(node.Name, ""),
			ast.PutMap(node.Attrs),
		}
		t.StmtList = append(t.StmtList, stmt)
	}
	names := sortedSubGraphs(g.SubGraphs.SubGraphs)
	for _, name := range names {
		sub := g.SubGraphs.SubGraphs[name]
		subGraph := &ast.SubGraph{
			Id: ast.Id(name),
		}
		attrs := sub.Attrs
		keys := sortedKeys(attrs)
		for _, key := range keys {
			value := attrs[key]
			stmt := &ast.Attr{
				Field: ast.Id(key),
				Value: ast.Id(value),
			}
			subGraph.StmtList = append(subGraph.StmtList, stmt)
		}
		ns, ok := nodes[name]
		if !ok {
			continue
		}
		for _, n := range ns {
			subGraph.StmtList = append(subGraph.StmtList, n)
		}
		t.StmtList = append(t.StmtList, subGraph)
	}
	for _, edge := range g.Edges.Edges {
		stmt := &ast.EdgeStmt{
			Source: ast.MakeNodeId(edge.Src, edge.SrcPort),
			EdgeRHS: ast.EdgeRHS{
				&ast.EdgeRH{
					ast.EdgeOp(edge.Dir),
					ast.MakeNodeId(edge.Dst, edge.DstPort),
				},
			},
			Attrs: ast.PutMap(edge.Attrs),
		}
		t.StmtList = append(t.StmtList, stmt)
	}
	return t
}

//Returns a DOT string representing the Graph.
func (g *Graph) String() string {
	return g.Write().String()
}

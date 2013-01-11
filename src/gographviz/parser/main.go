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

package parser

import (
	"gographviz/scanner"
	"gographviz/token"
	"gographviz/ast"
	"io"
	"io/ioutil"
	"os"
	"fmt"
)

func ParseString(dotString string) (*ast.Graph, error) {
	return ParseBytes([]byte(dotString))
}

func ParseBytes(dotBytes []byte) (*ast.Graph, error) {
	lex := &scanner.Scanner{}
	lex.Init(dotBytes, token.DOTTokens)
	parser := NewParser(ActionTable, GotoTable, ProductionsTable, token.DOTTokens)
	st, err := parser.Parse(lex)
	if err != nil {
		return nil, err
	}
	g, ok := st.(*ast.Graph)
	if !ok {
		panic(fmt.Sprintf("Parser did not return an *ast.Graph, but rather a %T", st))
	}
	return g, err
}

func Parse(r io.Reader) (*ast.Graph, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParseBytes(bytes)
}

func ParseFile(filename string) (*ast.Graph, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return Parse(f)
}

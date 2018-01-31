package ast

import (
	"testing"
)

func TestMakeNodeID(t *testing.T) {
	tcs := []struct {
		id   string
		port string
		want NodeID
	}{
		{ // TC#1
			id:   "",
			port: "",
			want: NodeID{}},
		{ // TC#2
			id:   "id",
			port: "p1",
			want: NodeID{"id", Port{"p1", ""}},
		},
		{ // TC#3
			id:   "_id",
			port: "p1:p2",
			want: NodeID{"_id", Port{"p1", "p2"}},
		},
		{ // TC#4
			id:   "1id",
			port: "p1:p2:p3",
			want: NodeID{"1id", Port{"p1", "p2"}},
		},
		{ // TC#5
			id:   "?id",
			port: ":p2:p3",
			want: NodeID{"?id", Port{"", "p2"}},
		},
	}

	for i, tc := range tcs {
		n := MakeNodeID(tc.id, tc.port)

		if *n != tc.want {
			t.Fatalf("TC#%d: MakeNodeId(%q, %q)=%#v, want %#v", i+1, tc.id, tc.port, *n, tc.want)
		}
	}
}

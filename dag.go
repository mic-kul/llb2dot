package llb2dot

import (
	"io"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

var increment int64

var addednodes = map[string]int64{}

type node struct {
	id         int64
	nodeDigest string
	desc       string
}

func (n node) ID() int64 {
	return n.id
}

func (n node) DOTID() string {
	return n.desc
}

func (n node) Attributes() []encoding.Attribute {
	return []encoding.Attribute{{Key: "digest", Value: n.nodeDigest}}
}

func newNode(g *simple.DirectedGraph, nodeDigest, desc string) node {
	increment++
	addednodes[nodeDigest] = increment
	return node{id: increment, nodeDigest: nodeDigest, desc: desc}
}

// WriteDOT output graph to dot language.
func WriteDOT(w io.Writer, g graph.Graph) error {
	b, err := dot.Marshal(g, "llb", "", "")
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
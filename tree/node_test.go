package tree

import (
	"testing"
)

func TestNodeObjectCreatedWithID(t *testing.T) {
	ID := "a1"
	n := CreateNode(ID)
	if n.ID != ID {
		t.Errorf("CreateNode creates node, ID=%s, but expected was ID=%s", n.ID, ID)
	}
	stringRepr := "<Node id=" + ID + ", data= >"
	if n.String() != stringRepr {
		t.Errorf("Created node representation = %s, but expected was = %s", n.String(),
			stringRepr)
	}
}

func TestNodeObjectCreatedParams(t *testing.T) {
	ID := "a1"
	data := "data"
	n := CreateNodeFull(NodeParams{ID: ID, Data: data})
	if n.ID != ID {
		t.Errorf("CreateNodeFull creates node, ID=%s, but expected was ID=%s", n.ID, ID)
	}
	if n.Data != data {
		t.Errorf("CreateNodeFull creates node, data=%s, but expected was data=%s", n.Data, data)
	}
	stringRepr := "<Node id=" + ID + ", data=" + data + " >"
	if n.String() != stringRepr {
		t.Errorf("Created node representation = %s, but expected was = %s", n.String(),
			stringRepr)
	}
}

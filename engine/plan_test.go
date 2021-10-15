package engine

import (
	"strings"
	"testing"

	"github.com/chytilp/sudoku/tree"
	"github.com/google/go-cmp/cmp"
)

/*
func TestPlanFindSolutions(t *testing.T) {
	plan := NewPlan()
	g, err := structures.NewGameFromString(game1)
	if err != nil {
		t.Errorf("Game should be succesfully created, but err: %v", err)
	}
	ok, err := plan.FindSolutions(g)
	if err != nil {
		t.Errorf("FindSolutions should not return error: %v", err)
	}
	if !*ok {
		t.Errorf("FindSolutions should return ok.")
	}
	solutions := make(map[string]bool)
	for _, n := range plan.solutionTree.RootNodes() {
		solutions[n.ID] = true
		fmt.Printf("%s\n", n.ID)
	}
}*/

func createPlan() *Plan {
	p := NewPlan()
	p.AddNodes("a", []string{"b", "c", "d"})
	p.AddNodes("a1", []string{"a2", "a3", "a4"})
	p.AddNodes("a11", []string{"a12", "a13", "a14"})
	p.AddNodes("a111", []string{"a112", "a113", "a114"})
	p.SetCurrent(p.FindNode("b"))
	p.AddNodes("b1", []string{"b2", "b3"})
	p.AddNodes("b11", []string{"b12", "b13"})
	p.SetCurrent(p.FindNode("c"))
	p.AddNodes("c1", []string{"c2"})
	return p
}

func TestPlanAddNodes(t *testing.T) {
	p := createPlan()
	schema := p.solutionTree.Display(" - ")
	expected := "a - a1 - a11 - a111\n" +
		"a - a1 - a11 - a112\n" +
		"a - a1 - a11 - a113\n" +
		"a - a1 - a11 - a114\n" +
		"a - a1 - a12\n" +
		"a - a1 - a13\n" +
		"a - a1 - a14\n" +
		"a - a2\n" +
		"a - a3\n" +
		"a - a4\n" +
		"b - b1 - b11\n" +
		"b - b1 - b12\n" +
		"b - b1 - b13\n" +
		"b - b2\n" +
		"b - b3\n" +
		"c - c1\n" +
		"c - c2\n" +
		"d"
	if diff := cmp.Diff(schema, expected); diff != "" {
		t.Errorf("Plan display returns: %s, but expected: %s, diff: %s\n", schema, expected, diff)
	}
}

func TestPlanSetNodesDone(t *testing.T) {
	p := createPlan()
	ids := make([]string, 0, 25)
	ids = getDoneNodeIds(p.solutionTree.RootNodes(), ids)
	idecka := strings.Join(ids, ",")
	expected := "a,a1,a11,a111,b,b1,b11,c,c1"
	if diff := cmp.Diff(idecka, expected); diff != "" {
		t.Errorf("Plan done nodes ids: %s, but expected: %s, diff: %s\n", idecka, expected, diff)
	}
	p.SetNodesDone([]string{"a", "a1", "a11", "a111"})
	ids = make([]string, 0, 25)
	ids = getDoneNodeIds(p.solutionTree.RootNodes(), ids)
	idecka = strings.Join(ids, ",")
	expected = "a,a1,a11,a111"
	if diff := cmp.Diff(idecka, expected); diff != "" {
		t.Errorf("Plan done nodes ids: %s, but expected: %s, diff: %s\n", idecka, expected, diff)
	}
}

func TestPlanFindNearest(t *testing.T) {
	p := createPlan()
	p.SetNodesDone([]string{"a", "a1", "a11", "a111"})
	p.SetCurrent(p.FindNode("a111"))
	nearest := []string{"a112", "a113", "a114", "a12", "a13", "a14", "a2", "a3", "a4", "b", "c", "d"}
	var n *tree.Node
	for _, next := range nearest {
		n = p.FindNearest()
		if n.ID != next {
			t.Errorf("Plan FindNearest returns: %s, but expected: %s", n.ID, next)
		}
		p.SetCurrent(n)
	}
	//nothing more nearest
	n = p.FindNearest()
	if n != nil {
		t.Errorf("Plan FindNearest should not returns next node, but returned: %s\n", n.ID)
	}
}

func getDoneNodeIds(nodes []*tree.Node, ids []string) []string {
	if nodes == nil {
		return ids
	}
	for _, n := range nodes {
		if n.Data == doneTrue {
			ids = append(ids, n.ID)
		}
		ids = getDoneNodeIds(n.Children, ids)
	}
	return ids
}

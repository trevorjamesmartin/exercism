package tree

import (
	"fmt"
	"sort" // https://pkg.go.dev/sort
)

type Record struct {
	ID     int
	Parent int
}

// By is the type of a "less" function that defines the ordering of its Record arguments.
type By func(r1, r2 *Record) bool

// joins a By function and a slice of Records to be sorted
type RecordSorter struct {
	records []Record
	by      By
}

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(records []Record) {
	sort.Sort(&RecordSorter{records: records, by: by})
}

// Len is part of sort.Interface
func (r *RecordSorter) Len() int {
	return len(r.records)
}

// Swap is part of sort.Interface
func (r *RecordSorter) Swap(a, b int) {
	r.records[a], r.records[b] = r.records[b], r.records[a]
}

// Less is part of sort.Interface
func (r *RecordSorter) Less(a, b int) bool {
	return r.by(&r.records[a], &r.records[b])
}

type Node struct {
	ID       int
	Children []*Node
}

func (node *Node) Find(id int) *Node {
	if id == node.ID {
		return node
	}
	for _, child := range node.Children {
		c := child.Find(id)
		if c != nil {
			return c
		}
	}
	return nil
}

type RecordTree struct {
	Root  *Node
	Count int
}

func (tree *RecordTree) Insert(r Record) (bool, error) {
	switch {
	case tree.Count == 0 && r.ID > 0:
		return false, fmt.Errorf("error inserting record id: %v (no root)", r.ID)
	case tree.Count == 0 && r.Parent > 0:
		return false, fmt.Errorf("error inserting record id: %v (root node has parent)", r.ID)
	case tree.Count == 0 && r.ID == 0:
		// Create root
		tree.Root = &Node{ID: r.ID}
		tree.Count++
	case tree.Root != nil:
		parent := tree.Root.Find(r.Parent)
		if parent == nil {
			return false, fmt.Errorf("error finding parent (%v) of record: %v", r.Parent, r.ID)
		}
		if parent.Find(r.ID) != nil {
			return false, fmt.Errorf("error, duplicate (%v) found", r.ID)
		}
		// Append child
		parent.Children = append(parent.Children, &Node{ID: r.ID})
		tree.Count++
	}
	return true, nil
}

func Build(records []Record) (*Node, error) {
	var err error
	var tree = &RecordTree{}

	record_ID := func(r1, r2 *Record) bool {
		return r1.ID < r2.ID
	}

	By(record_ID).Sort(records)

	for i, record := range records {
		if i != record.ID {
			return nil, fmt.Errorf("error, non-continuous (%v != %v)", record.ID, i)
		}
		ok, err := tree.Insert(record)
		if !ok {
			return nil, err
		}
	}

	return tree.Root, err
}

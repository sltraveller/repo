package main

import (
	"fmt"
	"sort"
	"web/3w/nodes"
)

type Apnode struct {
	Value    string
	Label    string
	Children []*Apnode
}

func main() {
	node := nodes.GetNodes()
	show := []*Apnode{}
	for name, n := range node.Children {
		apps := &Apnode{
			Value:    name,
			Label:    name,
			Children: []*Apnode{},
		}
		for aname := range n.Children {
			child := &Apnode{
				Value: aname,
				Label: aname,
			}
			apps.Children = append(apps.Children, child)
		}
		show = append(show, apps)
	}
}
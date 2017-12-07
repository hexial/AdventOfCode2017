package day07

import (
	"AdventOfCode2017/util"
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Tree struct {
	Nodes []*TreeNode
}

type TreeNode struct {
	Name          string
	Weight        int
	Children      []*TreeNode
	ChildrenNames []string
	Parent        *TreeNode
}

func (this Tree) FindBase() *TreeNode {
	for _, node := range this.Nodes {
		if len(node.Children) > 0 && node.Parent == nil {
			return node
		}
	}
	return nil
}

func (this Tree) Balance() int {
	return this.FindBase().Balance()
}

func (this Tree) BalanceNew() int {
	base := this.FindBase()
	log.Infof("base=%s", base.Name)
	for _, child := range base.Children {
		log.Infof("child=%s (%d) : %d", child.Name, child.Weight, child.WeightWithChildren())
	}
	return 0
}

func (this *TreeNode) DebugInfo() string {
	return fmt.Sprintf("%s (%d) (%d)", this.Name, this.Weight, this.WeightSum())
}

func (this *TreeNode) Balance() int {
	//log.Infof("Balance : %s", this.Name)
	var incorrect *TreeNode
	var prev *TreeNode
	var wanted int
	if len(this.Children) > 1 {
		prev = this.Children[0]
		for i := 1; i < len(this.Children); i++ {
			curr := this.Children[i]
			if incorrect == nil && prev.WeightSum() != curr.WeightSum() {
				incorrect = prev
				wanted = curr.WeightSum()
			} else if incorrect != nil {
				if curr.WeightSum() != prev.WeightSum() {
					if curr.WeightSum() != incorrect.WeightSum() {
						incorrect = curr
						wanted = prev.WeightSum()
					} else {
						incorrect = prev
						wanted = curr.WeightSum()
					}
				}
			}
			prev = curr
			if incorrect != nil {
				//log.Infof("Balance : %d : prev=%s : curr=%s : incorrect=%s", i, prev.DebugInfo(), curr.DebugInfo(), incorrect.DebugInfo())
			} else {
				//log.Infof("Balance : %d : prev=%s : curr=%s", i, prev.DebugInfo(), curr.DebugInfo())
			}
		}
	}
	if incorrect != nil {
		log.Infof("************************")
		for _, child := range this.Children {
			log.Infof("%s : %d", child.Name, child.WeightSum())
		}
		newWeight := wanted - incorrect.WeightOfChildren()
		log.Infof("%s : wanted=%d : new=%d", incorrect.Name, wanted, newWeight)
		return newWeight
	}
	//
	// Balance children
	for _, child := range this.Children {
		n := child.Balance()
		if n > 0 {
			return n
		}
	}
	return 0
}

func (this *TreeNode) WeightWithChildren() int {
	w := this.Weight
	for _, child := range this.Children {
		w += child.WeightWithChildren()
	}
	return w
}

func (this *TreeNode) WeightOfChildren() int {
	sum := 0
	for _, child := range this.Children {
		sum += child.Weight
	}
	return sum
}

func (this *TreeNode) WeightSum() int {
	return this.Weight + this.WeightOfChildren()
}

func Load(filename string) Tree {
	var err error
	tree := Tree{}

	//
	// Parse file
	rows := util.FileAsLineArray(filename)
	for _, row := range rows {
		rowItems := strings.Split(row, " ")
		node := new(TreeNode)
		node.Name = rowItems[0]
		node.Weight, err = strconv.Atoi(rowItems[1][1 : len(rowItems[1])-1])
		if err != nil {
			log.Panic(err)
		}
		for i := 3; i < len(rowItems); i++ {
			if strings.HasSuffix(rowItems[i], ",") {
				node.ChildrenNames = append(node.ChildrenNames, rowItems[i][:len(rowItems[i])-1])
			} else {
				node.ChildrenNames = append(node.ChildrenNames, rowItems[i])
			}
		}
		tree.Nodes = append(tree.Nodes, node)
	}

	//
	// Build tree
	for _, node := range tree.Nodes {
		for _, childName := range node.ChildrenNames {
			for _, n := range tree.Nodes {
				if n.Name == childName {
					if n.Parent != nil {
						log.Panicf("Didn't except more than one parent")
					}
					node.Children = append(node.Children, n)
					n.Parent = node
				}
			}
		}
	}
	return tree
}

package main

import (
	"errors"
	"fmt"
)

type DNode struct {
	Id           int
	Name         string
	NextNode     *DNode
	PreviousNode *DNode
}

// InsertDNode insert to the last position
func InsertDNode(head *DNode, newNode *DNode) {
	tmp := head
	for {
		if tmp.NextNode == nil {
			break
		}
		tmp = tmp.NextNode
	}

	tmp.NextNode = newNode
	newNode.PreviousNode = tmp
}

func InsertDNodeOrderly(head *DNode, newNode *DNode) error {
	tmp := head
	for {
		if tmp.NextNode == nil {
			break
		} else if tmp.NextNode.Id > newNode.Id {
			// nice position
			break
		} else if tmp.NextNode.Id == newNode.Id {
			return errors.New("cannot insert newNode cause the same id")
		}
		tmp = tmp.NextNode
	}

	newNode.NextNode = tmp.NextNode
	newNode.PreviousNode = tmp
	tmp.NextNode = newNode
	if tmp.NextNode != nil {
		tmp.NextNode.PreviousNode = newNode
	}
	return nil
}

func DeleteDNode(head *DNode, id int) error {
	tmp := head
	for {
		if tmp.NextNode == nil {
			return errors.New("there is no such a node with indicated id, nothing happened")
		} else if tmp.NextNode.Id == id {
			tmp.NextNode = tmp.NextNode.NextNode
			if tmp.NextNode != nil {
				// tmp.NextNode here is equivalent with tmp.NextNode.NextNode previously
				tmp.NextNode.PreviousNode = tmp
			}
			return nil
		}
		tmp = tmp.NextNode
	}
}

func ShowAllDNodeOrder(head *DNode) {
	tmp := head
	if tmp.NextNode == nil {
		fmt.Println("There is nothing")
		return
	}

	for {
		fmt.Printf("[id: %d, name: %s] ==> ", tmp.NextNode.Id, tmp.NextNode.Name)
		tmp = tmp.NextNode
		if tmp.NextNode == nil {
			break
		}

	}
}

func ShowAllDNodeReverseOrder(head *DNode) {
	tmp := head
	if tmp.NextNode == nil {
		fmt.Println("There is nothing")
		return
	}

	for {
		if tmp.NextNode == nil {
			break
		}
		tmp = tmp.NextNode
	}

	for {
		fmt.Printf("[id: %d, name: %s] ==> ", tmp.Id, tmp.Name)
		tmp = tmp.PreviousNode

		if tmp.PreviousNode == nil {
			break
		}
	}
}

func main() {
	head := &DNode{}

	node1 := &DNode{
		Id:   1,
		Name: "rey",
	}

	node2 := &DNode{
		Id:   2,
		Name: "charlotte",
	}

	node3 := &DNode{
		Id:   3,
		Name: "martin",
	}
	node4 := &DNode{
		Id:   4,
		Name: "dora",
	}

	node5 := &DNode{
		Id:   4,
		Name: "lucy",
	}
	InsertDNode(head, node1)
	InsertDNode(head, node2)
	fmt.Println("======order===========")
	ShowAllDNodeOrder(head)
	fmt.Println()
	fmt.Println("===========reverse order===============")
	ShowAllDNodeReverseOrder(head)
	fmt.Println()

	InsertDNodeOrderly(head, node4)
	ShowAllDNodeOrder(head)
	fmt.Println()

	InsertDNodeOrderly(head, node3)
	ShowAllDNodeOrder(head)
	fmt.Println()

	err := InsertDNodeOrderly(head, node5)
	if err != nil {
		fmt.Println(err.Error())
	}
	ShowAllDNodeOrder(head)
	fmt.Println()

	err = DeleteDNode(head, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	ShowAllDNodeOrder(head)
}

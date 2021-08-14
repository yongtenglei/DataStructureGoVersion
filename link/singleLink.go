package main

import (
	"errors"
	"fmt"
)

type SNode struct {
	Id       int
	Name     string
	NextNode *SNode
}

// InsertSNode insert to the last position
func InsertSNode(head *SNode, newNode *SNode) {
	tmp := head
	for {
		if tmp.NextNode == nil {
			break
		}
		tmp = tmp.NextNode
	}

	tmp.NextNode = newNode
}

func InsertSNodeOrderly(head *SNode, newNode *SNode) error {
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
	tmp.NextNode = newNode
	return nil
}

func DeleteSNode(head *SNode, id int) error {
	tmp := head
	for {
		if tmp.NextNode == nil {
			return errors.New("there is no such a node with indicated id, nothing happened")
		} else if tmp.NextNode.Id == id {
			tmp.NextNode = tmp.NextNode.NextNode
			return nil
		}
		tmp = tmp.NextNode
	}
}

func ShowAllSNode(head *SNode) {
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

func main() {
	head := &SNode{}

	node1 := &SNode{
		Id:   1,
		Name: "rey",
	}

	node2 := &SNode{
		Id:   2,
		Name: "charlotte",
	}

	node3 := &SNode{
		Id:   3,
		Name: "martin",
	}
	node4 := &SNode{
		Id:   4,
		Name: "dora",
	}

	node5 := &SNode{
		Id:   4,
		Name: "lucy",
	}
	InsertSNodeOrderly(head, node1)
	ShowAllSNode(head)
	fmt.Println()
	InsertSNodeOrderly(head, node4)
	ShowAllSNode(head)
	fmt.Println()
	InsertSNodeOrderly(head, node3)
	ShowAllSNode(head)
	fmt.Println()
	InsertSNodeOrderly(head, node2)
	ShowAllSNode(head)
	fmt.Println()

	err := InsertSNodeOrderly(head, node5)
	if err != nil {
		fmt.Println(err.Error())
	}

	ShowAllSNode(head)
	fmt.Println()
	err = DeleteSNode(head, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	ShowAllSNode(head)
}

package main

import (
	"errors"
	"fmt"
)

type CNode struct {
	Id       int
	Name     string
	NextNode *CNode
}

func InsertCNode(head *CNode, newNode *CNode) {
	if head.NextNode == nil {
		head.Id = newNode.Id
		head.Name = newNode.Name
		head.NextNode = head
		return
	}

	tmp := head
	for {
		if tmp.NextNode == head {
			break
		}
		tmp = tmp.NextNode
	}

	tmp.NextNode = newNode
	newNode.NextNode = head
}

// DeleteCNode 需要返回头结点，因为删除头结点后，头结点已经改变。
func DeleteCNode(head *CNode, id int) (*CNode, error) {
	tmp := head    //辅助节点指向链表头节点head
	helper := head //辅助节点指向链表最后一个节点

	//判断链表为空
	if head.NextNode == nil {
		return head, errors.New("there is nothing")
	}

	//判断只有一个节点，且这个节点的no编号等于id时删除
	if head.NextNode == head && head.Id == id {
		head.NextNode = nil
		return head, nil
	} else if head.NextNode == head && head.Id != id {
		return head, errors.New("only one node, but have another id")
	}

	// 有多个节点时， 将helper指向链表最后一个节点
	for {
		if helper.NextNode == head {
			break
		}
		helper = helper.NextNode
	}

	for {
		// 此时tmp是最后一个节点，但没有比较是否为需要删除的节点
		if tmp.NextNode == head {
			break
		}
		if tmp.Id == id {
			// 多个节点，且要删除的节点是第一个节点， 将head后移
			if tmp == head {
				head = head.NextNode
			}
			helper.NextNode = tmp.NextNode
			return head, nil
		}
		tmp = tmp.NextNode
		helper = helper.NextNode
	}

	if tmp.Id == id {
		helper.NextNode = tmp.NextNode
		return head, nil
	} else {
		return head, errors.New("no such a node to delete")
	}
}

func ShowAllCNode(head *CNode) {
	tmp := head

	if tmp.NextNode == nil {
		fmt.Println("There is nothing")
		return
	}

	for {
		fmt.Printf("[id: %d, name: %s] ==> ", tmp.Id, tmp.Name)

		if tmp.NextNode == head {
			break
		}
		tmp = tmp.NextNode
	}
}

func main() {
	head := &CNode{}

	node1 := &CNode{
		Id:   1,
		Name: "rey",
	}

	node2 := &CNode{
		Id:   2,
		Name: "charlotte",
	}

	node3 := &CNode{
		Id:   3,
		Name: "martin",
	}
	node4 := &CNode{
		Id:   4,
		Name: "dora",
	}

	node5 := &CNode{
		Id:   4,
		Name: "lucy",
	}
	ShowAllCNode(head)
	InsertCNode(head, node1)
	ShowAllCNode(head)
	fmt.Println()
	InsertCNode(head, node2)
	ShowAllCNode(head)
	fmt.Println()
	InsertCNode(head, node3)
	ShowAllCNode(head)
	fmt.Println()
	InsertCNode(head, node4)
	ShowAllCNode(head)
	fmt.Println()
	InsertCNode(head, node5)
	ShowAllCNode(head)
	fmt.Println()

	head, err := DeleteCNode(head, 1)
	if err != nil {
		fmt.Printf("%v, where id = %d\n", err.Error(), 1)
	}
	ShowAllCNode(head)
	fmt.Println()

	head, err = DeleteCNode(head, 3)
	if err != nil {
		fmt.Printf("%v, where id = %d\n", err.Error(), 3)
	}
	ShowAllCNode(head)
	fmt.Println()

	head, err = DeleteCNode(head, 5)
	if err != nil {
		fmt.Printf("%v, where id = %d\n", err.Error(), 5)
	}
	ShowAllCNode(head)
	fmt.Println()
}

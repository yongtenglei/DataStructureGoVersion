package main

import (
	"errors"
	"fmt"
)

type BNode struct {
	Id       int
	NextNode *BNode
}

func AddBoy(num int) (*BNode, error) {
	if num < 0 {
		return nil, errors.New("can not add node less than 0")
	}

	head := &BNode{}
	current := &BNode{}
	for i := 1; i <= num; i++ {
		boy := &BNode{
			Id: i,
		}

		if i == 1 {
			head = boy
			current = boy
			current.NextNode = head
		} else {
			current.NextNode = boy
			current = boy
			boy.NextNode = head
		}
	}
	return head, nil
}

func ShowAllBNode(head *BNode) {
	if head.NextNode == nil {
		fmt.Println("There is nothing")
		return
	}

	current := head
	for {
		fmt.Printf("[id: %d] ==> ", current.Id)

		if current.NextNode == head {
			break
		}
		current = current.NextNode
	}
}
func CountBoy(head *BNode) int {
	counter := 0
	if head.NextNode == nil {
		fmt.Println("There is nothing")
		return 0
	}

	current := head
	for {
		if current.NextNode == head {
			break
		}
		counter++
		current = current.NextNode
	}
	return counter
}
func PlayGame(head *BNode, startId int, countNum int) {
	if head.NextNode == nil {
		fmt.Println("There is nothing, game over")
		return
	}

	if startId > CountBoy(head) {
		fmt.Printf("Cannot start with startId = %d, game over", startId)
		return
	}

	tail := head

	// let tail go to last position
	for {
		if tail.NextNode == head {
			break
		}
		tail = tail.NextNode
	}

	// let head go to BNode which have startId
	for i := 0; i < startId-1; i++ {
		head = head.NextNode
		tail = tail.NextNode
	}

	for {
		// count...
		for i := 0; i < countNum-1; i++ {
			head = head.NextNode
			tail = tail.NextNode
		}

		fmt.Printf("node %d leave away\n", head.Id)
		// delete boy
		head = head.NextNode
		tail.NextNode = head

		// only one boy here
		if head == tail {
			fmt.Printf("remain node %d\n", head.Id)
			return
		}
	}
}
func main() {
	head, err := AddBoy(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	ShowAllBNode(head)
	fmt.Println()
	PlayGame(head, 2, 3)
}

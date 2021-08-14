package main

import (
  "errors"
  "fmt"
)

type Queue struct {
  maxSize int
  array   [5]int
  front   int
  rear    int
}

func (q *Queue) QueueIsFull() bool {
  if q.rear == q.maxSize-1 {
    return true
  }
  return false
}

func (q *Queue) QueueIsEmpty() bool {
  if q.rear == q.front {
    return true
  }
  return false
}

func (q *Queue) AddQueue(val int) error {

  if q.QueueIsFull() {
    return errors.New("this queue is full")
  }

  q.rear++
  q.array[q.rear] = val
  fmt.Println("add successfully")
  return nil
}

func (q *Queue) Getqueue() (int, error) {

  if q.QueueIsEmpty() {
    return -1, errors.New("queue empty")
  }

  q.front++
  val := q.array[q.front]
  return val, nil
}

func (q *Queue) ShowQueue() {

  fmt.Println("Queue: ")

  for i := q.front; i < q.rear; i++ {
    fmt.Printf("arr[%d]=%d\n", i, q.array[i])
  }
}

func main() {
  queue := &Queue{
    maxSize: 5,
    front:   0,
    rear:    0,
  }

  var key int
  var val int
  for {
    fmt.Println("1 for add")
    fmt.Println("2 for get")
    fmt.Println("3 for show")
    fmt.Println("4 for exit")
    fmt.Scanln(&key)

    switch key {
    case 1:
      fmt.Println("Add： ")
      fmt.Scanln(&val)
      err := queue.AddQueue(val)
      if err == nil {
        fmt.Println("Add successfully")
      } else {
        fmt.Println(err.Error())
      }

    case 2:
      val, err := queue.Getqueue()
      if err != nil {
        fmt.Println(err.Error())
      } else {
        fmt.Printf("get %v\n", val)
      }

    case 3:
      queue.ShowQueue()

    case 4:
      break
    }
  }
}

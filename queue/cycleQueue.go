package main

import (
  "errors"
  "fmt"
)

type CycleQueue struct {
  maxSize int
  array   [5]int
  front   int
  rear    int
}

func (q *CycleQueue) CycQueueIsFull() bool {

  if (q.rear+1)%q.maxSize == q.front {
    return true
  }
  return false
}

func (q *CycleQueue) CycQueueIsEmpty() bool {
  if q.rear == q.front {
    return true
  }
  return false
}

func (q *CycleQueue) AddCycQueue(val int) error {

  if q.CycQueueIsFull() {
    return errors.New("this queue is full")
  }

  q.array[q.rear] = val
  q.rear = (q.rear + 1) % q.maxSize
  fmt.Println("add successfully")
  return nil
}

func (q *CycleQueue) GetCycQueue() (int, error) {

  if q.CycQueueIsEmpty() {
    return -1, errors.New("queue empty")
  }

  val := q.array[q.front]
  q.front = (q.front + 1) % q.maxSize
  return val, nil
}

func (q *CycleQueue) CycQueueSize() int {
  return (q.rear + q.maxSize - q.front) % q.maxSize

}

func (q *CycleQueue) ShowCycQueue() {

  fmt.Println("Queue: ")

  size := q.CycQueueSize()

  if size == 0 {
    fmt.Println("queue is empty")
  }

  tmp := q.front
  for i := 0; i < size; i++ {
    fmt.Printf("arr[%d]=%d\n", tmp, q.array[tmp])
    tmp = (tmp + 1) % q.maxSize
  }
  fmt.Println()
}

func main() {
  queue := &CycleQueue{
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
      err := queue.AddCycQueue(val)
      if err == nil {
        fmt.Println("Add successfully")
      } else {
        fmt.Println(err.Error())
      }

    case 2:
      val, err := queue.GetCycQueue()
      if err != nil {
        fmt.Println(err.Error())
      } else {
        fmt.Printf("get %v\n", val)
      }

    case 3:
      queue.ShowCycQueue()

    case 4:
      break
    }
  }
}

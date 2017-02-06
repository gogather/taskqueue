package taskqueue

import "testing"
import "fmt"

func Test_TaskQueue(t *testing.T) {
	tq := New()
	tq.Add("task1", "task1")
	tq.Add("task2", "task2")
	tq.Add("task3", "task3")

	fmt.Println(tq.Length())

	err := tq.Remove("task2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tq.Length())

	err = tq.Remove("task2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tq.Length())

	fmt.Println(tq.Top())
	fmt.Println(tq.Top())
	fmt.Println(tq.Top())

	fmt.Println(tq.Length())
}

func Test_TaskQueueThreadSafe(t *testing.T) {
	tq := New()

	count := 10000

	go func() {
		for i := 0; i < count; i++ {
			tq.Add(fmt.Sprintf("%d", i), i)
		}
	}()

	go func() {
		for true {
			notEmpty, taskID, task := tq.Top()
			if notEmpty {
				fmt.Println(notEmpty, taskID, task)
			}
		}
	}()

	for {
	}

}

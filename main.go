package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	logrus.Println("第一个mod项目测试")
	fmt.Println("测试一下")

	chanDemo()
}

func constDemo() {
	const (
		t1 = iota
		t2
		t3
		t4
	)
	fmt.Println(t4)
}

func chanDemo() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}()
outerLoop:
	for {
		select {
		case i, ok := <-ch:
			if ok {
				fmt.Println(i)
			} else {
				fmt.Println("关闭")
				break outerLoop
			}
		}
	}
}

func strDemo() {
	s := "中国人"
	b := []byte(s)
	r := []rune(s)
	fmt.Printf("0x%x\n", b)
	fmt.Printf("0x%x\n", r)
	fmt.Println(utf8.RuneCountInString(s))
	sb := strings.Builder{}
	sb.WriteString(s)
	sb.WriteString("最好")
	fmt.Println(sb.String())
}

func ExampleWithCancel() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				logrus.Printf("发射 %d", n)
				select {
				case <-ctx.Done():
					logrus.Println("外层取消了")
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			logrus.Println("结束")
			break
		}
	}
	cancel() // cancel when we are finished consuming integers
	time.Sleep(time.Second)
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleWithDeadline() {
	d := time.Now().Add(time.Millisecond * 2)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()
	c := make(chan struct{})
	go doSomeThing(ctx, c)
	select {
	case <-c:
		fmt.Println("按时完成")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	// Output:
	// context deadline exceeded
}

func doSomeThing(ctx context.Context, c chan<- struct{}) {
	logrus.Println("做事情....")
	time.Sleep(time.Millisecond * 3)
	c <- struct{}{}
}

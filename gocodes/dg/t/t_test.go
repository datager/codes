package t

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func BenchmarkTemplateParallel(b *testing.B) {
	tmpl := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			_ = tmpl.Execute(&buf, "World")
		}
	})
}

func ExampleHello() {
	fmt.Println("hello")
	// Output:hello
}

func ExampleSalutation() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}

func ExamplePerm() {
	for _, value := range rand.Perm(5) {
		fmt.Println(value)
	}
	// Unordered output: 4
	// 2
	// 1
	// 3
	// 0
}

func TestTimeConsuming(t *testing.T) {
	//if testing.Short() {
	//	t.Skip("skipping test in short mode.")
	//}
	t.Log("TestTimeConsuming")
}

func TestFoo(t *testing.T) {
	// setup code
	t.Run("A=1", func(t *testing.T) {
		t.Log("A=1111")
	})
	t.Run("A=2", func(t *testing.T) {
		t.Log("A=222")
	})
	t.Run("B=1", func(t *testing.T) {
		t.Log("B=333")
	})
	// tear down code
}

//func TestGroupedParallel(t *testing.T) {
//	for _,tc := range tests {
//		tc := tc
//		t.Run(tc.Name, func(t *testing.T) {
//			t.Parallel()
//		})
//	}
//}

func TestMain(m *testing.M) {
}

func TestSyncMap(t *testing.T) {
	m := sync.Map{}
	for i := 0; i < 10; i++ {
		m.Store(i, i)
	}
	m.Range(func(k, v interface{}) bool {
		fmt.Println(v)
		<-time.After(1 * time.Second)
		return true
	})
	<-time.After(2 * time.Second)
	m.Delete(1)
	m.Delete(2)
	m.Delete(3)
	m.Delete(4)
	m.Delete(5)
	m.Delete(6)
	m.Delete(7)
	m.Delete(8)
	m.Delete(9)
	m.Delete(10)

	for {
		time.Sleep(1 * time.Second)
	}
}

func TestChanCloseInSyncMap(t *testing.T) {
	type Session struct {
		sessionId string
		ch        chan int
	}

	var sessions *sync.Map = new(sync.Map)
	RemoveSession := func(sessionId string) {
		if session, ok := sessions.Load(sessionId); ok && session != nil {
			if session.(*Session).ch != nil {
				close(session.(*Session).ch)
				session.(*Session).ch = nil
			}
			sessions.Delete(sessionId)
		}
	}

	name := 1
	for name <= 1000 {
		session := &Session{
			sessionId: fmt.Sprintf("consumer_%d", name),
			ch:        make(chan int, 10),
		}
		name++
		sessions.Store(session.sessionId, session)
		go func(session *Session) {
			for {
				select {
				case num := <-session.ch:
					//fmt.Println("[out]", session.sessionId, num)
					if num > 10 {
						RemoveSession(session.sessionId)
					}
				}
			}
		}(session)
	}
	go func() {
		i := 0
		for {
			i++

			sessions.Range(func(sessionId, session interface{}) bool {
				//fmt.Printf("[in] sessionId:%v, i:%v\n", sessionId, i)
				select {
				case session.(*Session).ch <- i:
				default:
					fmt.Printf("[in-final] sessionId:%v, i:%v\n", sessionId, i)
					break
				}
				return true
			})
		}
	}()

	for {
		time.Sleep(3 * time.Second)
	}
}

package snowflake

import "testing"
import "fmt"
import "time"

func TestGen(t *testing.T) {
	node, _ := New(23)
	fmt.Printf("%b\n", node.Next())
	<-time.After(500 * time.Microsecond)
	fmt.Printf("%b\n", node.Next())
	<-time.After(500 * time.Microsecond)
	fmt.Printf("%b\n", node.Next())
	<-time.After(500 * time.Microsecond)
}

func TestUnixMillis(t *testing.T) {
	u := uint64(time.Now().Unix())
	fmt.Printf("unix: %064b\n", u)

	n := uint64(time.Now().UnixNano())
	fmt.Printf("nano: %064b\n", n)
	fmt.Printf("mili: %064b\n", n/1e6)
}

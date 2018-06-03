package snowflake

import "time"
import "sync"
import "fmt"

var (
	bitNode uint8  = 10
	bitStep uint8  = 12
	maxNode uint64 = 1<<(bitNode+1) - 1
	maxID   uint64 = 1<<(bitStep+1) - 1
	epoch   uint64
	nodes   = make(map[uint64]*Node, maxNode)
	mu      sync.Mutex

	timeShift = bitNode + bitStep // 时间戳向左的偏移量
	nodeShift = bitStep           // 节点 ID 向左的偏移量

	// ErrNodeInvalid node id is illegal
	ErrNodeInvalid = fmt.Errorf("node id must between 0...%d ", maxNode)
)

func init() {
	e, _ := time.Parse("2013-02-03 00:00:00 +0000 UTC", "2018-01-01 00:00:00 +0800 UTC")
	epoch = uint64(e.UnixNano() / 1e6)
}

// ID snowflake
type ID uint64

// Node generate ID
type Node struct {
	mu        sync.Mutex
	step      uint64
	node      uint64
	timestamp uint64
}

// New return a new node
func New(node uint64) (*Node, error) {
	if node < 0 || node > maxNode {
		return nil, ErrNodeInvalid
	}
	mu.Lock()
	defer mu.Unlock()

	n := nodes[node]
	if n != nil {
		return n, nil
	}

	n = &Node{node: node}
	nodes[node] = n
	return n, nil
}

// Next generate ID
func (n *Node) Next() ID {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := uint64(time.Now().UnixNano() / 1e6)

	if now == n.timestamp {
		n.step++
		if n.step > maxID {
			for now <= n.timestamp {
				now = uint64(time.Now().UnixNano() / 1e6)
			}
			n.step = 0
			n.timestamp = now
		}
	} else {
		n.step = 0
		n.timestamp = now
	}

	id := ID(((now - epoch) << timeShift) | (n.node << nodeShift) | (n.step))
	return id
}

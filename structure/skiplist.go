package structure

import (
	"math/rand"
	"time"
)

const (
	skipListInitLevel = 1  //跳跃表初始level
	skipListMaxLevel  = 32 //跳跃表最大的节点层数
	skipListP         = 0.25
)

type (
	SkipList struct {
		header *SkipListNode
		tail   *SkipListNode
		length uint
		level  int8
	}

	SkipListNode struct {
		ele      *sds
		score    float32
		backward *SkipListNode
		level    []SkipListLevel
	}

	SkipListLevel struct {
		forward *SkipListNode
		span    uint
	}
)

func NewSkipList() *SkipList {
	sl := &SkipList{
		header: NewSkipListNode(EmptySds(), 0, skipListMaxLevel),
		length: 0,
		level:  skipListInitLevel,
	}
	for i := 0; i < skipListMaxLevel; i++ {
		sl.header.level[i] = *&SkipListLevel{
			span: 0,
		}
	}
	return sl
}

func NewSkipListNode(ele *sds, score float32, level int8) *SkipListNode {
	return &SkipListNode{
		ele:   ele,
		score: score,
		level: make([]SkipListLevel, level),
	}
}

func randomLevel() int8 {
	level := skipListInitLevel
	for {
		rand.Seed(time.Now().UnixNano())
		if float64(rand.Int()&0xFFFF) < skipListP*0xFFFF {
			level++
		} else {
			break
		}
	}
	if level < skipListMaxLevel {
		return int8(level)
	}
	return skipListMaxLevel
}

func (sl *SkipList) Insert(ele *sds, score float32) {
	update, rank, node :=
		[skipListMaxLevel]*SkipListNode{},
		[skipListMaxLevel]uint{},
		sl.header

	for i := sl.level - 1; i >= 0; i-- {
		if i == sl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for node.level[i].forward != nil &&
			(score > node.level[i].forward.score ||
				(score == node.level[i].forward.score &&
					ele.Cmp(node.level[i].forward.ele) > 0)) {
			rank[i] += node.level[i].span
			node = node.level[i].forward
		}
		update[i] = node
	}

	level := randomLevel()
	node = NewSkipListNode(ele, score, level)

	if level > sl.level {
		for i := sl.level; i < level; i++ {
			rank[i], update[i] = 0, sl.header
			update[i].level[i].span = sl.length
		}
		sl.level = level
	}

	for i := int8(0); i < level; i++ {
		node.level[i].forward, update[i].level[i].forward = update[i].level[i].forward, node
		node.level[i].span, update[i].level[i].span = update[i].level[i].span-(rank[0]-rank[i]), rank[0]-rank[i]+1
	}

	for i := level; i < sl.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == sl.header {
		node.backward = nil
	} else {
		node.backward = update[0]
	}

	if node.level[0].forward == nil {
		sl.tail = node
	} else {
		node.level[0].forward.backward = node
	}

	sl.length++
}

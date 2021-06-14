package structure

import (
	"math/rand"
	"time"
)

const skipListMaxLevel = 32 //跳跃表最大的节点层数
const skipListP = 0.25

type (
	SkipList struct {
		header *SkipListNode
		tail   *SkipListNode
		length int
		level  int
	}

	SkipListNode struct {
		ele      sds
		score    float32
		backward *SkipListNode
		level    []SkipListLevel
	}

	SkipListLevel struct {
		forward *SkipListNode
		span    uint32
	}
)

func NewSkipList() *SkipList {
	sl := new(SkipList)
	sl.length, sl.level, sl.tail = 0, 1, nil
	sl.header = new(SkipListNode)
	sl.header.ele, sl.header.score, sl.header.backward = *new(sds), 0, nil
	for i := 0; i < skipListMaxLevel; i++ {
		sl.header.level = append(sl.header.level, *&SkipListLevel{
			forward: nil,
			span:    0,
		})
	}
	return sl
}

func NewSkipListNode(ele sds, score float32) *SkipListNode {
	return &SkipListNode{
		ele:   ele,
		score: score,
	}
}

func randomLevel() int {
	level := 1
	for {
		rand.Seed(time.Now().UnixNano())
		if float64(rand.Int()&0xFFFF) < skipListP*0xFFFF {
			level++
		} else {
			break
		}
	}
	if level < skipListMaxLevel {
		return level
	}
	return skipListMaxLevel
}

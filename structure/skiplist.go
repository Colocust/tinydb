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
		level  uint8
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
		header: NewSkipListNode(EmptySds(), 0),
		length: 0,
		level:  skipListInitLevel,
	}
	for i := 0; i < skipListMaxLevel; i++ {
		sl.header.level = append(sl.header.level, *&SkipListLevel{
			span: 0,
		})
	}
	return sl
}

func NewSkipListNode(ele *sds, score float32) *SkipListNode {
	return &SkipListNode{
		ele:   ele,
		score: score,
	}
}

func randomLevel() int {
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
		return level
	}
	return skipListMaxLevel
}

func (sl *SkipList) Insert(ele sds, score float32) {

}

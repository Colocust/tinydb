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

func (sl *SkipList) Insert(ele *sds, score float32) *SkipListNode {
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
					node.level[i].forward.ele.Cmp(ele) < 0)) {
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
	return node
}

func (sl *SkipList) Delete(ele *sds, score float32) (result bool, node *SkipListNode) {
	update, node, result := [skipListMaxLevel]*SkipListNode{}, sl.header, false

	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil &&
			(score > node.level[i].forward.score ||
				(score == node.level[i].forward.score &&
					node.level[i].forward.ele.Cmp(ele) > 0)) {
			node = node.level[i].forward
		}
		update[i] = node
	}

	node = node.level[0].forward

	if node != nil && node.score == score && node.ele.Cmp(ele) == 0 {
		sl.delete(node, update)
		result = true
		return
	}
	return
}

func (sl *SkipList) delete(node *SkipListNode, update [skipListMaxLevel]*SkipListNode) {
	for i := sl.level - 1; i >= 0; i-- {
		if update[i].level[i].forward == node {
			update[i].level[i].span += node.level[i].span - 1
			update[i].level[i].forward = node.level[i].forward
		} else {
			update[i].level[i].span--
		}
	}
	//如果被删除的node是最后一个元素
	if node.level[0].forward == nil {
		sl.tail = node.backward
	} else {
		node.level[0].forward.backward = node.backward
	}

	for sl.level > 1 && sl.header.level[sl.level-1].forward == nil {
		sl.level--
	}

	sl.length--
}

func (sl *SkipList) UpdateScore(ele *sds, curScore float32, newScore float32) *SkipListNode {
	update, node := [skipListMaxLevel]*SkipListNode{}, sl.header

	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil &&
			(curScore > node.level[i].forward.score ||
				(curScore == node.level[i].forward.score &&
					node.level[i].forward.ele.Cmp(ele) > 0)) {
			node = node.level[i].forward
		}
		update[i] = node
	}

	node = node.level[0].forward
	if node != nil && node.score == curScore && node.ele.Cmp(ele) == 0 {
		if (node.backward == nil || node.backward.score < newScore) &&
			(node.level[0].forward == nil || node.level[0].forward.score > newScore) {
			node.score = newScore
			return node
		}
		sl.delete(node, update)
		return sl.Insert(ele, newScore)
	}
	return nil
}

// IsInRange 判断跳表中是否由元素的值在指定区间内
func (sl *SkipList) IsInRange(zrs *ZRangeSpec) bool {
	if zrs.min > zrs.max || (zrs.min == zrs.max && (zrs.minex || zrs.maxex)) {
		return false
	}

	node := sl.tail
	if node == nil || !zrs.isValueGteMin(node.score) {
		return false
	}
	node = sl.header.level[0].forward
	if node == nil || !zrs.isValueLteMax(node.score) {
		return false
	}

	return true
}

// FirstInRange 寻找在指定区间内跳表中的第一个元素
func (sl *SkipList) FirstInRange(zrs *ZRangeSpec) *SkipListNode {
	var node *SkipListNode

	if !sl.IsInRange(zrs) {
		return nil
	}

	node = sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && !zrs.isValueGteMin(node.level[i].forward.score) {
			node = node.level[i].forward
		}
	}

	node = node.level[0].forward

	// 检查当前节点是否大于最大值
	if !zrs.isValueLteMax(node.score) {
		return nil
	}

	return node
}

func (sl *SkipList) LastInRange(zrs *ZRangeSpec) *SkipListNode {
	var node *SkipListNode

	if !sl.IsInRange(zrs) {
		return nil
	}

	node = sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && zrs.isValueLteMax(node.level[i].forward.score) {
			node = node.level[i].forward
		}
	}

	// 检查当前节点是否小于最小值
	if !zrs.isValueGteMin(node.score) {
		return nil
	}

	return node
}

func (sl *SkipList) DeleteByScore(zrs *ZRangeSpec, dict *Dict) int {
	update, node, removed := [skipListMaxLevel]*SkipListNode{}, sl.header, 0

	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && zrs.isValueLtMin(node.level[i].forward.score) {
			node = node.level[i].forward
		}
		update[i] = node
	}

	node = node.level[0].forward

	for node != nil && zrs.isValueLteMax(node.score) {
		next := node.level[0].forward
		sl.delete(node, update)
		dict.Remove(node.ele)
		removed++
		node = next
	}
	return removed
}

func (sl *SkipList) GetRank() {

}

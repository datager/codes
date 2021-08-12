package clients

import "container/heap"

type RankItemPriorityQueue []*RankItem

func (pq RankItemPriorityQueue) Len() int { return len(pq) }

func (pq RankItemPriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Score > pq[j].Score
}

func (pq RankItemPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *RankItemPriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*RankItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *RankItemPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *RankItemPriorityQueue) UpdateItemScore(item *RankItem, score float32) {
	item.Score = score
	heap.Fix(pq, item.index)
}

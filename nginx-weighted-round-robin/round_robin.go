package weighted_round_robin

import (
	"errors"
	"strconv"
)

/*
	go实现nginx polling 和weight polling
*/

//polling
type RoundRobinBalance struct {
	curIndex int
	rss      []string
}

func (rrb *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len is 1 at least")
	}
	//@todo 校验params参数addr
	for _, addr := range params {
		rrb.rss = append(rrb.rss, addr)
	}
	return nil
}
func (rrb *RoundRobinBalance) Next() string {
	lens := len(rrb.rss)
	if lens == 0 {
		return ""
	}
	if rrb.curIndex >= lens {
		rrb.curIndex = 0
	}
	curAddr := rrb.rss[rrb.curIndex]
	rrb.curIndex = (rrb.curIndex + 1) % lens
	return curAddr
}

//weight polling
/*
	算法逻辑：
	1. 轮询所有节点的effectiveWeight作为totalWeight
	2. 更新每个节点的currentWeight = currentWeight + effectiveWeight
		选中最大的currentWeight作为选中节点
	3. 选中的节点更新currentWeight = currentWeight - totalWeight
*/
type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
}

type WeightNode struct {
	weight          int    //配置的权重
	currentWeight   int    //节点当前权重
	effectiveWeight int    //节点有效权重，初始值为weight
	addr            string //服务器addr
}

func (r *WeightRoundRobinBalance) Add(params ...string) error {
	lens := len(params)
	if lens != 2 {
		return errors.New("params len need 2")
	}
	//获取addr和weight
	addr := params[0]
	wt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}

	node := &WeightNode{
		weight:          int(wt),
		effectiveWeight: int(wt),
		currentWeight:   int(wt),
		addr:            addr,
	}
	r.rss = append(r.rss, node)
	return nil
}

func (r *WeightRoundRobinBalance) Next() string {
	//服务为kong
	if len(r.rss) == 0 {
		return ""
	}

	//选中权重最大的节点 maxWeightNode
	totalWeight := 0
	var maxWeightNode *WeightNode
	for key, node := range r.rss {
		//计算当前状态所有节点的effectiveWeight为totalWeight
		totalWeight += node.effectiveWeight
		//更新每个节点的当前weight
		node.currentWeight = node.currentWeight + node.effectiveWeight
		//寻求权重最大的节点
		if maxWeightNode == nil || maxWeightNode.currentWeight < node.currentWeight {
			maxWeightNode = node
			r.curIndex = key
		}
	}
	//更新选中节点的当前权重
	maxWeightNode.currentWeight = maxWeightNode.currentWeight - totalWeight
	return maxWeightNode.addr
}

package hashmap

//hashmap的数组个数
const bucketCount = 20

type HashMap struct {
	//数组元素为连表的头指针
	Buckets [bucketCount]*LinkNode
}

//连表结构
type LinkNode struct {
	//存储key value
	Data KV
	//下一个节点
	Next *LinkNode
}

type KV struct {
	Key string
	Value string
}

func CreateLink() *LinkNode {
	//头结点数据为空 是为了标识这个链表还没有存储键值对
	return &LinkNode{
		KV{"", ""},
		nil,
	}
}

//添加节点，返回链表长度
func (link *LinkNode) AddNode(data KV) int {
	var count = 0
	//查找当前节点的尾节点
	tail := link
	for {
		count += 1
		if tail.Next == nil {
			break
		} else {
			tail = tail.Next
		}
	}
	var newNode = &LinkNode{data, nil}
	tail.Next = newNode
	return count + 1
}

//创建hashmap
func CreateHashMap() *HashMap {
	hmap := &HashMap{}
	//为数组每个bucket创建链表头
	for i := 0; i < bucketCount; i++ {
		hmap.Buckets[i] = CreateLink()
	}
	return hmap
}

//自定义散列算法将key转化到不同的bucket数组下标上
func HashCode(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}
	return sum % bucketCount
}

//往hashmap中添加key value
func (hmap *HashMap) AddKeyValue(key, value string) {
	//将key转成map数组的下标
	var index = HashCode(key)
	//获取对应数组的头节点
	link := hmap.Buckets[index]
	//判断当前链头节点是否为nil
	if link.Data.Key == "" && link.Next == nil {
		link.Data.Key = key
		link.Data.Value = value
	} else {
		link.AddNode(KV{key, value})
	}
}

//按key取值
func (hmap *HashMap) GetValeuOfKey(key string) string {
	index := HashCode(key)
	link := hmap.Buckets[index]
	//遍历链表获取value
	var value string
	head := link
	for {
		if head.Data.Key == key {
			value = head.Data.Value
			break
		} else {
			head = head.Next
		}
	}
	return value
}

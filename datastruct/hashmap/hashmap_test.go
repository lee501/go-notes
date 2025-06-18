package hashmap

import (
	"fmt"
	"testing"
)

func TestHashMap_AddKeyValue(t *testing.T) {
	hmap := CreateHashMap()
	fmt.Println(HashCode("test2"))
	hmap.AddKeyValue("test1", "value1")
	hmap.AddKeyValue("test2", "value2")
	fmt.Printf("hmap is %v\n", hmap.Buckets[17])
}

func TestHashMap_GetValeuOfKey(t *testing.T) {
	hmap := CreateHashMap()
	hmap.AddKeyValue("test1", "value1")
	fmt.Println(hmap.GetValeuOfKey("test1"))
}

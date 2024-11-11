package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed manu.json
var menuData []byte

type Tree struct {
	Path     string
	Label    string
	Children []*Tree
}

type Menu struct {
	ID        int
	Path      string
	Label     string
	ParentID  int
	Order     int
	Authority string
}

func main() {
	var data = make(map[string][]*Tree)
	if err := json.Unmarshal(menuData, &data); err != nil {
		fmt.Println(err)
		return
	}
	var manus []Menu
	var order = 1
	for key, trees := range data {
		generateMenus(key, &order, 0, trees, &manus)
	}
	fmt.Println(order)
	//fmt.Println(manus)
	order = 1
	fmt.Println(order)
}

func generateMenus(key string, order *int, parentId int, trees []*Tree, manus *[]Menu) {
	var menu Menu
	for _, tree := range trees {
		menu.ID = *order
		menu.Path = tree.Path
		menu.Label = tree.Label
		menu.ParentID = parentId
		menu.Order = *order
		menu.Authority = key
		*manus = append(*manus, menu)
		*order++
		if tree.Children != nil {
			generateMenus(key, order, menu.ID, tree.Children, manus)
		}
	}
}

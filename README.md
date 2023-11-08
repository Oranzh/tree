Tree
==============
Transform a two-dimensional array into a tree structure based on id and parent id value.

## Usage
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/oranzh/tree"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Region struct {
	Id     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Parent int    `gorm:"column:parent" json:"parent"`
}

type Regions struct {
	Region
	Children []*Regions `json:"children"`
}

func (r *Regions) GetID() int {
	return r.Id
}

func (r *Regions) GetParentID() int {
	return r.Parent
}

func (r *Regions) AddChild(node tree.Node) {
	r.Children = append(r.Children, node.(*Regions))
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var r []Region
	db.Find(&r)
	fmt.Printf("%+v", r)

	nodes := make([]tree.Node, len(r))
	for i := range r {
		nodes[i] = &Regions{
			Region: r[i],
		}
	}

	root := tree.GenerateTree(nodes)
	jsonTree, _ := json.MarshalIndent(root, "", "    ")
	fmt.Println(string(jsonTree))
}
```
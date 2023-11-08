package tree

import (
	"encoding/json"
	"testing"
)

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name" `
	ParentId int    `json:"parentId" `
}

type CategoryTree struct {
	Category
	SubCategories []*CategoryTree `json:"subCategories"`
}

func (c *CategoryTree) GetID() int {
	return c.Category.ID
}

func (c *CategoryTree) GetParentID() int {
	return c.Category.ParentId
}

func (c *CategoryTree) AddChild(node Node) {
	c.SubCategories = append(c.SubCategories, node.(*CategoryTree))
}

// newCategoryList return a list of CategoryTree
func newCategoryList() []*CategoryTree {
	return []*CategoryTree{
		{
			Category: Category{
				ID:       0,
				Name:     "root",
				ParentId: -1,
			},
		},
		{
			Category: Category{
				ID:       6,
				Name:     "six",
				ParentId: 0,
			},
		},
		{
			Category: Category{
				ID:       1,
				Name:     "one",
				ParentId: 0,
			},
		},

		{
			Category: Category{
				ID:       2,
				Name:     "two",
				ParentId: 1,
			},
		},
		{
			Category: Category{
				ID:       3,
				Name:     "three",
				ParentId: 1,
			},
		},
		{
			Category: Category{
				ID:       4,
				Name:     "four",
				ParentId: 2,
			},
		},
		{
			Category: Category{
				ID:       5,
				Name:     "five",
				ParentId: 2,
			},
		},
	}
}

func TestGenerateTree(t *testing.T) {
	input := newCategoryList()
	nodesList := make([]Node, len(input))
	for i, v := range input {
		nodesList[i] = v
	}
	tree := GenerateTree(nodesList)
	jsonTree, err := json.Marshal(tree)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(jsonTree))
}

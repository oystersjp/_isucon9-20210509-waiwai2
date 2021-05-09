package main

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
)

var allCategories []Category
var categoryByID map[int]Category
var childCategories map[int][]int

func setCategories() {
	err := dbx.Select(&allCategories, "SELECT * FROM categories")
	if err != nil {
		log.Fatal(err)
	}

	// カテゴリをID：配列で格納
	categoryByID = make(map[int]Category, len(allCategories))
	// カテゴリをID：配列（子カテゴリ）で格納
	childCategories = make(map[int][]int)
	for _, c := range allCategories {
		categoryByID[c.ID] = c
	}

	for _, c := range categoryByID {
		if c.ParentID > 0 {
			c.ParentCategoryName = categoryByID[c.ParentID].CategoryName
		}
		categoryByID[c.ID] = c

		var children []int
		for _, child := range categoryByID {
			if child.ParentID == c.ID {
				children = append(children, child.ID)
			}

			if len(children) > 0 {
				childCategories[c.ID] = children
			}
		}
	}
}

func getCategoryByID(_ sqlx.Queryer, id int) (Category, error) {
	c, ok := categoryByID[id]
	if !ok {
		return Category{}, sql.ErrNoRows
	}
	return c, nil
}

func getChildCategories(parent int) []int {
	return childCategories[parent]
}

func getAllCategories() []Category {
	return allCategories
}

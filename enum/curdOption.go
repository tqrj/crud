package enum

import "github.com/gin-gonic/gin"

type ListOption struct {
	Enable   bool
	Omit     []string
	LimitMax int
}

type GetOption struct {
	Enable bool
	Omit   []string
}

type UpdateOption struct {
	Enable bool
	Omit   []string
}

type CreateOption struct {
	Enable bool
	Omit   []string
}

type DelOption struct {
	Enable bool
}

// CrudGroup is options to construct the router group.
//
// By adding GetNested, CreateNested, DeleteNested to Crud,
// you can add CRUD routes for a nested model (Parent.Child).
//
// Or use CrudNested to add all three options above.
type CrudGroup func(group *gin.RouterGroup) *gin.RouterGroup

type CurdOption struct {
	ListOption
	GetOption
	UpdateOption
	CreateOption
	DelOption
}
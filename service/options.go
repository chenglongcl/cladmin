package service

import (
	"cladmin/util"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/schema"
)

type Option func(*Options)

type Options struct {
	IsHiddenCommentCount bool
	WithoutCount         bool
}

func OptionFindListWithoutCount(b bool) Option {
	return func(options *Options) {
		options.WithoutCount = b
	}
}

func OptionFindListIsHiddenCommentCount(b bool) Option {
	return func(options *Options) {
		options.IsHiddenCommentCount = b
	}
}

type ListParams struct {
	PS      util.PageSetting
	Options struct {
		WithoutCount  bool
		Scenes        string
		CustomDBOrder string
		CustomFunc    func() interface{}
	}
	Fields     []field.Expr
	Conditions []gen.Condition
	Joins      []struct {
		Table schema.Tabler
		On    []field.Expr
	}
	LeftJoins []struct {
		Table schema.Tabler
		On    []field.Expr
	}
	RightJoins []struct {
		Table schema.Tabler
		On    []field.Expr
	}
	Groups []field.Expr
	Orders []field.Expr
}

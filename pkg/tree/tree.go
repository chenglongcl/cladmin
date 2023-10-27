package tree

type Tree[T any] interface {
	IsEqual(father *T, child *T) bool
	SetChildren(father *T, children []*T)
	RetFather(father *T) bool
}

func ToTree[T any](data []*T, e Tree[T]) []*T {
	return filter(data, func(father *T) bool {
		children := filter(data, func(child *T) bool {
			return e.IsEqual(father, child)
		})
		if len(children) > 0 {
			e.SetChildren(father, children)
		}
		return e.RetFather(father)
	})
}

func filter[T any](data []*T, f func(item *T) bool) (res []*T) {
	for _, el := range data {
		if f(el) {
			res = append(res, el)
		}
	}
	return res
}

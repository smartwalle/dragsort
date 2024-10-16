package sortable

type Element interface {
	GetUniqueID() int64

	GetSortIndex() int

	UpdateSortIndex(sortIndex int)
}

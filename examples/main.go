package main

import (
	"fmt"
	"github.com/smartwalle/dragsort"
	"sort"
)

func main() {
	var ds = &DataSource{}
	for i := 1; i <= 10; i++ {
		ds.users = append(ds.users, &User{
			ID:        int64(i),
			SortIndex: i,
		})
	}

	dragsort.Sort(ds, ds.users[2], ds.users[1])
	dragsort.Sort(ds, ds.users[2], ds.users[3])
	dragsort.Sort(ds, ds.users[2], ds.users[4])
}

type User struct {
	ID        int64
	SortIndex int
}

func (u *User) String() string {
	return fmt.Sprintf("[%d-%d]", u.ID, u.SortIndex)
}

func (u *User) GetUniqueID() int64 {
	return u.ID
}

func (u *User) GetSortIndex() int {
	return u.SortIndex
}

func (u *User) UpdateSortIndex(sortIndex int) {
	u.SortIndex = sortIndex
}

type DataSource struct {
	users []*User
}

func (ds *DataSource) GetSortable(uniqueID int64) (dragsort.Sortable, error) {
	for _, u := range ds.users {
		if u.ID == uniqueID {
			return u, nil
		}
	}
	return nil, nil
}

func (ds *DataSource) GetSortableList(minSortIndex, maxSortIndex int) ([]dragsort.Sortable, error) {
	var elements = make([]dragsort.Sortable, 0, len(ds.users))
	for _, u := range ds.users {
		if u.SortIndex <= maxSortIndex && u.SortIndex >= minSortIndex {
			elements = append(elements, u)
		}
	}
	return elements, nil
}

func (ds *DataSource) UpateSortableList(elements []dragsort.Sortable) error {
	sort.SliceStable(ds.users, func(i, j int) bool {
		if ds.users[i].SortIndex < ds.users[j].SortIndex {
			return true
		}
		return false
	})
	fmt.Println("更新后的顺序：", ds.users)
	return nil
}

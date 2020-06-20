type Data struct {
}

type By func(s1, s2 *Data) bool

func (by By) Sort(data []Data) {
    sorter := &dataSorter{
        data: data,
        by: by,
    }
    sort.Sort(sorter)
}

type dataSorter struct {
    data []Data
    by func(s1, s2 *Data) bool
}

func (s *dataSorter) Len() int {
}

func (s *studentSorter) Less(i,j int) bool {
}

func (s *studentSorter) Swap(i, j int) bool {
}
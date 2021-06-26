package structure

type Dict struct {
	table map[*sds]interface{}
}

func NewDict() *Dict {
	return &Dict{
		table: make(map[*sds]interface{}),
	}
}
func (d *Dict) Set(key *sds, value interface{}) {
	d.table[key] = value
}

func (d *Dict) Get(key *sds) interface{} {
	return d.table[key]
}

func (d *Dict) Remove(key *sds) {
	delete(d.table, key)
}

func (d *Dict) GetLen() int {
	return len(d.table)
}

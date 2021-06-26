package structure

type Dict struct {
	table map[string]interface{}
}

func NewDict() *Dict {
	return &Dict{
		table: make(map[string]interface{}),
	}
}
func (d *Dict) Set(key string, value interface{}) {
	d.table[key] = value
}

func (d *Dict) Get(key string) interface{} {
	return d.table[key]
}

func (d *Dict) Remove(key string) {
	delete(d.table, key)
}

func (d *Dict) GetLen() int {
	return len(d.table)
}

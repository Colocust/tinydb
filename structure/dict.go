package structure

type dict struct {
	table map[string]interface{}
}

func NewDict() *dict {
	return &dict{
		table: make(map[string]interface{}),
	}
}
func (d *dict) Set(key string, value interface{}) {
	d.table[key] = value
}

func (d *dict) Get(key string) interface{} {
	return d.table[key]
}

func (d *dict) Remove(key string) {
	delete(d.table, key)
}

func (d *dict) GetLen() int {
	return len(d.table)
}

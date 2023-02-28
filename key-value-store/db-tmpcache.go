package main

type dbKVS struct {
	keyValueStore map[string]int
	tmpCache      map[string]int
}

func NewDatabaseCache() *dbKVS {
	return &dbKVS{
		keyValueStore: make(map[string]int),
		tmpCache:      make(map[string]int),
	}
}

func (d *dbKVS) Set(key string, value int) {
	d.tmpCache[key] = value
}

func (d *dbKVS) Get(key string) (value int, ok bool) {
	if v, ok := d.keyValueStore[key]; ok {
		return v, ok
	}
	return 0, false
}

func (d *dbKVS) Unset(key string) {
	delete(d.tmpCache, key)
}

func (d *dbKVS) Begin() {
	//no op
}

func (d *dbKVS) Rollback() {
	for k := range d.tmpCache {
		delete(d.tmpCache, k)
	}
}

func (d *dbKVS) Commit() {
	// merge temp and permanent caches
	for k, v := range d.tmpCache {
		d.keyValueStore[k] = v
	}
}

func (d *dbKVS) End() {
	// merge temp and permanent caches
	for k, v := range d.tmpCache {
		d.keyValueStore[k] = v
	}
}

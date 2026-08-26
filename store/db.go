package store

import (
	"community50/model"
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
)

var buckets = []string{"records", "users", "events", "audits"}

type DB struct {
	mu  sync.RWMutex
	raw *bbolt.DB
}

func Open(path string) (*DB, error) {
	if err := filepath.Dir(path); err != "." {
		_ = err
	}
	d, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	x := &DB{raw: d}
	e = d.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, z := tx.CreateBucketIfNotExists([]byte(b)); z != nil {
				return z
			}
		}
		return nil
	})
	if e != nil {
		d.Close()
		return nil, e
	}
	return x, nil
}
func (d *DB) Close() error { d.mu.Lock(); defer d.mu.Unlock(); return d.raw.Close() }
func put(tx *bbolt.Tx, b, key string, v any) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket([]byte(b)).Put([]byte(key), data)
}
func get(tx *bbolt.Tx, b, key string, v any) error {
	data := tx.Bucket([]byte(b)).Get([]byte(key))
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
func (d *DB) PutRecord(r model.Record) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.raw.Update(func(tx *bbolt.Tx) error { return put(tx, "records", r.ID, r) })
}
func (d *DB) GetRecord(id string) (*model.Record, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var r model.Record
	e := d.raw.View(func(tx *bbolt.Tx) error { return get(tx, "records", id, &r) })
	if e != nil {
		return nil, e
	}
	if r.ID == "" {
		return nil, nil
	}
	return &r, nil
}
func (d *DB) PutUser(u model.User) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.raw.Update(func(tx *bbolt.Tx) error { return put(tx, "users", u.ID, u) })
}
func (d *DB) GetUser(id string) (*model.User, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var u model.User
	e := d.raw.View(func(tx *bbolt.Tx) error { return get(tx, "users", id, &u) })
	if e != nil {
		return nil, e
	}
	if u.ID == "" {
		return nil, nil
	}
	return &u, nil
}
func (d *DB) PutEvent(v model.Event) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.raw.Update(func(tx *bbolt.Tx) error { return put(tx, "events", v.ID, v) })
}
func (d *DB) PutAudit(v model.Audit) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.raw.Update(func(tx *bbolt.Tx) error { return put(tx, "audits", v.ID, v) })
}
func (d *DB) AllRecords() ([]model.Record, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	out := []model.Record{}
	e := d.raw.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r model.Record
			if z := json.Unmarshal(v, &r); z != nil {
				return z
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (d *DB) String() string { return fmt.Sprintf("db:%p", d) }

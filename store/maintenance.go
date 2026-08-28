package store

import (
	"community50/model"
	"time"
)

func (d *DB) ReplaceStatus(id, status string) error {
	r, e := d.GetRecord(id)
	if e != nil {
		return e
	}
	if r == nil {
		return nil
	}
	r.Touch(status)
	return d.PutRecord(*r)
}
func (d *DB) SeedDemo() error {
	r := model.NewRecord("seed-50", "社区50值班公告", "今日服务时间", "system")
	r.CreatedAt = time.Now().UTC()
	return d.PutRecord(*r)
}

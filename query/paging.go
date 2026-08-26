package query

import "community50/model"

func Page(rs []model.Record, offset, size int) []model.Record {
	if offset < 0 {
		offset = 0
	}
	if offset >= len(rs) {
		return []model.Record{}
	}
	end := len(rs)
	if size > 0 && offset+size < end {
		end = offset + size
	}
	return rs[offset:end]
}

/* This file is auto-generated, manual edits in this file will be overwritten! */
package utils

import "github.com/samber/lo"

const INITIAL_PAGE = 0
const MAX_PAGE_SIZE = 15

func ConvertQueryPaging(page *int, size *int) (*int, *int) {
	if page == nil {
		page = lo.ToPtr(INITIAL_PAGE)
	}
	if size == nil {
		size = lo.ToPtr(MAX_PAGE_SIZE)
	}

	return page, size
}
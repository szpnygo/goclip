package clip

type searchOptions struct {
	pageNum  int32
	pageSize int32
}

type SearchOptionFunc func(opts *searchOptions)

func WithPageNum(pageNum int32) SearchOptionFunc {
	return func(opts *searchOptions) {
		opts.pageNum = pageNum
	}
}

func WithPageSize(pageSize int32) SearchOptionFunc {
	return func(opts *searchOptions) {
		opts.pageSize = pageSize
	}
}

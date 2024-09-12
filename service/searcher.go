package service

import (
	_ "embed"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
)

//go:embed ip2region.xdb
var ip2region_xdb []byte

type SearchService struct {
	searcher *xdb.Searcher
}

func (s *SearchService) Search(ip string) (string, error) {
	res, err := s.searcher.SearchByStr(ip)
	if err != nil {
		return "", err
	}

	arr := lo.Filter(strings.Split(res, "|"), func(item string, _ int) bool {
		return item != "0"
	})

	return strings.Join(lo.Union(arr), " "), nil
}

func NewSearchService(i do.Injector) (*SearchService, error) {
	searcher, err := xdb.NewWithBuffer(ip2region_xdb)
	if err != nil {
		return nil, err
	}
	return &SearchService{searcher}, nil
}

func init() {
	do.Provide(nil, NewSearchService)
}

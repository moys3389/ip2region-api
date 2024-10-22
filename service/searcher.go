package service

import (
	_ "embed"
	"log"
	"os"
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
	xdbFileName := "/data/ip2region.xdb"
	buff := ip2region_xdb
	_, err := os.Lstat(xdbFileName)
	if !os.IsNotExist(err) {
		log.Println("使用自定义xdb")
		buff, err = os.ReadFile(xdbFileName)
		if err != nil {
			return nil, err
		}
	}
	searcher, err := xdb.NewWithBuffer(buff)
	if err != nil {
		return nil, err
	}
	return &SearchService{searcher}, nil
}

func init() {
	do.Provide(nil, NewSearchService)
}

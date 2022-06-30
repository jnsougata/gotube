package search

import (
	"github.com/gotube/channel"
	"github.com/gotube/utils"
	"regexp"
)

type Search struct {
	keywords string
}

func New(SearchTerms string) *Search {
	return &Search{keywords: SearchTerms}
}

func (obj *Search) Channel() *channel.Channel {
	ep := "https://www.youtube.com/results?search_query=" + utils.KeywordParser(obj.keywords) + "&sp=EgIQAg%253D%253D"
	html := utils.FetchHtml(ep)
	re := regexp.MustCompile(`channelId":"(.*?)"`)
	url := "https://www.youtube.com/channel/" + re.FindStringSubmatch(html)[1]
	return channel.New(url)
}

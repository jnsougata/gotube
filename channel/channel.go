package channel

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Channel struct {
	html string
	root string
}

func (obj *Channel) _FetchHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		return string(body)
	} else {
		log.Fatal("Error: " + resp.Status)
	}
	return ""
}

func New(url string) *Channel {
	resp, err := http.Get(url + "/about")
	if err != nil {
		log.Fatal(err)
		return nil
	} else if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		return &Channel{html: string(body), root: url}
	} else {
		log.Fatal("Error: " + resp.Status)
		return nil
	}
}

func (obj *Channel) Name() string {
	re := regexp.MustCompile(`channelMetadataRenderer":{"title":"(.*?)"`)
	return re.FindStringSubmatch(obj.html)[1]
}

func (obj *Channel) Id() string {
	re := regexp.MustCompile(`channelId":"(.*?)"`)
	return re.FindStringSubmatch(obj.html)[1]
}

func (obj *Channel) Url() string {
	url := "https://www.youtube.com/channel/" + obj.Id()
	return url
}

func (obj *Channel) Totalviews() string {
	re := regexp.MustCompile(`viewCountText":{"simpleText":"(.*?)"}`)
	return strings.Split(re.FindStringSubmatch(obj.html)[1], " ")[0]
}

func (obj *Channel) Description() string {
	re := regexp.MustCompile(`{"description":{"simpleText":"(.*?)"}`)
	return re.FindStringSubmatch(obj.html)[1]
}

func (obj *Channel) Subscribers() string {
	re := regexp.MustCompile(`}},"simpleText":"(.*?) `)
	return re.FindStringSubmatch(obj.html)[1]
}

func (obj *Channel) Verified() bool {
	re := regexp.MustCompile(`"label":"Verified"`)
	return re.MatchString(obj.html)
}

func (obj *Channel) Avatar() string {
	re := regexp.MustCompile(`height":88},{"url":"(.*?)"`)
	return re.FindStringSubmatch(obj.html)[1]
}

func (obj *Channel) Banner() string {
	re := regexp.MustCompile(`width":1280,"height":351},{"url":"(.*?)"`)
	return re.FindStringSubmatch(obj.html)[1]
}

func (obj *Channel) Connections() string {
	re := regexp.MustCompile(`q=https%3A%2F%2F(.*?)"`)
	return "https://" + strings.Replace(re.FindStringSubmatch(obj.html)[1], "%2F", "/", 10)
}

func (obj *Channel) Country() string {
	re := regexp.MustCompile(`country":{"simpleText":"(.*?)"}`)
	return re.FindStringSubmatch(obj.html)[1]
}

func (obj *Channel) CustomUrl() string {
	re := regexp.MustCompile(`canonicalChannelUrl":"(.*?)"`)
	return re.FindStringSubmatch(obj.html)[1]
}

func (obj *Channel) CreationDate() string {
	re := regexp.MustCompile(`{"text":"Joined "},{"text":"(.*?)"}`)
	return re.FindStringSubmatch(obj.html)[1]
}

//Playlists To-Do: fix it
func (obj *Channel) Playlists() []string {
	html := obj._FetchHtml(obj.root + "/pls")
	re := regexp.MustCompile(`,"playlistId":"(.*?)"`)
	res := re.FindAllStringSubmatch(html, -1)
	empty := make([]string, len(res))
	for _, val := range res {
		if len(val[1]) > 0 {
			empty = append(empty, val[1])
		}
	}
	return empty
}

func (obj *Channel) IsLive() bool {
	html := obj._FetchHtml(obj.root + "/videos?view=2&live_view=501")
	re := regexp.MustCompile(`{"text":"LIVE"}`)
	return re.MatchString(html)
}

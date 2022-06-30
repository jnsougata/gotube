package channel

import (
	"github.com/gotube/utils"
	"regexp"
	"strings"
)

type Channel struct {
	html string
	url  string
}

func New(url string) *Channel {
	return &Channel{html: utils.FetchHtml(url + "/about"), url: url}
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

func (obj *Channel) TotalViews() string {
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

func (obj *Channel) Playlists() []string {
	html := utils.FetchHtml(obj.url + "/pls")
	re := regexp.MustCompile(`,"playlistId":"(.*?)"`)
	res := re.FindAllStringSubmatch(html, -1)
	return utils.Sanitize2D(res)
}

func (obj *Channel) IsLive() bool {
	html := utils.FetchHtml(obj.url + "/videos?view=2&live_view=501")
	re := regexp.MustCompile(`{"text":"LIVE"}`)
	return re.MatchString(html)
}

func (obj *Channel) Uploads() []string {
	html := utils.FetchHtml(obj.url + "/videos?view=0&sort=dd&flow=grid")
	re := regexp.MustCompile(`videoId":"(.*?)"`)
	res := re.FindAllStringSubmatch(html, -1)
	return utils.Sanitize2D(res)
}

func (obj *Channel) LatestUploaded() string {
	html := utils.FetchHtml(obj.url + "/videos?view=0&sort=dd&flow=grid")
	reUpChunk := regexp.MustCompile(`gridVideoRenderer":{(.*?)"navigationEndpoint`)
	reCheckStream := regexp.MustCompile(`simpleText":"Streamed`)
	reCheckLive := regexp.MustCompile(`default_live.`)
	fL1 := make([][]string, 0)
	fL2 := make([]string, 0)
	for _, v := range reUpChunk.FindAllStringSubmatch(html, -1) {
		if reCheckStream.MatchString(v[1]) != true {
			fL1 = append(fL1, v)
		}
	}
	for _, v := range fL1 {
		if reCheckLive.MatchString(v[1]) != true {
			fL2 = append(fL2, v[1])
		}
	}
	targetChunk := utils.Sanitize1D(fL2)[0]
	re := regexp.MustCompile(`videoId":"(.*?)"`)
	return re.FindStringSubmatch(targetChunk)[1]
}

func (obj *Channel) PersistentLiveStream() string {
	html := utils.FetchHtml(obj.url + "/videos?view=2&live_view=501")
	lre := regexp.MustCompile(`{"text":"LIVE"}`)
	if lre.MatchString(html) == true {
		re := regexp.MustCompile(`videoId":"(.*?)"`)
		res := re.FindStringSubmatch(html)
		return res[1]
	} else {
		return ""
	}
}

func (obj *Channel) LatestLiveStreamed() string {
	html := utils.FetchHtml(obj.url + "/videos?view=0&sort=dd&flow=grid")
	reUpChunk := regexp.MustCompile(`gridVideoRenderer":{(.*?)"navigationEndpoint`)
	reCheckStream := regexp.MustCompile(`simpleText":"Streamed`)
	reCheckLive := regexp.MustCompile(`default_live.`)
	fL1 := make([][]string, 0)
	fL2 := make([]string, 0)
	for _, v := range reUpChunk.FindAllStringSubmatch(html, -1) {
		if reCheckStream.MatchString(v[1]) == true {
			fL1 = append(fL1, v)
		}
	}
	for _, v := range fL1 {
		if reCheckLive.MatchString(v[1]) != true {
			fL2 = append(fL2, v[1])
		}
	}
	targetChunk := utils.Sanitize1D(fL2)[0]
	re := regexp.MustCompile(`videoId":"(.*?)"`)
	return re.FindStringSubmatch(targetChunk)[1]
}

func (obj *Channel) PreviousLiveStreams() []string {
	html := utils.FetchHtml(obj.url + "/videos?view=2&live_view=503")
	re := regexp.MustCompile(`videoId":"(.*?)"`)
	res := re.FindAllStringSubmatch(html, -1)
	return utils.Sanitize2D(res)
}

func (obj *Channel) UpcomingVideos() []string {
	html := utils.FetchHtml(obj.url + "/videos?view=2&live_view=502")
	reCheck := regexp.MustCompile(`"title":"Upcoming live streams"`)
	reSearch := regexp.MustCompile(`gridVideoRenderer:{"videoId":"(.*?)"`)
	if reCheck.MatchString(html) == true {
		return utils.Sanitize1D(reSearch.FindStringSubmatch(html))
	} else {
		return []string{}
	}
}

func (obj *Channel) Info() map[string]interface{} {
	contentMap := make(map[string]interface{})
	contentMap["name"] = obj.Name()
	contentMap["id"] = obj.Id()
	contentMap["subscribers"] = obj.Subscribers()
	contentMap["total_views"] = obj.TotalViews()
	contentMap["country"] = obj.Country()
	contentMap["custom_url"] = obj.CustomUrl()
	contentMap["creation_date"] = obj.CreationDate()
	contentMap["playlists"] = obj.Playlists()
	contentMap["verified"] = obj.Verified()
	contentMap["description"] = obj.Description()
	contentMap["avatar"] = obj.Avatar()
	contentMap["banner"] = obj.Banner()
	contentMap["url"] = obj.Url()
	return contentMap
}

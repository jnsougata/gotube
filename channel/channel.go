package channel

import (
	"github.com/gotube/https"
	"github.com/gotube/utils"
	"regexp"
	"strings"
)

type Channel struct {
	html string
	url  string
}

func New(url string) *Channel {
	return &Channel{html: https.GetChannelAbout(url), url: url}
}

func (obj *Channel) Name() string {
	re := regexp.MustCompile(`channelMetadataRenderer":{"title":"(.*?)"`)
	nameArray := re.FindStringSubmatch(obj.html)
	if len(nameArray) > 1 {
		return nameArray[1]
	} else {
		return ""
	}
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
	descArray := re.FindStringSubmatch(obj.html)
	if len(descArray) > 1 {
		return descArray[1]
	} else {
		return ""
	}
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
	bannerArray := re.FindStringSubmatch(obj.html)
	if len(bannerArray) > 1 {
		return bannerArray[1]
	} else {
		return ""
	}
}

func (obj *Channel) Connections() string {
	re := regexp.MustCompile(`q=https%3A%2F%2F(.*?)"`)
	connArray := re.FindStringSubmatch(obj.html)
	if len(connArray) > 1 {
		return "https://" + strings.Replace((connArray)[1], "%2F", "/", -1)
	} else {
		return ""
	}
}

func (obj *Channel) Country() string {
	re := regexp.MustCompile(`country":{"simpleText":"(.*?)"}`)
	countryArray := re.FindStringSubmatch(obj.html)
	if len(countryArray) > 1 {
		return countryArray[1]
	} else {
		return ""
	}
}

func (obj *Channel) CustomUrl() string {
	re := regexp.MustCompile(`canonicalChannelUrl":"(.*?)"`)
	curlArray := re.FindStringSubmatch(obj.html)
	if len(curlArray) > 1 {
		return curlArray[1]
	} else {
		return ""
	}
}

func (obj *Channel) CreationDate() string {
	re := regexp.MustCompile(`{"text":"Joined "},{"text":"(.*?)"}`)
	crArray := re.FindStringSubmatch(obj.html)
	if len(crArray) > 1 {
		return crArray[1]
	} else {
		return ""
	}
}

func (obj *Channel) Playlists() []string {
	re := regexp.MustCompile(`,"playlistId":"(.*?)"`)
	res := re.FindAllStringSubmatch(https.GetChannelPlaylists(obj.url), -1)
	return utils.Sanitize2D(res)
}

func (obj *Channel) IsLive() bool {
	re := regexp.MustCompile(`{"text":"LIVE"}`)
	return re.MatchString(https.GetLivestreamData(obj.url))
}

func (obj *Channel) Uploads() []string {
	re := regexp.MustCompile(`videoId":"(.*?)"`)
	res := re.FindAllStringSubmatch(https.GetChannelUploads(obj.url), -1)
	return utils.Sanitize2D(res)
}

func (obj *Channel) LatestUploaded() string {
	reUpChunk := regexp.MustCompile(`gridVideoRenderer":{(.*?)"navigationEndpoint`)
	reCheckStream := regexp.MustCompile(`simpleText":"Streamed`)
	reCheckLive := regexp.MustCompile(`default_live.`)
	fL1 := make([][]string, 0)
	fL2 := make([]string, 0)
	for _, v := range reUpChunk.FindAllStringSubmatch(https.GetChannelUploads(obj.url), -1) {
		if reCheckStream.MatchString(v[1]) != true {
			fL1 = append(fL1, v)
		}
	}
	for _, v := range fL1 {
		if reCheckLive.MatchString(v[1]) != true {
			fL2 = append(fL2, v[1])
		}
	}
	targetChunk := fL2[0]
	re := regexp.MustCompile(`videoId":"(.*?)"`)
	return re.FindStringSubmatch(targetChunk)[1]
}

func (obj *Channel) PersistentLiveStream() string {
	html := https.GetLivestreamData(obj.url)
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
	reUpChunk := regexp.MustCompile(`gridVideoRenderer":{(.*?)"navigationEndpoint`)
	reCheckStream := regexp.MustCompile(`simpleText":"Streamed`)
	reCheckLive := regexp.MustCompile(`default_live.`)
	fL1 := make([][]string, 0)
	fL2 := make([]string, 0)
	for _, v := range reUpChunk.FindAllStringSubmatch(https.GetChannelUploads(obj.url), -1) {
		if reCheckStream.MatchString(v[1]) == true {
			fL1 = append(fL1, v)
		}
	}
	for _, v := range fL1 {
		if reCheckLive.MatchString(v[1]) != true {
			fL2 = append(fL2, v[1])
		}
	}
	targetChunk := fL2[0]
	re := regexp.MustCompile(`videoId":"(.*?)"`)
	return re.FindStringSubmatch(targetChunk)[1]
}

func (obj *Channel) PreviousLiveStreams() []string {
	re := regexp.MustCompile(`videoId":"(.*?)"`)
	res := re.FindAllStringSubmatch(https.GetPastLiveStreams(obj.url), -1)
	return utils.Sanitize2D(res)
}

func (obj *Channel) UpcomingVideos() []string {
	html := https.GetUpcomingVideos(obj.url)
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
	contentMap["Name"] = obj.Name()
	contentMap["Id"] = obj.Id()
	contentMap["Subscribers"] = obj.Subscribers()
	contentMap["TotalViews"] = obj.TotalViews()
	contentMap["Country"] = obj.Country()
	contentMap["CustomUrl"] = obj.CustomUrl()
	contentMap["CreationDate"] = obj.CreationDate()
	contentMap["Playlists"] = obj.Playlists()
	contentMap["verified"] = obj.Verified()
	contentMap["Description"] = obj.Description()
	contentMap["Avatar"] = obj.Avatar()
	contentMap["Banner"] = obj.Banner()
	contentMap["Url"] = obj.Url()
	return contentMap
}

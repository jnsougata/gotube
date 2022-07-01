package https

import "github.com/gotube/utils"

func GetChannelAbout(url string) string {
	return utils.FetchHtml(url + "/about")
}

func GetChannelPlaylists(url string) string {
	return utils.FetchHtml(url + "/playlists")
}

func GetLivestreamData(url string) string {
	return utils.FetchHtml(url + "/videos?view=2&live_view=501")
}

func GetChannelUploads(url string) string {
	return utils.FetchHtml(url + "/videos?view=0&sort=dd&flow=grid")
}

func GetPastLiveStreams(url string) string {
	return utils.FetchHtml(url + "/videos?view=2&live_view=503")
}

func GetUpcomingVideos(url string) string {
	return utils.FetchHtml(url + "/videos?view=2&live_view=502")
}

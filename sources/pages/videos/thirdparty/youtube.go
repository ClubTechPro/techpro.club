package thirdparty

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
	"golang.org/x/oauth2"
	"techpro.club/sources/common"
)

type YoutubeOutputStruct struct{
	VideoId string `json:"videoId"`
	VideoTitle string `json:"videoTitle"`
	VideoDescription string `json:"videoDescription"`
	VideoPublishedAt string `json:"videoPublishedAt"`
	VideoThumbnail string `json:"videoThumbnail"`
	VideoStatus string `json:"videoStatus"`
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {

	f, err := os.Open(common.CONST_SESSION_PATH + file)
	if err != nil {
			return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// Fetch my channel
func FetchMyChannel(login string)(status bool, msg string, channelId string){

	var uploadID string
	status = false
	msg = ""

	tokenFile := login + ".json"
	token, err := tokenFromFile(tokenFile)
	if err != nil {
		msg = err.Error()
	} else {

		reqUrl := "https://youtube.googleapis.com/youtube/v3/channels?part=snippet%2CcontentDetails%2Cstatistics&mine=true&key="+token.AccessToken
		req, err := http.NewRequest("GET", reqUrl, nil)
		if err != nil {
			msg = err.Error()
		} else {
			req.Header.Add("Authorization", "Bearer "+token.AccessToken)
			req.Header.Add("Accept", "application/json")
			req.Header.Add("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				msg = err.Error()
			} else {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					msg = err.Error()
				} else {
					bodyString := string(body)
					uploadID = gjson.Get(bodyString, "items.0.contentDetails.relatedPlaylists.uploads").String()

					msg = ""

				}
			}
		}
	}

	return status, msg, uploadID
}


// Fetch videos from my channel
func FetchMyVideos(channelId, login, pageToken string)(status bool, msg string, prevPageToken, nextPageToken string, videoList []YoutubeOutputStruct){
	status = false
	msg = ""

	var totalResults int64

	tokenFile := login + ".json"
	token, err := tokenFromFile(tokenFile)
	if err != nil {
		msg = err.Error()
	} else {

		reqUrl := "https://youtube.googleapis.com/youtube/v3/playlistItems?part=contentDetails%2Cstatus%2Csnippet%2Cid&playlistId=UUw2rxHtnlUfRykkLcqV4J9g&maxResults=10&key="+token.AccessToken +"pageToken="+pageToken
		req, err := http.NewRequest("GET", reqUrl, nil)
		if err != nil {
			msg = err.Error()
		} else {
			req.Header.Add("Authorization", "Bearer "+token.AccessToken)
			req.Header.Add("Accept", "application/json")
			req.Header.Add("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				msg = err.Error()
			} else {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					msg = err.Error()
				} else {
					bodyString := string(body)
					totalResults = gjson.Get(bodyString, "pageInfo.totalResults").Int()
					nextPageToken = gjson.Get(bodyString, "nextPageToken").String()
					prevPageToken = gjson.Get(bodyString, "prevPageToken").String()
					for i := 0; i < int(totalResults); i++ {
						
						videoId := gjson.Get(bodyString, "items."+fmt.Sprint(i)+".contentDetails.videoId").String()
						videoTitle := gjson.Get(bodyString, "items."+fmt.Sprint(i)+".snippet.title").String()
						videoDescription := gjson.Get(bodyString, "items."+fmt.Sprint(i)+".snippet.description").String()
						videoPublishedAt := gjson.Get(bodyString, "items."+fmt.Sprint(i)+".snippet.publishedAt").String()
						videoThumbnail := gjson.Get(bodyString, "items."+fmt.Sprint(i)+".snippet.thumbnails.medium.url").String()
						videoStatus := gjson.Get(bodyString, "items."+fmt.Sprint(i)+".status.privacyStatus").String()

						videoList = append(videoList, YoutubeOutputStruct{
							VideoId: videoId,
							VideoTitle: videoTitle,
							VideoDescription: videoDescription,
							VideoPublishedAt: videoPublishedAt,
							VideoThumbnail: videoThumbnail,
							VideoStatus: videoStatus,
						})

						fmt.Println(videoTitle)
					}
					msg = ""
					status = true

					
				}
			}
		}
	}

	return status, msg, prevPageToken, nextPageToken, videoList
}
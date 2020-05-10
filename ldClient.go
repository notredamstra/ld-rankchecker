package ldjam_rank

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const ldStaticBaseUrl string = "http://static.jam.vg/"

type Client struct {
	BaseURL   string
	httpClient *http.Client
}

func NewClient(baseURL string, httpClient *http.Client) *Client {
	return &Client{BaseURL: baseURL, httpClient: httpClient}
}

func (c *Client) GetGameRank(game *LDGame) (int, error){
	offset := 0
	limit := 250

	rank, err := c.findGamePositionInFeed(game, offset, limit)

	if err != nil {
		return -1, err
	}

	return rank, err
}

func (c *Client) findGamePositionInFeed(game *LDGame, offset int, limit int) (int, error){
	fmt.Println("Offset: "+strconv.Itoa(offset))
	path := url.URL{Path: "/vx/node/feed/" + strconv.Itoa(game.EventId) + "/grade+parent/item/game/" + game.Type, RawQuery: "offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit)} //todo remove ugly way of building url

	// get raw response data
	resp, err := c.httpClient.Get(c.BaseURL + path.RequestURI())
	fmt.Println(c.BaseURL + path.RequestURI())
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// transform raw data into slice of bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// parse slice of bytes into LD response struct
	var page *LDGameFeedPage
	err = json.Unmarshal(body, &page)
	if err != nil {
		fmt.Println(err)
	}

	for i, g := range page.Feed {
		if g.ID == game.Id {
			idx := offset + i + 1
			return idx, err
		}
	}

	return c.findGamePositionInFeed(game, offset + limit, limit)
}

func (c *Client) GetEventStatsFromGame(game *LDGame) (*LDEvent, error){
	path := url.URL{Path: "/vx/stats/" + strconv.Itoa(game.EventId)}

	// get raw response data
	resp, err := c.httpClient.Get(c.BaseURL + path.RequestURI())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// transform raw data into slice of bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse slice of bytes into LD response struct
	var ldResp *LDStats
	err = json.Unmarshal(body, &ldResp)
	if err != nil {
		return nil, err
	}

	event := LDEvent{}

	if ldResp.Type == "event"{
		event = LDEvent{
			Id:          ldResp.ID,
			Type:        ldResp.Subtype,
			Signups:     ldResp.Stats.Signups,
			Authors:     ldResp.Stats.Authors,
			Unpublished: ldResp.Stats.Unpublished,
			Game:        ldResp.Stats.Game,
			Craft:       ldResp.Stats.Craft,
			Tool:        ldResp.Stats.Tool,
			Demo:        ldResp.Stats.Demo,
			Jam:         ldResp.Stats.Jam,
			Compo:       ldResp.Stats.Compo,
			Warmup:      ldResp.Stats.Warmup,
			Late:        ldResp.Stats.Late,
			Release:     ldResp.Stats.Release,
			Unfinished:  ldResp.Stats.Unfinished,
		}
	}

	return &event, err
}

func (c *Client) GetGameFromURL(userUrl *url.URL) (*LDGame, error){
	nodePath := userUrl.Path
	path := url.URL{Path: "/vx/node2/walk/1" + nodePath, RawQuery: "node&parent"}

	// get raw response data
	resp, err := c.httpClient.Get(c.BaseURL + path.RequestURI())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// transform raw data into slice of bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse slice of bytes into LD response struct
	var ldResp *LDNodes
	err = json.Unmarshal(body, &ldResp)
	if err != nil {
		return nil, err
	}

	game := LDGame{}

	if len(ldResp.Nodes) == 2 && ldResp.Nodes[0].Subtype == "game" {
		node := &ldResp.Nodes[0]
		game = LDGame{
			Id:          node.Id,
			Name:        node.Name,
			Description: node.Description,
			Grade:       node.Magic.Grade,
			Cover:       genCoverURI(node.Meta.Cover),
			EventId:     node.Parent,
			Type:     	 strings.Title(strings.ToLower(node.Subsubtype)),
		}

		game.Rank, err = c.GetGameRank(&game)
	}

	return &game, err
}

func genCoverURI(rel string) string{
	rel = strings.TrimPrefix(rel, "///")
	return ldStaticBaseUrl+rel
}
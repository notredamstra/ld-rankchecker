package ldmdare

import (
	"net/http"
	"net/url"
	"time"
)

type LudumApi interface {
	GetGameRank(game *LDGame) (int, error)
	GetEventStatsFromGame(game *LDGame) (*LDEvent, error)
	GetGameFromURL(userUrl *url.URL) (*LDGame, error)
}

type Client struct {
	BaseURL    string
	httpClient *http.Client
}

type LDNodes struct {
	Nodes []LDNode `json:"node"`
}

type LDNode struct {
	Id          int    `json:"id"`
	Type        string `json:"type"`
	Subtype     string `json:"subtype"`
	Subsubtype  string `json:"subsubtype"`
	Name        string `json:"name"`
	Description string `json:"body"`
	Magic       struct {
		Grade    float64 `json:"grade"`
		Cool     float64 `json:"cool"`
		Feedback float64 `json:"feedback"`
		Smart    float64 `json:"smart"`
		Given    float64 `json:"given"`
	}
	Meta struct {
		Cover string `json:"cover"`
	}
	Parent int `json:"parent"`
}

type LDStats struct {
	Status   int    `json:"status"`
	CallerID int    `json:"caller_id"`
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Subtype  string `json:"subtype"`
	Stats    struct {
		Signups     int       `json:"signups"`
		Authors     int       `json:"authors"`
		Unpublished int       `json:"unpublished"`
		Game        int       `json:"game"`
		Craft       int       `json:"craft"`
		Tool        int       `json:"tool"`
		Demo        int       `json:"demo"`
		Jam         int       `json:"jam"`
		Compo       int       `json:"compo"`
		Warmup      int       `json:"warmup"`
		Late        int       `json:"late"`
		Release     int       `json:"release"`
		Unfinished  int       `json:"unfinished"`
		Grade20Plus int       `json:"grade-20-plus"`
		Grade1520   int       `json:"grade-15-20"`
		Grade1015   int       `json:"grade-10-15"`
		Grade510    int       `json:"grade-5-10"`
		Grade05     int       `json:"grade-0-5"`
		Grade0Only  int       `json:"grade-0-only"`
		Timestamp   time.Time `json:"timestamp"`
	} `json:"stats"`
}

type LDGame struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Grade       float64 `json:"grade"`
	Cover       string  `json:"cover"`
	EventId     int     `json:"event_id"`
	Type        string  `json:"type"`
	Rank        int     `json:"rank"`
}

type LDGameFeedPage struct {
	Status      int      `json:"status"`
	CallerID    int      `json:"caller_id"`
	Root        int      `json:"root"`
	Method      []string `json:"method"`
	Types       []string `json:"types"`
	Subtypes    []string `json:"subtypes"`
	Subsubtypes []string `json:"subsubtypes"`
	Offset      int      `json:"offset"`
	Limit       int      `json:"limit"`
	Feed        []struct {
		ID       int       `json:"id"`
		Modified time.Time `json:"modified"`
		Score    float64   `json:"score"`
	} `json:"feed"`
}

type LDEvent struct {
	Id          int    `json:"id"`
	Type        string `json:"type"`
	Signups     int    `json:"signups"`
	Authors     int    `json:"authors"`
	Unpublished int    `json:"unpublished"`
	Game        int    `json:"game"`
	Craft       int    `json:"craft"`
	Tool        int    `json:"tool"`
	Demo        int    `json:"demo"`
	Jam         int    `json:"jam"`
	Compo       int    `json:"compo"`
	Warmup      int    `json:"warmup"`
	Late        int    `json:"late"`
	Release     int    `json:"release"`
	Unfinished  int    `json:"unfinished"`
}

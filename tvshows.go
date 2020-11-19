package tvgo

import (
	"fmt"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func tvshowsUUID() (UUID string) {
	aseed := time.Now()
	aSeed := aseed.UnixNano()
	rand.Seed(aSeed)
	u := rand.Int63n(aSeed)
	UUID = strconv.FormatInt(u, 10)
	return
}

// TVShowInfoS is needed because I want it
type TVShowInfoS struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	FilePath      string        `bson:"filepath"`
	Catagory      string        `bson:"catagory"`
	MediaID       string        `bson:"MediaID"`
	Genre         string        `bson:"genre"`
	Season        string        `bson:"season"`
	Episode       string        `bson:"episode"`
	Title         string        `bson:"title"`
	Series        string        `bson:"series"`
	TVShowPicPath string        `bson:"tvshowpicpath"`
	ThumbPath     string        `bson:"thumbpath"`
	TvFSPath      string        `bson:"tvfspath"`
}

func getTvShowInfo(apath string, tvshowpicPath string) (TvSI TVShowInfoS) {
	// _, filename := path.Split(apath)
	// // /root/fsData/TVShows/Enterprise/S1/filename.mp4
	// fspath := apath[20:]
	// boo := len(filename) - 4
	// TvSI.ID = bson.NewObjectId()
	// TvSI.FilePath = apath
	// TvSI.MediaID = tvshowsUUID()
	// TvSI.Genre = "TVShows"
	// TvSI.TVShowPicPath = tvshowpicPath
	// TvSI.TvFSPath = fspath
	switch {
	case strings.Contains(apath, "TVShows/TNG"):
		_, filename := path.Split(apath)
		// /root/fsData/TVShows/Enterprise/S1/filename.mp4
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "TNG"
		TvSI.Season = filename[15:17]
		TvSI.Episode = filename[18:20]
		TvSI.Title = filename[21:boo]
		TvSI.Series = filename[21:boo]
	case strings.Contains(apath, " STTV "):
		_, filename := path.Split(apath)
		// /root/fsData/TVShows/Enterprise/S1/filename.mp4
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "STTV"
		TvSI.Season = filename[16:18]
		TvSI.Episode = filename[19:21]
		TvSI.Title = filename[21:boo]
		TvSI.Series = "Star Trek"
	case strings.Contains(apath, "Orville"):
		_, filename := path.Split(apath)
		// /root/fsData/TVShows/Enterprise/S1/filename.mp4
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Orville"
		TvSI.Season = filename[13:15]
		TvSI.Episode = filename[16:18]
		TvSI.Title = filename[19:boo]
		TvSI.Series = "The Orville"
	case strings.Contains(apath, "Voyager"):
		_, filename := path.Split(apath)
		// /root/fsData/TVShows/Enterprise/S1/filename.mp4
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Voyager"
		TvSI.Season = filename[19:21]
		TvSI.Episode = filename[22:24]
		TvSI.Title = filename[24:boo]
		TvSI.Series = "Voyager"
	case strings.Contains(apath, "Discovery"):
		_, filename := path.Split(apath)
		// /root/fsData/TVShows/Enterprise/S1/filename.mp4
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Discovery"
		TvSI.Season = filename[21:23]
		TvSI.Episode = filename[24:26]
		TvSI.Title = filename[27:boo]
		TvSI.Series = "Discovery"
	case strings.Contains(apath, "ENT"):
		_, filename := path.Split(apath)
		// /root/fsData/TVShows/Enterprise/S1/filename.mp4
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Enterprise"
		TvSI.Season = filename[15:17]
		TvSI.Episode = filename[18:20]
		TvSI.Title = filename[20:boo]
		TvSI.Series = "Enterprise"
	case strings.Contains(apath, "The Last Ship"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Last Ship"
		TvSI.Season = filename[15:17]
		TvSI.Episode = filename[18:20]
		TvSI.Title = filename[21:boo]
		TvSI.Series = "The Last Ship"
	case strings.Contains(apath, "Lost In Space"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Lost In Space"
		TvSI.Season = filename[15:17]
		TvSI.Episode = filename[18:20]
		TvSI.Title = filename[21:boo]
		TvSI.Series = "Lost In Space"
	case strings.Contains(apath, "Picard"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Picard"
		TvSI.Season = filename[18:20]
		TvSI.Episode = filename[21:23]
		TvSI.Title = filename[24:boo]
		TvSI.Series = "Picard"
	case strings.Contains(apath, "Mandalorian"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Mandalorian"
		TvSI.Season = filename[17:19]
		TvSI.Episode = filename[20:22]
		TvSI.Title = filename[23:boo]
		TvSI.Series = "Mandalorian"

	case strings.Contains(apath, "Lower Decks"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "LowerDecks"
		TvSI.Season = filename[23:25]
		TvSI.Episode = filename[26:28]
		TvSI.Title = filename[29:boo]
		TvSI.Series = "LowerDecks"

	case strings.Contains(apath, "Altered Carbon"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "AlteredCarbon"
		TvSI.Season = filename[16:18]
		TvSI.Episode = filename[19:21]
		TvSI.Title = filename[21:boo]
		TvSI.Series = "AlteredCarbon"

	case strings.Contains(apath, "Raised By Wolves"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "RaisedByWolves"
		TvSI.Season = filename[16:18]
		TvSI.Episode = filename[19:21]
		TvSI.Title = filename[21:boo]
		TvSI.Series = "RaisedByWolves"
	}
	fmt.Printf("\n THIS IS TVI FROM TVSHOWS \n %s \n", TvSI)
	return
}

package tvgo

import (
	"fmt"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"
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
	switch {
	case strings.Contains(apath, "TVShows/TNG"):
		_, filename := path.Split(apath)
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
		TvSI.Season = filename[18:20]
		TvSI.Episode = filename[21:23]
		TvSI.Title = filename[23:boo]
		TvSI.Series = "RaisedByWolves"

	case strings.Contains(apath, "For All Mankind"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "ForAllManKind"
		TvSI.Season = filename[17:19]
		TvSI.Episode = filename[20:22]
		TvSI.Title = filename[22:boo]
		TvSI.Series = "ForAllManKind"

	case strings.Contains(apath, "Alien Worlds"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "AlienWorlds"
		TvSI.Season = filename[14:16]
		TvSI.Episode = filename[17:19]
		TvSI.Title = filename[20:boo]
		TvSI.Series = "AlienWorlds"

	case strings.Contains(apath, "WandaVision"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "WandaVision"
		TvSI.Season = filename[13:15]
		TvSI.Episode = filename[16:18]
		TvSI.Title = filename[19:boo]
		TvSI.Series = "WandaVision"

	case strings.Contains(apath, "FalconWinterSoldier"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "FalconWinterSoldier"
		TvSI.Season = filename[35:37]
		TvSI.Episode = filename[38:40]
		TvSI.Title = filename[40:boo]
		TvSI.Series = "FalconWinterSoldier"

	case strings.Contains(apath, "Loki"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Loki"
		TvSI.Season = filename[6:8]
		TvSI.Episode = filename[9:11]
		TvSI.Title = filename[11:boo]
		TvSI.Series = "Loki"
		fmt.Printf("\n THIS IS TVI FROM Loki \n %s \n", TvSI)
		fmt.Println(TvSI.Season)
		fmt.Println(TvSI.Episode)
		
		// /media/pi/PiTB/TVShows/Loki/s1/Loki S01E01 Glorious Purpose.mp4

		// 710203042933435476 TVShows 1E 01  Glorious Purpose Loki 


	case strings.Contains(apath, "MastersOfTheUniverse"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "MastersOfTheUniverse"
		TvSI.Season = filename[36:38]
		TvSI.Episode = filename[39:41]
		TvSI.Title = filename[41:boo]
		TvSI.Series = "MastersOfTheUniverse"
		fmt.Printf("\n THIS IS TVI FROM MastersOfTheUniverse \n %s \n", TvSI)
		fmt.Println(TvSI.Season)
		fmt.Println(TvSI.Episode)











	case strings.Contains(apath, "TheBadBatch"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "TheBadBatch"
		TvSI.Season = filename[25:27]
		TvSI.Episode = filename[28:30]
		TvSI.Title = filename[30:boo]
		TvSI.Series = "TheBadBatch"
		fmt.Printf("\n THIS IS TVI FRO the bad batch \n %s \n", TvSI)

	case strings.Contains(apath, "SpaceTime"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "SpaceTime"
		TvSI.Season = "01"
		TvSI.Episode = "0"
		TvSI.Title = filename[11:boo]
		TvSI.Series = "SpaceTime"
	
	case strings.Contains(apath, "SeanCarrol"):
		_, filename := path.Split(apath)
		fspath := apath[20:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "SeanCarrol"
		TvSI.Season = "01"
		TvSI.Episode = "0"
		TvSI.Title = filename[:boo]
		TvSI.Series = "SeanCarrol"
	}
	return
}

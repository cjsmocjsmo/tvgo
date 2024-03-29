package tvgo

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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

func startLogging() string {
	logtxtfile := os.Getenv("MEDIACENTER_LOG_BASE_PATH") + "/tvgo_setup.log"
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logtxtfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	// log.Println("tvgo_setup logging started")
	return "tvgo_setup logging started"
}

func getTvShowInfo(apath string, tvshowpicPath string) (TvSI TVShowInfoS) {
	startLogging()
	switch {
	case strings.Contains(apath, "TVShows/TNG"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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

	case strings.Contains(apath, "Lost In Space"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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
		fspath := apath[21:]
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

	case strings.Contains(apath, "MastersOfTheUniverse"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
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

	case strings.Contains(apath, "TheBadBatch"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
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

	case strings.Contains(apath, "WhatIf"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "WhatIf"
		TvSI.Season = filename[9:11]
		TvSI.Episode = filename[12:14]
		TvSI.Title = filename[14:boo]
		TvSI.Series = "WhatIf"

	case strings.Contains(apath, "YTheLastMan"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "YTheLastMan"
		TvSI.Season = filename[16:18]
		TvSI.Episode = filename[19:21]
		TvSI.Title = filename[21:boo]
		TvSI.Series = "YTheLastMan"

	case strings.Contains(apath, "Foundation"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Foundation"
		TvSI.Season = filename[12:14]
		TvSI.Episode = filename[15:17]
		TvSI.Title = filename[17:boo]
		TvSI.Series = "Foundation"

	case strings.Contains(apath, "Visions"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Visions"
		TvSI.Season = filename[19:21]
		TvSI.Episode = filename[22:24]
		TvSI.Title = filename[24:boo]
		TvSI.Series = "Visions"

	case strings.Contains(apath, "Prodigy"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Prodigy"
		TvSI.Season = filename[19:21]
		TvSI.Episode = filename[22:24]
		TvSI.Title = filename[24:boo]
		TvSI.Series = "Prodigy"

	case strings.Contains(apath, "WheelOfTime"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "WheelOfTime"
		TvSI.Season = filename[19:21]
		TvSI.Episode = filename[22:24]
		TvSI.Title = filename[24:boo]
		TvSI.Series = "WheelOfTime"

	case strings.Contains(apath, "CowboyBebop"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "CowboyBebop"
		TvSI.Season = filename[14:16]
		TvSI.Episode = filename[17:19]
		TvSI.Title = filename[19:boo]
		TvSI.Series = "CowboyBebop"
		log.Println("Starting CowboyBebop")
		log.Println(filename[14:16])
		log.Println(filename[17:19])
		log.Println(filename[19:boo])

	case strings.Contains(apath, "Hawkeye"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Hawkeye"
		TvSI.Season = filename[9:11]
		TvSI.Episode = filename[13:15]
		TvSI.Title = filename[15:boo]
		TvSI.Series = "Hawkeye"
		log.Println("Starting Hawkeye")
		log.Println(filename[9:11])
		log.Println(filename[13:15])
		log.Println(filename[15:boo])

	case strings.Contains(apath, "BookOfBobaFett"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "BookOfBobaFett"
		TvSI.Season = filename[23:25]
		TvSI.Episode = filename[27:29]
		TvSI.Title = filename[29:boo]
		TvSI.Series = "BookOfBobaFett"
		log.Println("Starting BookOfBobaFett")

	case strings.Contains(apath, "Reacher"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Reacher"
		TvSI.Season = filename[9:11]
		TvSI.Episode = filename[12:14]
		TvSI.Title = filename[15:boo]
		TvSI.Series = "Reacher"
		log.Println("Starting Reacher")
		log.Println(filename[10:12])
		log.Println(filename[12:14])
		log.Println(filename[15:boo])

	case strings.Contains(apath, "Halo"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Halo"
		TvSI.Season = filename[6:8]
		TvSI.Episode = filename[9:11]
		TvSI.Title = filename[12:boo]
		TvSI.Series = "Halo"
		log.Println("Starting Halo")
		fmt.Println(filename[12:boo])

	case strings.Contains(apath, "MoonKnight"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "MoonKnight"
		TvSI.Season = filename[13:15]
		TvSI.Episode = filename[16:18]
		TvSI.Title = filename[19:boo]
		TvSI.Series = "MoonKnight"
		log.Println("Starting MoonKnight")
		fmt.Println(filename[19:boo])

	case strings.Contains(apath, "StrangeNewWorlds"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "StrangeNewWorlds"
		TvSI.Season = filename[30:32]
		TvSI.Episode = filename[33:35]
		TvSI.Title = filename[36:boo]
		TvSI.Series = "StrangeNewWorlds"
		log.Println("Starting StrangeNewWorlds")
		fmt.Println(filename[24:boo])

	case strings.Contains(apath, "PrehistoricPlanet"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "PrehistoricPlanet"
		TvSI.Season = filename[25:27]
		TvSI.Episode = filename[28:30]
		TvSI.Title = filename[30:boo]
		TvSI.Series = "PrehistoricPlanet"
		log.Println("Starting PrehistoricPlanet")
		fmt.Println(filename[24:boo])

	case strings.Contains(apath, "ObiWanKenobi"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "ObiWanKenobi"
		TvSI.Season = filename[16:18]
		TvSI.Episode = filename[19:21]
		TvSI.Title = filename[21:boo]
		TvSI.Series = "ObiWanKenobi"
		log.Println("Starting ObiWanKenobi")
		fmt.Println(filename[24:boo])

	case strings.Contains(apath, "MSMarvel"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "MSMarvel"
		TvSI.Season = filename[11:13]
		TvSI.Episode = filename[14:16]
		TvSI.Title = filename[16:boo]
		TvSI.Series = "MSMarvel"
		log.Println("Starting MSMarvel")
		fmt.Println(filename[16:boo])

	
	// /media/pi/PiTB/media/ TVShows/prehistoricplanet/s1/IAmGroot S01E01 Glorious Purpose.mp4
	case strings.Contains(apath, "IAmGroot"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "IAmGroot"
		TvSI.Season = filename[12:14]
		TvSI.Episode = filename[15:17]
		TvSI.Title = filename[17:boo]
		TvSI.Series = "IAmGroot"
		log.Println("Starting IAmGroot")
		fmt.Println(filename[17:boo])
	
	case strings.Contains(apath, "SheHulk"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "SheHulk"
		TvSI.Season = filename[25:27]
		TvSI.Episode = filename[28:30]
		TvSI.Title = filename[31:boo]
		TvSI.Series = "SheHulk"
		log.Println("Starting SheHulk")
		fmt.Println(filename[31:boo])


// /media/pi/PiTB/media/ TVShows/prehistoricplanet/s1/House Of The Dragon S01E01 Glorious Purpose.mp4
	case strings.Contains(apath, "HouseOfTheDragon"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "HouseOfTheDragon"
		TvSI.Season = filename[21:23]
		TvSI.Episode = filename[24:26]
		TvSI.Title = filename[26:boo]
		TvSI.Series = "HouseOfTheDragon"
		log.Println("Starting HouseOfTheDragon")
		fmt.Println(filename[26:boo])

// /media/pi/PiTB/media/ TVShows/prehistoricplanet/s1/The Lord Of The Rings The Rings Of Power S01E01 Glorious Purpose.mp4
	case strings.Contains(apath, "TheLordOfTheRingsTheRingsOfPower"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "TheLordOfTheRingsTheRingsOfPower"
		TvSI.Season = filename[42:44]
		TvSI.Episode = filename[45:47]
		TvSI.Title = filename[47:boo]
		TvSI.Series = "TheLordOfTheRingsTheRingsOfPower"
		log.Println("Starting TheLordOfTheRingsTheRingsOfPower")
		fmt.Println(filename[47:boo])

// /media/pi/PiTB/media/ TVShows/prehistoricplanet/s1/Andor S01E01 Glorious Purpose.mp4
	case strings.Contains(apath, "Andor"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "Andor"
		TvSI.Season = filename[7:9]
		TvSI.Episode = filename[10:12]
		TvSI.Title = filename[12:boo]
		TvSI.Series = "Andor"
		log.Println("Starting Andor")
		fmt.Println(filename[12:boo])

	// /media/pi/PiTB/media/ TVShows/prehistoricplanet/s1/Night Sky S01E01 Glorious Purpose.mp4
	case strings.Contains(apath, "NightSky"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "NightSky"
		TvSI.Season = filename[11:13]
		TvSI.Episode = filename[14:16]
		TvSI.Title = filename[17:boo]
		TvSI.Series = "NightSky"
		log.Println("Starting NightSky")
		fmt.Println(filename[17:boo])

	// /media/pi/PiTB/media/ TVShows/prehistoricplanet/s1/Star Wars Tales Of The Jedi S01E01 Glorious Purpose.mp4
	case strings.Contains(apath, "TalesOfTheJedi"):
		_, filename := path.Split(apath)
		fspath := apath[21:]
		boo := len(filename) - 4
		TvSI.ID = bson.NewObjectId()
		TvSI.FilePath = apath
		TvSI.MediaID = tvshowsUUID()
		TvSI.Genre = "TVShows"
		TvSI.TVShowPicPath = tvshowpicPath
		TvSI.TvFSPath = fspath
		TvSI.Catagory = "TalesOfTheJedi"
		TvSI.Season = filename[29:31]
		TvSI.Episode = filename[32:34]
		TvSI.Title = filename[34:boo]
		TvSI.Series = "TalesOfTheJedi"
		log.Println("Starting TalesOfTheJedi")
		fmt.Println(filename[34:boo])
	}

	return
}

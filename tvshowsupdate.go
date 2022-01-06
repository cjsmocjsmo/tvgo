package tvgo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/globalsign/mgo/bson"
)

// var finished2 bool = false

func scanFileNames() {
	err := filepath.Walk(os.Getenv("MEDIACENTER_TVSHOWS_PATH"), tvshowUpdateDirVisit)
	if err != nil {
		fmt.Println(err)
	}
}

func tvshowUpdateDirVisit(pAth string, f os.FileInfo, err error) error {
	log.Printf("this is path: %s", pAth)
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if f.IsDir() {
		return nil // not a file.  ignore.
	}

	ext := filepath.Ext(pAth)
	switch ext {
		case "":
			return nil
		case ".nfo":
			os.Remove(pAth)
			return nil
		case ".txt":
			os.Remove(pAth)
			return nil
		case ".mp4":
			if !tvshowNameInDb(pAth) {
				ProcessTVShowInfo(pAth)
			}
		case ".mkv":
			if !tvshowNameInDb(pAth) { 
				ProcessTVShowInfo(pAth)
			}
		case ".avi":
			if !tvshowNameInDb(pAth) {
				ProcessTVShowInfo(pAth)
			}
		case ".m4v":
			if !tvshowNameInDb(pAth) {
				ProcessTVShowInfo(pAth)
			}
	}
	return nil
}

func tvshowNameInDb(fn string) (result bool) {
	sess := TVDBcon()
	defer sess.Close()
	MTc := sess.DB("tvgobs").C("tvgobs")
	b1 := bson.M{"filepath": fn}
	b2 := bson.M{"_Id": 0}
	var PMedia []map[string]string
	err := MTc.Find(b1).Select(b2).All(&PMedia)
	if err != nil {
		log.Println(err)
	}
	num := len(PMedia)
	if num != 0 {
		result = true
	} else {
		result = false
	}
	return
}

// tvshowUpdate needs to be exported
func TvShowsUpdate() (finished2 bool) {
	scanFileNames()
	finished2 = true
	return
}
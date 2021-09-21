package tvgo

import (
	"fmt"
	"github.com/globalsign/mgo"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

//TVDBcon is exported for all our db connection objects
func TVDBcon() *mgo.Session {
	fmt.Println("Starting Update db session")
	s, err := mgo.Dial(os.Getenv("MEDIACENTER_MONGODB_ADDRESS"))
	if err != nil {
		fmt.Println("this is dial err")
		panic(err)
	}
	return s
}

func isDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()
	// read in ONLY one file
	_, err = f.Readdir(1)
	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func ProcessTVShowInfo(pAth string) {
	log.Println("\n\n Process_TVShows has started")
	tvpicPath := os.Getenv("MEDIACENTER_NO_ART_PIC_PATH")
	TvI := getTvShowInfo(pAth, tvpicPath)
	ses := TVDBcon()
	defer ses.Close()
	MTc := ses.DB("tvgobs").C("tvgobs")
	err := MTc.Insert(&TvI)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func myDirVisit(pAth string, f os.FileInfo, err error) error {
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
		return nil
	case ".txt":
		return nil
	case ".mp4":
		ProcessTVShowInfo(pAth)
	case ".mkv":
		ProcessTVShowInfo(pAth)
	case ".avi":
		ProcessTVShowInfo(pAth)
	case ".m4v":
		ProcessTVShowInfo(pAth)
	}
	return nil
}

var finished bool = false

// TVUpdate needs to be exported
func TVUpdate() (finished bool) {
	TvShowsUpdate()
	finished = true
	fmt.Println("Update Complete")
	return
}

func setupLogging() {
	logfile := os.Getenv("MEDIACENTER_LOG_BASE_PATH") + "moviegobsTVsetup.log"
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Logging started \n")
}


//TVSetUp is exported to main
func TVSetUp() (ExStat int) {
	setupLogging()
	//Start the timer
	startTime := time.Now().Unix()
	fmt.Printf("setup function has started at: %T", startTime)
	//Connect to the DB
	sess := TVDBcon()
	err := sess.DB("tvgobs").DropDatabase()
	if err != nil {
		fmt.Println(err)
	}
	err = filepath.Walk(os.Getenv("MEDIACENTER_TVSHOWS_PATH"), myDirVisit)
	if err != nil {
		fmt.Println(err)
	}

	os.Setenv("MEDIACENTER_TVGO_SETUP", "0")
	fmt.Println(startTime)
	stopTime := time.Now().Unix()
	fmt.Println(stopTime)
	etime := stopTime - startTime
	fmt.Println(etime)
	fmt.Println("SETUP IS COMPLETE")
	ExStat = 0
	return
}




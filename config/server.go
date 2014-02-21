package config

import (
	"database/sql"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/elgs/filesync/api"
	"github.com/elgs/filesync/index"
	"github.com/howeyc/fsnotify"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
)

func StartServer(configFile string) {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(configFile, " not found")
		return
	}
	json, _ := simplejson.NewJson(b)
	ip := json.Get("ip").MustString("127.0.0.1")
	port := json.Get("port").MustInt(6776)

	monitors := json.Get("monitors").MustMap()

	for _, v := range monitors {
		watcher, _ := fsnotify.NewWatcher()
		monitored, _ := v.(string)
		monitored = index.PathSafe(monitored)
		db, _ := sql.Open("sqlite3", index.SlashSuffix(monitored)+".sync/index.db")
		defer db.Close()
		db.Exec("VACUUM;")
		index.InitIndex(monitored, db)
		index.WatchRecursively(watcher, monitored, monitored)
		go index.ProcessEvent(watcher, monitored)
	}

	api.RunWeb(ip, port, monitors)
	//watcher.Close()
}

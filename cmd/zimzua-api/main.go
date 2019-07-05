package main

import (
	"fmt"
	"github.com/YonghoChoi/zimzua/cmd/zimzua-api/auth"
	"github.com/YonghoChoi/zimzua/cmd/zimzua-api/storage"
	"github.com/YonghoChoi/zimzua/pkg/version"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/zenazn/goji"
	"go.elastic.co/apm/module/apmgorilla"
	"log"
	"net/http"
	"os"
	"github.com/YonghoChoi/zimzua/cmd/zimzua-api/config"
)

var (
	store = sessions.NewCookieStore([]byte("secret???"))
)

func main() {
	if len(os.Args) < 3 || !(os.Args[1] == "-c" || os.Args[1] == "--config") {
		fmt.Println("Please input argument.")
		fmt.Println("-c, --config : config file path")
		fmt.Println("ex) ./zimzua-api --config ./config.yml")
		os.Exit(1)
	}

	if c := config.GetInstance(os.Args[2]); c == nil {
		fmt.Println("invalid config path")
		os.Exit(1)
	}

	os.Args = os.Args[:1]

	printVersion()
	vaildRequire(initRoute())
	startServe()
}

// 출저 : http://patorjk.com/software/taag/#p=display&f=Stop&t=zim%20zua%20api
func printVersion() {
	var logo string
	logo += "          _                                           _  \n"
	logo += "         (_)                                         (_) \n"
	logo += " 	 _____ _ ____     _____ _   _  ____     ____ ____  _  \n"
	logo += `	(___  ) |    \   (___  ) | | |/ _  |   / _  |  _ \| |` + "\n"
	logo += "	 / __/| | | | |   / __/| |_| ( ( | |  ( ( | | | | | | \n"
	logo += `	(_____)_|_|_|_|  (_____)\____|\_||_|   \_||_| ||_/|_|` + "\n"
	logo += "	                                            |_|      " + " v" + version.GetVersion() + "\n"

	fmt.Println(logo)
}

func vaildRequire(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func vaildOptional(err error) {
	if err != nil {
		log.Println(err)
	}
}

// NotFound is a 404 handler.
func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "If you have a problem, Please send to yongho1037@gmail.com", 404)
}

func initRoute() error {
	goji.Post("/regUser", auth.RegisterUser)
	goji.Post("/loginUser", auth.LoginUser)
	goji.Get("/getStorageList", storage.GetStorageList)

	// APM 사용
	if len(os.Getenv("ELASTIC_APM_SERVER_URL")) > 0 {
		router := mux.NewRouter()
		router.Use(apmgorilla.Middleware())
	}

	goji.NotFound(NotFound)

	return nil
}

func startServe() {
	log.Println("Running ZIM ZUA API...")
	goji.Serve()
}

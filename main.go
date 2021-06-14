package main
/*
#cgo CFLAGS: -I./iota/include -DIOTA_STREAMS_CHANNELS_CLIENT
//Choose one of the 2 below for compilation. Use .so for linux and .dylib for mac
#cgo LDFLAGS: ./iota/include/libiota_streams_c.so
//#cgo LDFLAGS: ./iota/include/libiota_streams_c.dylib
#include <channels.h>
*/
import "C"
import (
	"fmt"
	"github.com/project-alvarium/Alvarium-Server/api"
	"github.com/project-alvarium/Alvarium-Server/collections"
	"github.com/project-alvarium/Alvarium-Server/configuration"
	"github.com/project-alvarium/Alvarium-Server/subscriber"
	"log"
	"net/http"
	"time"
)

func main() {
	// drop subscribers
	subscriber.Shut()

	//VERY simple demonstration that the IOTA C bindings are included and callable
	C.drop_str(C.CString("A"))
	httpRouter := api.NewRouter()
	configuration.InitConfig()
	srv := &http.Server{
		Handler: httpRouter,
		Addr:    ":" + fmt.Sprint(configuration.Config.HTTPPort),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	collections.Database()

	// init the Annsubscriber 
	subscriber.Init()

	log.Fatal(srv.ListenAndServe())
	log.Println("listening")

}
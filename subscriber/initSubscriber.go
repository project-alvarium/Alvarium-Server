package subscriber
import (
	"database-manager/iota"
	"database-manager/configuration"
	"fmt"
	"os"
	"os/signal"
	"syscall"
);

// Global instance for ann subscriber and array of subscriber to drop on shutdown
var Subs []iota.Subscriber
var AnnSubscriber iota.Subscriber

func Init(){
	// configurations for the subscriber 
    subConfig := configuration.NewSubConfig()
	nodeConfig :=  configuration.NewNodeConfig()

	// Create a subscriber instance for annotator and await connection
	AnnSubscriber = iota.NewSubscriber(nodeConfig, subConfig)

	AnnSubscriber.AwaitKeyload()
	Subs = append(Subs, AnnSubscriber)

}
func Shut()  {
	setupShutdownHandler(&Subs)
}

//dropping on shutdown
func setupShutdownHandler(subs *[]iota.Subscriber) {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		fmt.Println("Shutdown called\nDropping Subscribers")
		for _, sub := range *subs {
			sub.Drop()
		}
		fmt.Println("Dropped\nExiting...")
		os.Exit(0)
	}()
}
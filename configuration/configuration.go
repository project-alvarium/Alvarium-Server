package configuration
import(
	"fmt"
	"math/rand"
	"time"
)
const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// [Streams Configs]
// Place announcement address here
const AnnAddress = "d1d2afc496ae15260d4296551b3dc09e1b029be3335d1b89990cb8b5d7446b7f0000000000000000:66fbee6b196ead77fe057e57"

// URL for author console
const AuthConsoleUrl = "http://127.0.0.1:8080"

// URL for IOTA node
const NodeUrl = "http://localhost:14265"

// Min Weight Magnitude
const NodeMwm = 5

// Max number of readings to conduct
const MaxReadings = 100




// Configuration holder
type Configuration struct {
	DatabaseName string
	DatabaseURL  string
	HTTPPort     string
}
type NodeConfig struct {
	Url string
	Mwm int8
}

type SubConfig struct {
	Seed       string
	Encoding   string
	AnnAddress string
}

// Config object
var (
	Config Configuration
)
func NewNodeConfig() NodeConfig {
	return NodeConfig{NodeUrl, NodeMwm}
}

func NewSubConfig() SubConfig {
	bytes := make([]byte, 64)
	rand.Seed(time.Now().UnixNano())
	for i := range bytes {
		bytes[i] = LetterBytes[rand.Intn(len(LetterBytes))]
	}

	seed := string(bytes)
	fmt.Println("Seed: ", seed)
	encoding := "utf-8"

	return SubConfig{seed, encoding, AnnAddress}
}
func InitConfig() {
	Config = Configuration{
		DatabaseName: "alvarium-db",
		DatabaseURL:  "mongodb://localhost:27017",
		HTTPPort:     "9090",
	}
	
}


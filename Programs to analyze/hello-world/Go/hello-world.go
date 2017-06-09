package main
import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
    fmt.Println("hello world")
    p.Stop()
}
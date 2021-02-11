package speedlog

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
)

func main() {
	user, _ := speedtest.FetchUserInfo()

	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)
		showServerResult(s)
	}
}

func showServerResult(server *speedtest.Server) {
	fmt.Printf(" \n")

	fmt.Printf("Download: %5.2f MB/s\n", server.DLSpeed/8)
	fmt.Printf("Upload: %5.2f MB/s\n\n", server.ULSpeed/8)
	valid := server.CheckResultValid()
	if !valid {
		fmt.Println("Warning: Result seems to be wrong. Please speedtest again.")
	}
}

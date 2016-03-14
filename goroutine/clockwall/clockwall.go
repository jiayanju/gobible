package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"

	"sort"
	"time"
)

func main() {
	cityServers := make(map[string]string)
	for _, v := range os.Args[1:] {
		values := strings.Split(v, "=")
		cityServers[values[0]] = values[1]
	}

	tc := time.Tick(1 * time.Second)

	for {
		select {
		case <-tc:
			tw := tabwriter.NewWriter(os.Stdout, 10, 4, 2, ' ', 0)

			const format = "%v\t%v\t\n"

			fmt.Fprintf(tw, format, "City", "Clock")
			fmt.Fprintf(tw, format, "-----", "-----")

			timeInfos := make(chan TimeInfo)
			for city, addr := range cityServers {
				go getTime(city, addr, timeInfos)
			}

			// fmt.Printf("map len: %d\n", len(cityServers))

			result := make(TimeInfoList, 0)
			for i := 0; i < len(cityServers); i++ {
				info := <-timeInfos
				// fmt.Printf("Rec Info: %s %s\n", info.city, info.time)
				result = append(result, info)
			}

			// fmt.Printf("result len : %d\n", len(result))

			sort.Sort(result)

			for _, info := range result {
				fmt.Fprintf(tw, format, info.city, info.time)
			}

			tw.Flush()
		}
	}
}

// TimeInfoList time info list
type TimeInfoList []TimeInfo

func (l TimeInfoList) Len() int {
	return len(l)
}

func (l TimeInfoList) Less(i, j int) bool {
	return l[i].city < l[j].city
}

func (l TimeInfoList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// TimeInfo include time info
type TimeInfo struct {
	city string
	time string
}

func getTime(city, addr string, infos chan TimeInfo) {
	// defer util.Trace("getTime")()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	time, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	info := TimeInfo{
		city: city,
		time: time,
	}
	// fmt.Printf("%s %s\n", info.city, info.time)
	infos <- info
}

func mustCopy(dest io.Writer, src io.Reader) {
	_, err := io.Copy(dest, src)
	if err != nil {
		log.Fatal(err)
	}
}

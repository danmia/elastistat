package main

import  (
	"fmt"
	"strconv"
	"time"
	"flag"
	"net/http"
	"encoding/json"
 	"github.com/danmia/stats"
)

func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}


func main()  {

	outticker := time.NewTicker(60 * time.Second)
	grabticker := time.NewTicker(10 * time.Second)
	docstats := stats.NewCounter()
	nodePtr := flag.String("node", "", "Elastic Node")
	var node string

	flag.Parse()

	if(*nodePtr != "")  {
        node = *nodePtr
    } else  {
		node = "localhost"
	}

	for  {
		select  {

			case <-grabticker.C:
				s := new(ClusterStats)
                getJson("http://" + node + ":9200/_cluster/stats", s)
				fmt.Println("Grabbing stats from: ",node, " Docs: ",  strconv.Itoa(s.Indices.Docs.Count))
                docstats.Set(int64(s.Indices.Docs.Count))

			case <-outticker.C:
				fmt.Println("Outputting stats")
				fmt.Printf("Doc count: %d, Docs/Sec: %.2f\n", docstats.Value(), docstats.Rate1m())
				
		}				
	}
}

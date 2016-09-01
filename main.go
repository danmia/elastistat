package main

import  (
	"fmt"
	// "strconv"
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
	var lastcpu int
	var fsused, fstot float64
	var node string
	var heapmax, heapused float64

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
				// fmt.Println("Grabbing stats from: ",node, " Docs: ",  strconv.Itoa(s.Indices.Docs.Count))
				lastcpu = s.Nodes.Process.CPU.Percent

				fsused = (float64(s.Nodes.Fs.TotalInBytes) - float64(s.Nodes.Fs.AvailableInBytes)) / 1024 / 1024 / 1024
				fstot = float64(s.Nodes.Fs.TotalInBytes) / 1024 / 1024 / 1024
				//											KB	   MB	  GB	
				heapmax = float64(s.Nodes.Jvm.Mem.HeapMaxInBytes) / 1024 / 1024 / 1024
				heapused = float64(s.Nodes.Jvm.Mem.HeapUsedInBytes) / 1024 / 1024 / 1024
		
                docstats.Set(int64(s.Indices.Docs.Count))

			case <-outticker.C:
				t := time.Now()
				fmt.Printf("Time: %s,  CPU: %d,  Heap Used/Total: %.2fGB/%.2fGB, Doc count: %d, Docs/Sec: %.2f Disk Size/Avail: %.2fGB/%.2fGB\n", t.Format("2006-01-02 15:04:05"), lastcpu, heapused, heapmax, docstats.Value(), docstats.Rate1m(), fsused, fstot)
				
		}				
	}
}

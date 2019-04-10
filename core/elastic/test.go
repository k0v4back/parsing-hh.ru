package elastic

import (
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"time"
)


func Test() {
	_, err := elastic.NewClient(
		elastic.SetURL("http://elasticsearch:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println(err)
		time.Sleep(3 * time.Second)
	} else {
		fmt.Println("YES")
	}
}

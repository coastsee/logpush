package engines

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Zinc struct {
	Index    string
	Url      string
	Username string
	Password string
	mux      sync.Mutex
}

// Flush only define Flush func
func (z *Zinc) Flush(pool []string) error {
	z.mux.Lock()
	defer z.mux.Unlock()

	var ndJson string
	if z.Index == "" {
		z.Index = "logpush"
	}
	for k, v := range pool {
		index := fmt.Sprintf(`{ "index" : { "_index" : "%s" } }`, z.Index) + "\n"
		if k > 0 {
			index = "\n" + index
		}
		ndJson = ndJson + index + v
	}
	err := z.bulkUploadDocument(ndJson)
	if err != nil {
		return err
	}
	return nil
}

func (z *Zinc) bulkUploadDocument(nJson string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", z.Url+"/api/_bulk", bytes.NewBuffer([]byte(nJson)))
	if err != nil {
		return fmt.Errorf("got error %s", err.Error())
	}
	req.SetBasicAuth(z.Username, z.Password)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("got error %s", err.Error())
	}
	defer response.Body.Close()

	//body, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))
	return nil
}

package requesthandler

import (
	"fmt"
	"time"

	"github.com/randyardiansyah25/wsbase-handler"
)

func BrowseNasabahHandler(client wsbase.WSClient, message wsbase.Message, body string) {
	fmt.Println("req >>", body)
	resp := message
	resp.Title = fmt.Sprint("Reply for : ", message.SenderId)
	resp.Body = "Ini body utk response browse nasabah"
	time.Sleep(10 * time.Second)
	client.SendMessage(resp)
}

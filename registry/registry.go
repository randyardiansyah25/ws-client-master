package registry

import (
	"ws-client-master/delivery/requesthandler"
	"ws-client-master/handler"
)

func RegisterHandler(handler handler.WSHandler) {
	handler.AddHandler("req_browse_nasabah", requesthandler.BrowseNasabahHandler)
	handler.AddHandler("req_cek_saldo", requesthandler.CekSaldoHandler)
}

package payment

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/payment/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	sessionDomain "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
)

type Handler struct {
	gateway Gateway
	logger  logging.Logger
}

func NewHandler(gateway Gateway, logger logging.Logger) Handler {
	return Handler{gateway: gateway, logger: logger}
}

func (h *Handler) Accept(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := r.ParseForm()
	if err != nil {
		h.logger.WithRequestID(r.Context()).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payment := domain.Payment{}
	_, err = fmt.Sscanf(r.PostForm.Get("AMOUNT"), "%d", &payment.Amount)
	if err != nil {
		h.logger.WithRequestID(r.Context()).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	payment.OrderID = r.Form.Get("MERCHANT_ORDER_ID")
	payment.Sign = r.Form.Get("SIGN")

	err = h.gateway.Accept(r.Context(), payment)
	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.logger.Trace("payment accepted")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("YES"))
}

func (h *Handler) GetPaymentLink(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	session, ok := r.Context().Value("session").(sessionDomain.Session)
	if !ok {
		h.logger.Trace("cant cast session from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	link, err := h.gateway.GetPaymentLink(r.Context(), session.UserID)
	if err != nil {
		h.logger.WithRequestID(r.Context()).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"link": link,
		},
	})

	if err != nil {
		h.logger.WithRequestID(r.Context()).Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

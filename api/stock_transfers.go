package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"oraluke.com/conn-ora1/domain"
)

// define struct param specificaly for all stock transfer

type CreateStockTransferRequest struct {
	CurrentWarehouseCode string `json:"current_warehouse_code"`
	ToWarehouseCode      string `json:"to_warehouse_code"`
}

// NOTE: handle passing context through all routes LATER
func (s *Server) CreateStockTransfer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Set("Content-Type", "application/json") // must be always first at handler

	var reqBody CreateStockTransferRequest

	defer r.Body.Close()
	errDecode := json.NewDecoder(r.Body).Decode(&reqBody)
	if errDecode != nil {
		s.writeError(w, http.StatusBadRequest, errDecode.Error())
		return
	}

	// NOTE: how to pass context to each service??
	// r.Context()

	errCreate := s.service.StockTransferSrv.Create(r.Context(), reqBody.CurrentWarehouseCode, reqBody.ToWarehouseCode)
	if errCreate != nil {
		s.writeError(w, http.StatusInternalServerError, errCreate.Error())
		return
	}

	s.writeJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "success create new stock transfer",
		Data:    reqBody, // TEST
	})

}

type StockTransferHeaderResponse struct {
	ID                 int        `json:"id"`
	StNumber           string     `json:"st_number"`
	TipeStID           int        `json:"tipe_st_id"`
	AccountOriginal    int        `json:"account_original"`
	AccountDestination int        `json:"account_destination"`
	CreatedBy          *int       `json:"created_by"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedBy          *int       `json:"updated_by"`
	UpdatedAt          *time.Time `json:"updated_at"`
	StatusApproval     string     `json:"status_approval"`
	StoCode            *string    `json:"sto_code"`
	Reference          *string    `json:"reference"`
	Quantity           *int       `json:"quantity"`
	Remarks            *string    `json:"remarks"`
	ApproveBy          *int       `json:"approve_by"`
	ApproveAt          *time.Time `json:"approve_at"`
	SubmittedBy        *int       `json:"submitted_by"`
	SubmittedAt        *time.Time `json:"submitted_at"`
	CancelledBy        *int       `json:"cancelled_by"`
	CancelledAt        *time.Time `json:"cancelled_at"`
}

// DTO
func toStockTransferHeaderResponse(source domain.StockTransferInfo) StockTransferHeaderResponse {
	return StockTransferHeaderResponse{
		ID:                 source.ID,
		StNumber:           source.StNumber,
		TipeStID:           source.TipeStID,
		AccountOriginal:    source.AccountOriginal,
		AccountDestination: source.AccountDestination,
		CreatedBy:          source.CreatedBy,
		CreatedAt:          source.CreatedAt,
		UpdatedBy:          source.UpdatedBy,
		UpdatedAt:          source.UpdatedAt,
		StatusApproval:     source.StatusApproval,
		StoCode:            source.StoCode,
		Reference:          source.Reference,
		Quantity:           source.Quantity,
		Remarks:            source.Remarks,
		ApproveBy:          source.ApproveBy,
		ApproveAt:          source.ApproveAt,
		SubmittedBy:        source.SubmittedBy,
		SubmittedAt:        source.SubmittedAt,
		CancelledBy:        source.CancelledBy,
		CancelledAt:        source.CancelledAt,
	}
}

func (s *Server) HeaderStockTransfer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json") // must be always first at handler
	reqNoTrans := "no_trans"

	// get Query param
	if !(r.URL.Query().Has(reqNoTrans)) {
		s.writeError(w, http.StatusBadRequest, "error query param")
		return
	}

	var responseHeader []StockTransferHeaderResponse

	noTrans := r.URL.Query().Get(reqNoTrans)
	infoHeader, errInfoHeader := s.service.StockTransferSrv.Header(r.Context(), noTrans)
	if errInfoHeader != nil {

		if errInfoHeader == sql.ErrNoRows {
			errMsg := fmt.Sprintf("nomor transaksi berikut ==> %s tidak ada di sistem", noTrans)
			s.writeJSON(w, http.StatusOK, Response{
				Success: false,
				Message: errMsg,
				Data:    infoHeader,
			})
			return
		} else {

			s.writeJSON(w, http.StatusInternalServerError, Response{
				Success: false,
				Message: errInfoHeader.Error(),
				Data:    infoHeader,
			})
			return
		}
	}

	for i := range infoHeader {
		responseHeader = append(responseHeader, toStockTransferHeaderResponse(infoHeader[i]))
	}

	s.writeJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "header berhasil didapatkan",
		Data:    responseHeader,
	})
}

type StockTransferDetailsResponse struct {
}

func (s *Server) DetailsStockTransfer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Set("Content-Type", "application/json") // must be always first at handler
	reqNoTrans := "no_trans"

	// get Query param
	if !(r.URL.Query().Has(reqNoTrans)) {
		s.writeError(w, http.StatusBadRequest, "error query param")
		return
	}

	// s.service.StockTransferSrv.Details()

}

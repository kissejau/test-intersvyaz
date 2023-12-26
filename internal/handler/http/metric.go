package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test-intersvyaz/internal/model"
	"test-intersvyaz/pkg/response"
)

func (h *Handler) Track(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Method != methodPost {
		response.NewErrorResponse(w, r, http.StatusMethodNotAllowed, fmt.Sprintf("Method %s not allowed", r.Method))
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		response.NewErrorResponse(w, r, http.StatusBadRequest, errInvalidBody)
		return
	}

	var metric model.Metric
	if err := json.Unmarshal(data, &metric); err != nil {
		response.NewErrorResponse(w, r, http.StatusBadRequest, errIncorrectBody)
		return
	}

	if err := h.amqpProducer.Track(ctx, metric); err != nil {
		response.NewErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(w, r, http.StatusAccepted, "metric was tracked")
}

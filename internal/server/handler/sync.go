package handler

import "net/http"

func (h *Handler) SyncHandler(w http.ResponseWriter, r *http.Request) {
	login, ok := r.Context().Value(LoginKey).(string)
	if !ok {
		http.Error(w, "no login in context", http.StatusInternalServerError)
		return
	}
	all, err := h.cardService.GetAll(r.Context(), login)
	if err != nil {
		return
	}
}

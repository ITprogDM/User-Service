package apperror

import (
	"errors"
	"net/http"
)

type ahandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h ahandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := h(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				}

				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(ErrNotFound.Marshal())
				return
			}

			w.WriteHeader(http.StatusTeapot)
			w.Write(systemError(err).Marshal())
		}
	}
}

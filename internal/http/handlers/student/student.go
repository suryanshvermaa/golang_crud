package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/suryanshvermaa/golang-crud/internal/types"
	"github.com/suryanshvermaa/golang-crud/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty error")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validateErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidatorError(validateErr))
			return
		}

		slog.Info("Creating a student")
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "true"})
	}
}

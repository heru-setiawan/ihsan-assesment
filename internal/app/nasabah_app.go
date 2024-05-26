package app

import (
	"assesment/internal/core/models"
	"assesment/pkg/exceptions"
	"context"
)

func (a *app) Daftar(ctx context.Context, data models.Nasabah) (*models.Nasabah, error) {
	errValidation := exceptions.Validation{
		Code: 400,
	}

	if data.NIK == "" {
		errValidation.Message = "nik tidak boleh kosong"
		a.log.Warn(map[string]any{"input": data, "error": errValidation}, "validation error")
		return nil, errValidation
	}

	if data.Nama == "" {
		errValidation.Message = "nama tidak boleh kosong"
		a.log.Warn(map[string]any{"input": data, "error": errValidation}, "validation error")

		return nil, errValidation
	}

	if data.NoHp == "" {
		errValidation.Message = "no hp tidak boleh kosong"
		a.log.Warn(map[string]any{"input": data, "error": errValidation}, "validation error")

		return nil, errValidation
	}

	data.NewRekening()

	result, err := a.nasabahDatabase.Daftar(ctx, data)
	if err != nil {
		a.log.Error(map[string]any{"input": data, "error": err}, "database error")
		return nil, err
	}

	return result, nil
}

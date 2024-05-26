package app

import (
	"assesment/internal/core/models"
	"assesment/pkg/exceptions"
	"context"
)

func (a *app) Tabung(ctx context.Context, noRekening string, nominal float64) (*models.Rekening, error) {
	errValidation := exceptions.Validation{
		Code: 400,
	}

	if noRekening == "" {
		errValidation.Message = "no rekening tidak boleh kosong"
		a.log.Warn(map[string]any{"input": struct {
			No      string
			Nominal float64
		}{
			No:      noRekening,
			Nominal: nominal,
		}, "error": errValidation}, "validation error")
		return nil, errValidation
	}

	if nominal < 0 {
		errValidation.Message = "nominal harus lebih dari 0"
		a.log.Warn(map[string]any{"input": struct {
			No      string
			Nominal float64
		}{
			No:      noRekening,
			Nominal: nominal,
		}, "error": errValidation}, "validation error")
		return nil, errValidation
	}

	result, err := a.rekeningDatabase.Transaksi(ctx, noRekening, nominal)
	if err != nil {
		a.log.Error(map[string]any{"input": struct {
			No      string
			Nominal float64
		}{
			No:      noRekening,
			Nominal: nominal,
		}, "error": err}, "database error")
		return nil, err
	}

	return result, nil
}

func (a *app) Tarik(ctx context.Context, noRekening string, nominal float64) (*models.Rekening, error) {
	errValidation := exceptions.Validation{
		Code: 400,
	}

	if noRekening == "" {
		errValidation.Message = "no rekening tidak boleh kosong"
		a.log.Warn(map[string]any{"input": struct {
			No      string
			Nominal float64
		}{
			No:      noRekening,
			Nominal: nominal,
		}, "error": errValidation}, "validation error")
		return nil, errValidation
	}

	if nominal < 0 {
		errValidation.Message = "nominal harus lebih dari 0"
		a.log.Warn(map[string]any{"input": struct {
			No      string
			Nominal float64
		}{
			No:      noRekening,
			Nominal: nominal,
		}, "error": errValidation}, "validation error")
		return nil, errValidation
	}

	nominal *= -1

	result, err := a.rekeningDatabase.Transaksi(ctx, noRekening, nominal)
	if err != nil {
		a.log.Error(map[string]any{"input": struct {
			No      string
			Nominal float64
		}{
			No:      noRekening,
			Nominal: nominal,
		}, "error": err}, "database error")
		return nil, err
	}

	return result, nil
}

func (a *app) CekSaldo(ctx context.Context, noRekening string) (*models.Rekening, error) {
	errValidation := exceptions.Validation{
		Code: 400,
	}

	if noRekening == "" {
		errValidation.Message = "no rekening tidak boleh kosong"
		a.log.Warn(map[string]any{"input": struct {
			No string
		}{
			No: noRekening,
		}, "error": errValidation}, "validation error")
		return nil, errValidation
	}

	result, err := a.rekeningDatabase.CekSaldo(ctx, noRekening)
	if err != nil {
		a.log.Error(map[string]any{"input": struct {
			No string
		}{
			No: noRekening,
		}, "error": err}, "database error")
		return nil, err
	}

	return result, nil
}

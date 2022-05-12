// Package handlers
//
//	________post_method_handler.go________
//
//	Calculate predict errors by the method passed in the body.
//	CRUD: POST
//	Path: /method/{id}
//	Body: { name : method }
//
//	Swagger handlers. Edit by alex.green
//
package handlers

import (
	agtm "CourseWork/algorithm"
	acr "CourseWork/algorithm/autocorrelation"
	cov "CourseWork/algorithm/covariance"
	"CourseWork/algorithm/ladder"
	"CourseWork/api/databases/mysql"
	"CourseWork/decoder/waveform"
	"CourseWork/gen/models"
	"CourseWork/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"strconv"
)

const (
	MAX_PREDICT = 16
	MIN_PREDICT = 2
)

func (handlers *Handlers) PostMethodHandler(log *log.Logger,
	mysql *mysql.Mysql) operations.PostMethodIDHandlerFunc {
	return func(params operations.PostMethodIDParams) middleware.Responder {
		var value float64
		predictError := make([]*models.EnergyItems0, MAX_PREDICT-MIN_PREDICT)
		decoder := new(waveform.WAVDecoder)

		desc, err := mysql.GetFileFromDB(log, params.ID)
		if err != nil {
			return errorEvent(err.Error(), log)
		}

		desc_decoder, err := decoder.DecodeFileFromDB(log, desc)
		if err != nil {
			return errorEvent(err.Error(), log)
		}

		if *params.Method.Method == agtm.AUTOCORRELATION_METHOD {
			for i := int32(MIN_PREDICT); i < MAX_PREDICT; i++ {
				value, err = acr.Autocorrelation(desc_decoder, i)
				if err != nil {
					log.Fatalln(err.Error())
				}
				predictError[i-MIN_PREDICT] = &models.EnergyItems0{
					strconv.FormatFloat(value, 'f', 10, 32),
				}
			}
		} else if *params.Method.Method == agtm.COVARIATION_METHOD {
			for i := int32(MIN_PREDICT); i < MAX_PREDICT; i++ {
				value, err = cov.Covariance(desc_decoder, i)
				if err != nil {
					log.Fatalln(err.Error())
				}
				predictError[i-MIN_PREDICT] = &models.EnergyItems0{
					strconv.FormatFloat(value, 'f', 10, 32),
				}
			}
		} else if *params.Method.Method == agtm.LADDER_METHOD {
			energy, err := ladder.Ladder(desc_decoder, MAX_PREDICT)
			if err != nil {
				log.Fatalln(err.Error())
			}
			for i := int32(MIN_PREDICT); i < MAX_PREDICT; i++ {
				predictError[i-MIN_PREDICT] = &models.EnergyItems0{
					strconv.FormatFloat(energy[i], 'f', 10, 32),
				}
			}
		} else {
			return errorEvent("The entered method was not found. Code error: -1", log)
		}

		return operations.NewPostMethodOK().WithPayload(&models.Energy{
			predictError,
		})
	}
}

func errorEvent(err string, log *log.Logger) *operations.PostMethodInternalServerError {
	log.Printf(err)
	return operations.NewPostMethodInternalServerError().WithPayload(&models.Error{
		Code:    nil,
		Message: &err,
	})
}

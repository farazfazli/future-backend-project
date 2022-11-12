package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	futuredb "github.com/farazfazli/future-backend-project/internal/db"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

func NewMux() *chi.Mux {
	r := chi.NewRouter()
	// Since this is a take-home development project, we are allowing wildcard CORS access
	r.Use(cors.Handler(cors.Options{}))
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.NoCache)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))
	// Limit the request body size
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, 1*1_000_000) // 1MB
			next.ServeHTTP(w, r)
		})
	})

	r.Get("/available-appointments", AvailableAppointments)
	r.Post("/appointments", AddAppointment)
	r.Get("/scheduled-appointments", ScheduledAppointments)
	return r
}

func AvailableAppointments(w http.ResponseWriter, r *http.Request) {
	trainerIdParam := r.URL.Query().Get("trainer_id")
	startsAtParam := r.URL.Query().Get("starts_at")
	endsAtParam := r.URL.Query().Get("ends_at")
	trainerId, err := uuid.Parse(trainerIdParam)
	if err != nil {
		log.Println("Trainer ID is not a valid UUID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	startsAt, err := time.Parse("2006-01-02T15:04:05-07:00", startsAtParam)
	if err != nil {
		log.Println("Starts at is not a valid date-time")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	endsAt, err := time.Parse("2006-01-02T15:04:05-07:00", endsAtParam)
	if err != nil {
		log.Println("Ends at is not a valid date-time")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	availableAppointmentsParams := futuredb.ListAvailableAppointmentsParams {
		TrainerID: trainerId,
		StartsAt: startsAt,
		EndsAt: endsAt,
	}
	res, err := futuredb.DBQueries.ListAvailableAppointments(context.Background(), availableAppointmentsParams)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	resJSON, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	_, err = w.Write(resJSON)
	if err != nil {
		log.Println(err)
	}
}

func AddAppointment(w http.ResponseWriter, r *http.Request) {
	var addAppointmentParams futuredb.PostNewAppointmentParams
	err := json.NewDecoder(r.Body).Decode(&addAppointmentParams)
		if err != nil {
		log.Println("Error decoding add appointment data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if addAppointmentParams.UserID == uuid.Nil || addAppointmentParams.TrainerID == uuid.Nil || 
	   addAppointmentParams.StartsAt.IsZero() || addAppointmentParams.EndsAt.IsZero() {
		log.Println("Invalid appointment data")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	searchAppointmentParams := futuredb.ListAvailableAppointmentsParams {
		TrainerID: addAppointmentParams.TrainerID,
		StartsAt: addAppointmentParams.StartsAt,
		EndsAt: addAppointmentParams.EndsAt,
	}
	res, err := futuredb.DBQueries.ListAvailableAppointments(context.Background(), searchAppointmentParams)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	diff := addAppointmentParams.EndsAt.Sub(addAppointmentParams.StartsAt)
	if diff.Minutes() != 30 {
		log.Println("Appointments must be 30 minutes long")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if addAppointmentParams.StartsAt.Minute() != 0 && addAppointmentParams.StartsAt.Minute() != 30 {
		log.Println("Appointments should be scheduled at :00 or :30 minutes after the hour")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	found := false
	for _, val := range res {
		if val.Equal(addAppointmentParams.StartsAt) {
			found = true
		}
	}
	if !found {
		log.Println("Appointment slot doesn't exist")
		w.WriteHeader(http.StatusConflict)
		return
	}
	
	newAppt, err := futuredb.DBQueries.PostNewAppointment(context.Background(), addAppointmentParams)
	newApptStruct := struct {
		Id int32 `json:"id"`
	}{newAppt}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	resJSON, err := json.Marshal(newApptStruct)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	_, err = w.Write(resJSON)
	if err != nil {
		log.Println(err)
	}
}

func ScheduledAppointments(w http.ResponseWriter, r *http.Request) {
	trainerIdParam := r.URL.Query().Get("trainer_id")
	trainerId, err := uuid.Parse(trainerIdParam)
	if err != nil {
		log.Println("Trainer ID is not a valid UUID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := futuredb.DBQueries.ListTrainerScheduledAppointments(context.Background(), trainerId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	resJSON, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	_, err = w.Write(resJSON)
	if err != nil {
		log.Println(err)
	}
}
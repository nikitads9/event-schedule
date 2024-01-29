package handlers

import (
	"event-schedule/internal/api"
	"event-schedule/internal/lib/logger/sl"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// GetVacantRooms godoc
//
//	@Summary		Get list of vacant rooms
//	@Description	Responds with event info with given EventID
//	@Tags			events
//	@Produce		json
//	@Param			user_id	path	int	true	"user_id"	Format(int64) default(1234)
//	@Param			start	path	string	true	"start"	Format(time.Time) default(2006-01-02T15:04:05Z07:00)
//	@Param			end	path	string	true	"end"	Format(time.Time) default(2006-01-02T15:04:05Z07:00)
//	@Success		200	{object}	api.GetVacantRoomsResponse
//	@Failure		400	{object}	api.GetVacantRoomsResponse
//	@Failure		404	{object}	api.GetVacantRoomsResponse
//	@Failure		422	{object}	api.GetVacantRoomsResponse
//	@Failure		503	{object}	api.GetVacantRoomsResponse
//	@Router			/events/{user_id}/get-vacant-rooms/{start}-{end} [get]
func (i *Implementation) GetVacantRooms(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.events.api.GetVacantRooms"
		var start, end string
		var err error

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		if start = chi.URLParam(r, "start"); start == "" {
			log.Error("invalid request", sl.Err(api.ErrNoInterval))
			render.Render(w, r, api.ErrInvalidRequest(api.ErrNoInterval))
			return
		}
		if end = chi.URLParam(r, "end"); end == "" {
			log.Error("invalid request", sl.Err(api.ErrNoInterval))
			render.Render(w, r, api.ErrInvalidRequest(api.ErrNoInterval))
			return
		}

		startDate, err := time.Parse("2006-01-02T15:04:05Z07:00", start)
		if err != nil {
			log.Error("invalid request", sl.Err(err))
			render.Render(w, r, api.ErrInvalidRequest(api.ErrInvalidDateFormat))
			return
		}
		endDate, err := time.Parse("2006-01-02T15:04:05Z07:00", end)
		if err != nil {
			log.Error("invalid request", sl.Err(err))
			render.Render(w, r, api.ErrInvalidRequest(api.ErrInvalidDateFormat))
			return
		}

		rooms, err := i.Service.GetVacantRooms(r.Context(), startDate, endDate) //TODO:GetVacantRooms
		if err != nil {
			log.Error("internal error", sl.Err(err))
			render.Render(w, r, api.ErrInternalError(err))
			return
		}

		log.Info("vacant room acquired", slog.Any("rooms", rooms))
		render.Status(r, http.StatusCreated)
		render.Render(w, r, api.GetVacantRoomsAPI(rooms))
	}
}

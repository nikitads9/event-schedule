package handlers

import (
	"booking-schedule/internal/app/api"
	"booking-schedule/internal/app/convert"
	"booking-schedule/internal/logger/sl"
	"booking-schedule/internal/middleware/auth"
	"log/slog"
	"time"

	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// GetBookings godoc
//
//	@Summary		Get several bookings info
//	@Description	Responds with series of booking info objects within given time period. The query parameters are start date and end date (start is to be before end and both should not be expired).
//	@ID				getMultipleBookingsByTag
//	@Tags			bookings
//	@Produce		json
//
//	@Param			start query		string	true	"start" Format(time.Time) default(2024-03-28T17:43:00)
//	@Param			end query		string	true	"end" Format(time.Time) default(2024-03-29T17:43:00)
//	@Success		200	{object}	api.GetBookingsResponse
//	@Failure		400	{object}	api.GetBookingsResponse
//	@Failure		401	{object}	api.GetBookingsResponse
//	@Failure		404	{object}	api.GetBookingsResponse
//	@Failure		422	{object}	api.GetBookingsResponse
//	@Failure		503	{object}	api.GetBookingsResponse
//	@Router			/get-bookings [get]
//
// @Security Bearer
func (i *Implementation) GetBookings(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "bookings.api.handlers.GetBookings"

		ctx := r.Context()

		log := logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(ctx)),
		)

		userID := auth.UserIDFromContext(ctx)
		if userID == 0 {
			log.Error("no user id in context", sl.Err(api.ErrNoUserID))
			err := render.Render(w, r, api.ErrUnauthorized(api.ErrNoAuth))
			if err != nil {
				log.Error("failed to render response", sl.Err(err))
				return
			}
			return
		}

		start := r.URL.Query().Get("start")
		if start == "" {
			log.Error("invalid request", sl.Err(api.ErrNoInterval))
			err := render.Render(w, r, api.ErrInvalidRequest(api.ErrNoInterval))
			if err != nil {
				log.Error("failed to render response", sl.Err(err))
				return
			}
			return
		}

		end := r.URL.Query().Get("end")
		if end == "" {
			log.Error("invalid request", sl.Err(api.ErrNoInterval))
			err := render.Render(w, r, api.ErrInvalidRequest(api.ErrNoInterval))
			if err != nil {
				log.Error("failed to render response", sl.Err(err))
				return
			}
			return
		}

		startDate, err := time.Parse("2006-01-02T15:04:05", start)
		if err != nil {
			log.Error("invalid request", sl.Err(err))
			err = render.Render(w, r, api.ErrInvalidRequest(api.ErrParse))
			if err != nil {
				log.Error("failed to render response", sl.Err(err))
				return
			}
			return
		}
		endDate, err := time.Parse("2006-01-02T15:04:05", end)
		if err != nil {
			log.Error("invalid request", sl.Err(err))
			err = render.Render(w, r, api.ErrInvalidRequest(api.ErrParse))
			if err != nil {
				log.Error("failed to render response", sl.Err(err))
				return
			}
			return
		}

		err = api.CheckDates(startDate, endDate)
		if err != nil {
			log.Error("invalid request", sl.Err(err))
			err = render.Render(w, r, api.ErrInvalidRequest(err))
			if err != nil {
				log.Error("failed to render response", sl.Err(err))
				return
			}
		}

		log.Info("received request", slog.Any("params:", start+" to "+end))

		bookings, err := i.Booking.GetBookings(ctx, startDate, endDate, userID)
		if err != nil {
			log.Error("internal error", sl.Err(err))
			render.Render(w, r, api.ErrInternalError(err))
			return
		}

		log.Info("bookings acquired", slog.Int("quantity: ", len(bookings)))

		render.Status(r, http.StatusCreated)
		err = render.Render(w, r, api.GetBookingsResponseAPI(convert.ToApiBookingsInfo(bookings)))
		if err != nil {
			log.Error("failed to render response", sl.Err(err))
			return
		}
	}

}

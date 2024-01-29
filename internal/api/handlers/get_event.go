package handlers

import (
	"event-schedule/internal/api"
	"event-schedule/internal/lib/logger/sl"
	"event-schedule/internal/model"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// GetEvent godoc
//
//	@Summary		Get event info
//	@Description	Responds with event info with given EventID
//	@Tags			events
//	@Produce		json
//	@Param			user_id	path	int	true	"user_id"	Format(int64) default(1234)
//	@Param			event_id	path	string	true	"event_id"	Format(uuid) default(550e8400-e29b-41d4-a716-446655440000)
//	@Success		200	{object}	api.GetEventResponse
//	@Failure		400	{object}	api.GetEventResponse
//	@Failure		404	{object}	api.GetEventResponse
//	@Failure		422	{object}	api.GetEventResponse
//	@Failure		503	{object}	api.GetEventResponse
//	@Router			/events/{user_id}/{event_id}/get [get]
func (i *Implementation) GetEvent(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.events.api.GetEvent"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		//TODO: проверить
		// Assume if we've reach this far, we can access the event
		// context because this handler is a child of the EventCtx
		// middleware. The worst case, the recoverer middleware will save us.
		event := r.Context().Value("event").(*model.EventInfo)

		if err := render.Render(w, r, api.GetEventResponseAPI(event)); err != nil {
			log.Error("internal error", sl.Err(err))
			render.Render(w, r, api.ErrRender(err))
			return
		}

		log.Info("event acquired", slog.Any("event", event))
	}
}

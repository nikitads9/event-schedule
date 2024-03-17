package user

import (
	"booking-schedule/internal/app/model"
	"booking-schedule/internal/logger/sl"
	"booking-schedule/internal/pkg/db"
	"context"
	"errors"
	"log/slog"

	t "booking-schedule/internal/app/repository/table"

	"github.com/go-chi/chi/middleware"
	"go.opentelemetry.io/otel/codes"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repository) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	const op = "users.repository.GetUser"

	log := r.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(ctx)),
	)
	ctx, span := r.tracer.Start(ctx, op)
	defer span.End()

	builder := sq.Select(t.ID, t.TelegramID, t.Name, t.TelegramNickname, t.CreatedAt, t.UpdatedAt).
		From(t.UserTable).
		Where(sq.Eq{t.ID: userID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("failed to build a query", sl.Err(err))
		return nil, ErrQueryBuild
	}

	span.AddEvent("query built")

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var res = new(model.User)
	err = r.client.DB().GetContext(ctx, res, q, args...)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		if errors.As(err, pgNoConnection) {
			log.Error("no connection to database host", sl.Err(err))
			return nil, ErrNoConnection
		}
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("booking with this id not found", sl.Err(err))
			return nil, ErrNotFound
		}
		log.Error("query execution error", sl.Err(err))
		return nil, ErrQuery
	}

	span.AddEvent("query successfully executed")

	return res, nil
}

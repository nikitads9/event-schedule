package user

import (
	"booking-schedule/internal/pkg/db"
	"context"
	"errors"
	"log/slog"

	t "booking-schedule/internal/app/repository/table"

	"github.com/go-chi/chi/middleware"
	"go.opentelemetry.io/otel/codes"

	sq "github.com/Masterminds/squirrel"
)

func (r *repository) DeleteUser(ctx context.Context, userID int64) error {
	const op = "users.repository.DeleteUser"

	log := r.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(ctx)),
	)
	ctx, span := r.tracer.Start(ctx, op)
	defer span.End()

	builder := sq.Delete(t.UserTable).
		Where(sq.Eq{t.ID: userID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("failed to build a query", err)
		return ErrQueryBuild
	}

	span.AddEvent("query built")

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	result, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		if errors.As(err, pgNoConnection) {
			log.Error("no connection to database host", err)
			return ErrNoConnection
		}
		log.Error("query execution error", err)
		return ErrQuery
	}

	if result.RowsAffected() == 0 {
		span.RecordError(ErrNoRowsAffected)
		span.SetStatus(codes.Error, ErrNoRowsAffected.Error())
		log.Error("unsuccessful delete", ErrNoRowsAffected)
		return ErrNotFound
	}

	span.AddEvent("query successfully executed")

	return nil

}

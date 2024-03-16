package booking

import (
	t "booking-schedule/internal/app/repository/table"
	"booking-schedule/internal/pkg/db"
	"context"
	"errors"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-chi/chi/middleware"
	"github.com/gofrs/uuid"
	"go.opentelemetry.io/otel/codes"
)

func (r *repository) DeleteBooking(ctx context.Context, bookingID uuid.UUID, userID int64) error {
	const op = "repository.booking.DeleteBooking"

	log := r.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(ctx)),
	)
	ctx, span := r.tracer.Start(ctx, op)
	defer span.End()

	builder := sq.Delete(t.BookingTable).
		Where(sq.And{
			sq.Eq{t.ID: bookingID},
			sq.Eq{t.UserID: userID},
		}).
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

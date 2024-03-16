package booking

import (
	"booking-schedule/internal/app/model"
	t "booking-schedule/internal/app/repository/table"
	"booking-schedule/internal/pkg/db"
	"context"
	"errors"
	"log/slog"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-chi/chi/middleware"
	"go.opentelemetry.io/otel/codes"
)

func (r *repository) UpdateBooking(ctx context.Context, mod *model.BookingInfo) error {
	const op = "repository.booking.UpdateBooking"

	log := r.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(ctx)),
	)
	ctx, span := r.tracer.Start(ctx, op)
	defer span.End()

	builder := sq.Update(t.BookingTable).
		Set(t.UpdatedAt, time.Now()).
		Set("start_date", mod.StartDate).
		Set("end_date", mod.EndDate).
		Set("suite_id", mod.SuiteID).
		Where(sq.And{
			sq.Eq{t.ID: mod.ID},
			sq.Eq{t.UserID: mod.UserID},
		}).
		PlaceholderFormat(sq.Dollar)

	if mod.NotifyAt != 0 {
		builder = builder.Set(t.NotifyAt, mod.NotifyAt)
	}

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
		log.Error("update unsuccessful", ErrNoRowsAffected)
		return ErrNotFound
	}

	span.AddEvent("query successfully executed")

	return nil
}

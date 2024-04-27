package mixin

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type condition func(context.Context, ent.Mutation) bool

type IdMixin struct {
	mixin.Schema
}

var TimeSchemaType = entgql.Type("DateTime")
var DefaultMixins = []ent.Mixin{IdMixin{}, MetadataMixin{}}

func (IdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique().
			Comment("Identifier").
			Annotations(
				entgql.OrderField("ID"),
				entgql.Type("ID"),
			),
	}
}

type MetadataMixin struct {
	mixin.Schema
}

func (MetadataMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp without time zone",
			}).
			Annotations(
				entgql.OrderField("CREATED_AT"),
				TimeSchemaType,
			),
		field.Time("updated_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp without time zone",
			}).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
				TimeSchemaType,
			),
	}
}

func hasOp(o ent.Op) condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(o)
	}
}

func (m MetadataMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		on(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, mut ent.Mutation) (ent.Value, error) {
					if s, ok := mut.(interface{ SetUpdatedAt(date time.Time) }); ok {
						s.SetUpdatedAt(time.Now())
					}

					v, err := next.Mutate(ctx, mut)
					if err != nil {
						return nil, fmt.Errorf("failed to run mutation hook: %w", err)
					}
					return v, nil
				})
			},
			ent.OpDelete|ent.OpUpdateOne,
		),
	}
}

func on(h ent.Hook, o ent.Op) ent.Hook {
	return iif(h, hasOp(o))
}

func iif(h ent.Hook, c condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if c(ctx, m) {
				res, err := h(next).Mutate(ctx, m)
				if err != nil {
					return nil, fmt.Errorf("failed to run mutation hook: %T: %w", h, err)
				}
				return res, nil
			}
			res, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, fmt.Errorf("failed to run mutation hook: %T: %w", h, err)
			}

			return res, nil
		})
	}
}

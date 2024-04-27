package mixin

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type IdMixin struct {
	mixin.Schema
}

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
				dialect.Postgres: "timestamp excluding timezone",
			}).
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Type("DateTime"),
			),
		field.Time("udated_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp excluding timezone",
			}).
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Type("DateTime"),
			),
	}
}

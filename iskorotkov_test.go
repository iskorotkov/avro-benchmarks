package avro_benchmarks

import (
	"testing"

	"github.com/iskorotkov/avro/v2"
	"github.com/stretchr/testify/assert"
)

func TestIskorotkovDecode(t *testing.T) {
	schema := avro.MustParse(Schema)

	superhero := Superhero{}
	err := avro.Unmarshal(schema, Payload, &superhero)

	want := Superhero{
		ID:            234765,
		AffiliationID: 9867,
		Name:          "Wolverine",
		Life:          85.25,
		Energy:        32.75,
		Powers: []*Superpower{
			{ID: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{ID: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{ID: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, want, superhero)
}

func BenchmarkIskorotkovDecode(b *testing.B) {
	schema := avro.MustParse(Schema)

	superhero := Superhero{}

	b.ReportAllocs()

	for b.Loop() {
		_ = avro.Unmarshal(schema, Payload, &superhero)
	}
}

func TestIskorotkovEncode(t *testing.T) {
	schema := avro.MustParse(Schema)

	superhero := Superhero{
		ID:            234765,
		AffiliationID: 9867,
		Name:          "Wolverine",
		Life:          85.25,
		Energy:        32.75,
		Powers: []*Superpower{
			{ID: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{ID: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{ID: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	b, err := avro.Marshal(schema, &superhero)

	assert.NoError(t, err)
	assert.Equal(t, Payload, b)
}

func BenchmarkIskorotkovEncode(b *testing.B) {
	schema := avro.MustParse(Schema)

	superhero := Superhero{
		ID:            234765,
		AffiliationID: 9867,
		Name:          "Wolverine",
		Life:          85.25,
		Energy:        32.75,
		Powers: []*Superpower{
			{ID: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{ID: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{ID: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	b.ReportAllocs()

	for b.Loop() {
		_, _ = avro.Marshal(schema, &superhero)
	}
}

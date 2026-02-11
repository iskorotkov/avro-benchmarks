package avro_benchmarks

import (
	"testing"

	"github.com/heetch/avro"
	"github.com/iskorotkov/avro-benchmarks/heetch"
	"github.com/stretchr/testify/assert"
)

func TestHeetchDecode(t *testing.T) {
	typ, err := avro.ParseType(Schema)
	if err != nil {
		t.Fatal(err)
	}

	superhero := heetch.Superhero{}
	_, err = avro.Unmarshal(Payload, &superhero, typ)

	want := heetch.Superhero{
		Id:             234765,
		Affiliation_id: 9867,
		Name:           "Wolverine",
		Life:           85.25,
		Energy:         32.75,
		Powers: []heetch.Superpower{
			{Id: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{Id: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{Id: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, want, superhero)
}

func BenchmarkHeetchDecode(b *testing.B) {
	typ, err := avro.ParseType(Schema)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for b.Loop() {
		superhero := heetch.Superhero{}
		_, _ = avro.Unmarshal(Payload, &superhero, typ)
	}
}

func TestHeetchEncode(t *testing.T) {
	superhero := heetch.Superhero{
		Id:             234765,
		Affiliation_id: 9867,
		Name:           "Wolverine",
		Life:           85.25,
		Energy:         32.75,
		Powers: []heetch.Superpower{
			{Id: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{Id: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{Id: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	b, _, err := avro.Marshal(superhero)

	assert.NoError(t, err)
	assert.Equal(t, PayloadSimple, b)
}

func BenchmarkHeetchEncode(b *testing.B) {
	superhero := heetch.Superhero{
		Id:             234765,
		Affiliation_id: 9867,
		Name:           "Wolverine",
		Life:           85.25,
		Energy:         32.75,
		Powers: []heetch.Superpower{
			{Id: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{Id: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{Id: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	b.ReportAllocs()

	for b.Loop() {
		_, _, _ = avro.Marshal(superhero)
	}
}

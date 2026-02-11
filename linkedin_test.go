package avro_benchmarks

import (
	"log"
	"testing"

	"github.com/linkedin/goavro/v2"
	"github.com/stretchr/testify/assert"
)

func TestLinkedinDecode(t *testing.T) {
	codec, err := goavro.NewCodec(Schema)
	if err != nil {
		log.Fatal(err)
	}

	superhero, _, err := codec.NativeFromBinary(Payload)

	want := map[string]any{
		"id":             int32(234765),
		"affiliation_id": int32(9867),
		"name":           "Wolverine",
		"life":           float32(85.25),
		"energy":         float32(32.75),
		"powers": []any{
			map[string]any{"id": int32(2345), "name": "Bone Claws", "damage": float32(5), "energy": float32(1.15), "passive": false},
			map[string]any{"id": int32(2346), "name": "Regeneration", "damage": float32(-2), "energy": float32(0.55), "passive": true},
			map[string]any{"id": int32(2347), "name": "Adamant skeleton", "damage": float32(-10), "energy": float32(0), "passive": true},
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, want, superhero)
}

func BenchmarkLinkedinDecode(b *testing.B) {
	codec, err := goavro.NewCodec(Schema)
	if err != nil {
		log.Fatal(err)
	}

	b.ReportAllocs()

	for b.Loop() {
		_, _, _ = codec.NativeFromBinary(Payload)
	}
}

func TestLinkedinEncode(t *testing.T) {
	codec, err := goavro.NewCodec(Schema)
	if err != nil {
		log.Fatal(err)
	}

	superhero := map[string]any{
		"id":             int32(234765),
		"affiliation_id": int32(9867),
		"name":           "Wolverine",
		"life":           float32(85.25),
		"energy":         float32(32.75),
		"powers": []any{
			map[string]any{"id": int32(2345), "name": "Bone Claws", "damage": float32(5), "energy": float32(1.15), "passive": false},
			map[string]any{"id": int32(2346), "name": "Regeneration", "damage": float32(-2), "energy": float32(0.55), "passive": true},
			map[string]any{"id": int32(2347), "name": "Adamant skeleton", "damage": float32(-10), "energy": float32(0), "passive": true},
		},
	}

	b, err := codec.BinaryFromNative(nil, superhero)

	assert.NoError(t, err)
	assert.Equal(t, PayloadSimple, b)
}

func BenchmarkLinkedinEncode(b *testing.B) {
	codec, err := goavro.NewCodec(Schema)
	if err != nil {
		log.Fatal(err)
	}

	superhero := map[string]any{
		"id":             int32(234765),
		"affiliation_id": int32(9867),
		"name":           "Wolverine",
		"life":           float32(85.25),
		"energy":         float32(32.75),
		"powers": []any{
			map[string]any{"id": int32(2345), "name": "Bone Claws", "damage": float32(5), "energy": float32(1.15), "passive": false},
			map[string]any{"id": int32(2346), "name": "Regeneration", "damage": float32(-2), "energy": float32(0.55), "passive": true},
			map[string]any{"id": int32(2347), "name": "Adamant skeleton", "damage": float32(-10), "energy": float32(0), "passive": true},
		},
	}

	b.ReportAllocs()

	for b.Loop() {
		_, _ = codec.BinaryFromNative(nil, superhero)
	}
}

package block

import (
	"phoenixbuilder/dragonfly/server/block/cube"
	"phoenixbuilder/dragonfly/server/item"
	"phoenixbuilder/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// Basalt is a type of igneous rock found in the Nether.
type Basalt struct {
	solid
	bassDrum

	// Polished specifies if the basalt is its polished variant.
	Polished bool
	// Axis is the axis which the basalt faces.
	Axis cube.Axis
}

// UseOnBlock ...
func (b Basalt) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(w, pos, face, b)
	if !used {
		return
	}
	b.Axis = face.Axis()

	place(w, pos, b, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (b Basalt) BreakInfo() BreakInfo {
	return newBreakInfo(1.25, pickaxeHarvestable, pickaxeEffective, oneOf(b))
}

// EncodeItem ...
func (b Basalt) EncodeItem() (name string, meta int16) {
	if b.Polished {
		return "minecraft:polished_basalt", 0
	}
	return "minecraft:basalt", 0
}

// EncodeBlock ...
func (b Basalt) EncodeBlock() (name string, properties map[string]interface{}) {
	if b.Polished {
		return "minecraft:polished_basalt", map[string]interface{}{"pillar_axis": b.Axis.String()}
	}
	return "minecraft:basalt", map[string]interface{}{"pillar_axis": b.Axis.String()}
}

// allBasalt ...
func allBasalt() (basalt []world.Block) {
	for _, axis := range cube.Axes() {
		basalt = append(basalt, Basalt{Axis: axis, Polished: false})
		basalt = append(basalt, Basalt{Axis: axis, Polished: true})
	}
	return
}

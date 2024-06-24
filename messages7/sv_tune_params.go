package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvTuneParams struct {
	ChunkHeader *chunk7.ChunkHeader

	GroundControlSpeed float32
	GroundControlAccel float32
	GroundFriction     float32
	GroundJumpImpulse  float32
	AirJumpImpulse     float32
	AirControlSpeed    float32
	AirControlAccel    float32
	AirFriction        float32
	HookLength         float32
	HookFireSpeed      float32
	HookDragAccel      float32
	HookDragSpeed      float32
	Gravity            float32
	VelrampStart       float32
	VelrampRange       float32
	VelrampCurvature   float32
	GunCurvature       float32
	GunSpeed           float32
	GunLifetime        float32
	ShotgunCurvature   float32
	ShotgunSpeed       float32
	ShotgunSpeeddiff   float32
	ShotgunLifetime    float32
	GrenadeCurvature   float32
	GrenadeSpeed       float32
	GrenadeLifetime    float32
	LaserReach         float32
	LaserBounceDelay   float32
	LaserBounceNum     float32
	LaserBounceCost    float32
	PlayerCollision    float32
	PlayerHooking      float32
}

func (msg *SvTuneParams) MsgId() int {
	return network7.MsgGameSvTuneParams
}

func (msg *SvTuneParams) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvTuneParams) System() bool {
	return false
}

func (msg *SvTuneParams) Vital() bool {
	return true
}

func (msg *SvTuneParams) Pack() []byte {
	return slices.Concat(
		packer.PackInt(int(msg.GroundControlSpeed*100)),
		packer.PackInt(int(msg.GroundControlAccel*100)),
		packer.PackInt(int(msg.GroundFriction*100)),
		packer.PackInt(int(msg.GroundJumpImpulse*100)),
		packer.PackInt(int(msg.AirJumpImpulse*100)),
		packer.PackInt(int(msg.AirControlSpeed*100)),
		packer.PackInt(int(msg.AirControlAccel*100)),
		packer.PackInt(int(msg.AirFriction*100)),
		packer.PackInt(int(msg.HookLength*100)),
		packer.PackInt(int(msg.HookFireSpeed*100)),
		packer.PackInt(int(msg.HookDragAccel*100)),
		packer.PackInt(int(msg.HookDragSpeed*100)),
		packer.PackInt(int(msg.Gravity*100)),
		packer.PackInt(int(msg.VelrampStart*100)),
		packer.PackInt(int(msg.VelrampRange*100)),
		packer.PackInt(int(msg.VelrampCurvature*100)),
		packer.PackInt(int(msg.GunCurvature*100)),
		packer.PackInt(int(msg.GunSpeed*100)),
		packer.PackInt(int(msg.GunLifetime*100)),
		packer.PackInt(int(msg.ShotgunCurvature*100)),
		packer.PackInt(int(msg.ShotgunSpeed*100)),
		packer.PackInt(int(msg.ShotgunSpeeddiff*100)),
		packer.PackInt(int(msg.ShotgunLifetime*100)),
		packer.PackInt(int(msg.GrenadeCurvature*100)),
		packer.PackInt(int(msg.GrenadeSpeed*100)),
		packer.PackInt(int(msg.GrenadeLifetime*100)),
		packer.PackInt(int(msg.LaserReach*100)),
		packer.PackInt(int(msg.LaserBounceDelay*100)),
		packer.PackInt(int(msg.LaserBounceNum*100)),
		packer.PackInt(int(msg.LaserBounceCost*100)),
		packer.PackInt(int(msg.PlayerCollision*100)),
		packer.PackInt(int(msg.PlayerHooking*100)),
	)
}

func (msg *SvTuneParams) Unpack(u *packer.Unpacker) error {
	msg.GroundControlSpeed = float32(u.GetInt()) / 100
	msg.GroundControlAccel = float32(u.GetInt()) / 100
	msg.GroundFriction = float32(u.GetInt()) / 100
	msg.GroundJumpImpulse = float32(u.GetInt()) / 100
	msg.AirJumpImpulse = float32(u.GetInt()) / 100
	msg.AirControlSpeed = float32(u.GetInt()) / 100
	msg.AirControlAccel = float32(u.GetInt()) / 100
	msg.AirFriction = float32(u.GetInt()) / 100
	msg.HookLength = float32(u.GetInt()) / 100
	msg.HookFireSpeed = float32(u.GetInt()) / 100
	msg.HookDragAccel = float32(u.GetInt()) / 100
	msg.HookDragSpeed = float32(u.GetInt()) / 100
	msg.Gravity = float32(u.GetInt()) / 100
	msg.VelrampStart = float32(u.GetInt()) / 100
	msg.VelrampRange = float32(u.GetInt()) / 100
	msg.VelrampCurvature = float32(u.GetInt()) / 100
	msg.GunCurvature = float32(u.GetInt()) / 100
	msg.GunSpeed = float32(u.GetInt()) / 100
	msg.GunLifetime = float32(u.GetInt()) / 100
	msg.ShotgunCurvature = float32(u.GetInt()) / 100
	msg.ShotgunSpeed = float32(u.GetInt()) / 100
	msg.ShotgunSpeeddiff = float32(u.GetInt()) / 100
	msg.ShotgunLifetime = float32(u.GetInt()) / 100
	msg.GrenadeCurvature = float32(u.GetInt()) / 100
	msg.GrenadeSpeed = float32(u.GetInt()) / 100
	msg.GrenadeLifetime = float32(u.GetInt()) / 100
	msg.LaserReach = float32(u.GetInt()) / 100
	msg.LaserBounceDelay = float32(u.GetInt()) / 100
	msg.LaserBounceNum = float32(u.GetInt()) / 100
	msg.LaserBounceCost = float32(u.GetInt()) / 100
	msg.PlayerCollision = float32(u.GetInt()) / 100
	msg.PlayerHooking = float32(u.GetInt()) / 100

	return nil
}

func (msg *SvTuneParams) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvTuneParams) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

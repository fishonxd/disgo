package discord

import (
	"github.com/disgoorg/omit"
	"github.com/disgoorg/snowflake/v2"
)

type ChannelUpdate interface {
	channelUpdate()
}

type GuildChannelUpdate interface {
	ChannelUpdate
	guildChannelUpdate()
}

type GuildTextChannelUpdate struct {
	Name                          *string                           `json:"name,omitempty"`
	Type                          *ChannelType                      `json:"type,omitempty"`
	Position                      omit.Omit[*int]                   `json:"position,omitzero"`
	Topic                         omit.Omit[*string]                `json:"topic,omitzero"`
	NSFW                          omit.Omit[*bool]                  `json:"nsfw,omitzero"`
	RateLimitPerUser              omit.Omit[*int]                   `json:"rate_limit_per_user,omitzero"`
	PermissionOverwrites          omit.Omit[*[]PermissionOverwrite] `json:"permission_overwrites,omitzero"`
	ParentID                      omit.Omit[*snowflake.ID]          `json:"parent_id,omitzero"`
	DefaultAutoArchiveDuration    omit.Omit[*AutoArchiveDuration]   `json:"default_auto_archive_duration,omitzero"`
	DefaultThreadRateLimitPerUser *int                              `json:"default_thread_rate_limit_per_user,omitempty"`
	Flags                         *ChannelFlags                     `json:"flags,omitempty"`
}

func (GuildTextChannelUpdate) channelUpdate()      {}
func (GuildTextChannelUpdate) guildChannelUpdate() {}

type GuildVoiceChannelUpdate struct {
	Name                 *string                           `json:"name,omitempty"`
	Position             omit.Omit[*int]                   `json:"position,omitzero"`
	RateLimitPerUser     omit.Omit[*int]                   `json:"rate_limit_per_user,omitzero"`
	Bitrate              omit.Omit[*int]                   `json:"bitrate,omitzero"`
	UserLimit            omit.Omit[*int]                   `json:"user_limit,omitzero"`
	PermissionOverwrites omit.Omit[*[]PermissionOverwrite] `json:"permission_overwrites,omitzero"`
	ParentID             omit.Omit[*snowflake.ID]          `json:"parent_id,omitzero"`
	RTCRegion            omit.Omit[*string]                `json:"rtc_region,omitzero"`
	NSFW                 omit.Omit[*bool]                  `json:"nsfw,omitzero"`
	VideoQualityMode     omit.Omit[*VideoQualityMode]      `json:"video_quality_mode,omitzero"`
	Flags                *ChannelFlags                     `json:"flags,omitempty"`
}

func (GuildVoiceChannelUpdate) channelUpdate()      {}
func (GuildVoiceChannelUpdate) guildChannelUpdate() {}

type GuildCategoryChannelUpdate struct {
	Name                 *string                           `json:"name,omitempty"`
	Position             omit.Omit[*int]                   `json:"position,omitzero"`
	PermissionOverwrites omit.Omit[*[]PermissionOverwrite] `json:"permission_overwrites,omitzero"`
}

func (GuildCategoryChannelUpdate) channelUpdate()      {}
func (GuildCategoryChannelUpdate) guildChannelUpdate() {}

type GuildNewsChannelUpdate struct {
	Name                       *string                           `json:"name,omitempty"`
	Type                       *ChannelType                      `json:"type,omitempty"`
	Position                   omit.Omit[*int]                   `json:"position,omitzero"`
	Topic                      omit.Omit[*string]                `json:"topic,omitzero"`
	RateLimitPerUser           *int                              `json:"rate_limit_per_user,omitempty"`
	PermissionOverwrites       omit.Omit[*[]PermissionOverwrite] `json:"permission_overwrites,omitzero"`
	ParentID                   omit.Omit[*snowflake.ID]          `json:"parent_id,omitzero"`
	DefaultAutoArchiveDuration omit.Omit[*int]                   `json:"default_auto_archive_duration,omitzero"`
	Flags                      *ChannelFlags                     `json:"flags,omitempty"`
}

func (GuildNewsChannelUpdate) channelUpdate()      {}
func (GuildNewsChannelUpdate) guildChannelUpdate() {}

type GuildThreadUpdate struct {
	Name                *string              `json:"name,omitempty"`
	Archived            *bool                `json:"archived,omitempty"`
	AutoArchiveDuration *AutoArchiveDuration `json:"auto_archive_duration,omitempty"`
	Locked              *bool                `json:"locked,omitempty"`
	Invitable           *bool                `json:"invitable,omitempty"`
	RateLimitPerUser    omit.Omit[*int]      `json:"rate_limit_per_user,omitzero"`
	Flags               *ChannelFlags        `json:"flags,omitempty"`
	AppliedTags         *[]snowflake.ID      `json:"applied_tags,omitempty"`
}

func (GuildThreadUpdate) channelUpdate()      {}
func (GuildThreadUpdate) guildChannelUpdate() {}

type GuildStageVoiceChannelUpdate struct {
	Name                 *string                           `json:"name,omitempty"`
	Position             omit.Omit[*int]                   `json:"position,omitzero"`
	RateLimitPerUser     omit.Omit[*int]                   `json:"rate_limit_per_user,omitzero"`
	Bitrate              omit.Omit[*int]                   `json:"bitrate,omitzero"`
	UserLimit            omit.Omit[*int]                   `json:"user_limit,omitzero"`
	PermissionOverwrites omit.Omit[*[]PermissionOverwrite] `json:"permission_overwrites,omitzero"`
	ParentID             omit.Omit[*snowflake.ID]          `json:"parent_id,omitzero"`
	RTCRegion            omit.Omit[*string]                `json:"rtc_region,omitzero"`
	NSFW                 omit.Omit[*bool]                  `json:"nsfw,omitzero"`
	VideoQualityMode     omit.Omit[*VideoQualityMode]      `json:"video_quality_mode,omitzero"`
}

func (GuildStageVoiceChannelUpdate) channelUpdate()      {}
func (GuildStageVoiceChannelUpdate) guildChannelUpdate() {}

type GuildForumChannelUpdate struct {
	Name                          *string                           `json:"name,omitempty"`
	Position                      omit.Omit[*int]                   `json:"position,omitzero"`
	Topic                         omit.Omit[*string]                `json:"topic,omitzero"`
	NSFW                          omit.Omit[*bool]                  `json:"nsfw,omitzero"`
	PermissionOverwrites          omit.Omit[*[]PermissionOverwrite] `json:"permission_overwrites,omitzero"`
	ParentID                      omit.Omit[*snowflake.ID]          `json:"parent_id,omitzero"`
	RateLimitPerUser              omit.Omit[*int]                   `json:"rate_limit_per_user,omitzero"`
	AvailableTags                 *[]ChannelTag                     `json:"available_tags,omitempty"`
	Flags                         *ChannelFlags                     `json:"flags,omitempty"`
	DefaultReactionEmoji          omit.Omit[*DefaultReactionEmoji]  `json:"default_reaction_emoji,omitzero"`
	DefaultThreadRateLimitPerUser *int                              `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              omit.Omit[*DefaultSortOrder]      `json:"default_sort_order,omitzero"`
	DefaultForumLayout            omit.Omit[*DefaultForumLayout]    `json:"default_forum_layout,omitzero"`
}

func (GuildForumChannelUpdate) channelUpdate()      {}
func (GuildForumChannelUpdate) guildChannelUpdate() {}

type GuildMediaChannelUpdate struct {
	Name                          *string                           `json:"name,omitempty"`
	Position                      omit.Omit[*int]                   `json:"position,omitzero"`
	Topic                         omit.Omit[*string]                `json:"topic,omitzero"`
	NSFW                          omit.Omit[*bool]                  `json:"nsfw,omitzero"`
	PermissionOverwrites          omit.Omit[*[]PermissionOverwrite] `json:"permission_overwrites,omitzero"`
	ParentID                      omit.Omit[*snowflake.ID]          `json:"parent_id,omitzero"`
	RateLimitPerUser              omit.Omit[*int]                   `json:"rate_limit_per_user,omitzero"`
	AvailableTags                 *[]ChannelTag                     `json:"available_tags,omitempty"`
	Flags                         *ChannelFlags                     `json:"flags,omitempty"`
	DefaultReactionEmoji          omit.Omit[*DefaultReactionEmoji]  `json:"default_reaction_emoji,omitzero"`
	DefaultThreadRateLimitPerUser *int                              `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              omit.Omit[*DefaultSortOrder]      `json:"default_sort_order,omitzero"`
}

func (GuildMediaChannelUpdate) channelUpdate()      {}
func (GuildMediaChannelUpdate) guildChannelUpdate() {}

type GuildPostUpdate struct {
	Name                *string              `json:"name,omitempty"`
	Archived            *bool                `json:"archived,omitempty"`
	AutoArchiveDuration *AutoArchiveDuration `json:"auto_archive_duration,omitempty"`
	Locked              *bool                `json:"locked,omitempty"`
	Invitable           *bool                `json:"invitable,omitempty"`
	RateLimitPerUser    omit.Omit[*int]      `json:"rate_limit_per_user,omitzero"`
	Flags               *ChannelFlags        `json:"flags,omitempty"`
	AppliedTags         *[]snowflake.ID      `json:"applied_tags,omitempty"`
}

func (GuildPostUpdate) channelUpdate()      {}
func (GuildPostUpdate) guildChannelUpdate() {}

type GuildChannelPositionUpdate struct {
	ID              snowflake.ID             `json:"id"`
	Position        omit.Omit[*int]          `json:"position,omitzero"`
	LockPermissions omit.Omit[*bool]         `json:"lock_permissions,omitzero"`
	ParentID        omit.Omit[*snowflake.ID] `json:"parent_id,omitzero"`
}

// VoiceChannelStatusUpdate is used with the SetVoiceChannelStatus endpoint (PUT /channels/{channel.id}/voice-status).
// Status is a short string (up to 500 characters) describing what's happening in the voice channel.
// Pass an empty string to clear the status.
type VoiceChannelStatusUpdate struct {
	Status string `json:"status"`
}

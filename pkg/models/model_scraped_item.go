package models

import (
	"database/sql"
)

type ScrapedItem struct {
	ID              int             `db:"id" json:"id"`
	Title           sql.NullString  `db:"title" json:"title"`
	Description     sql.NullString  `db:"description" json:"description"`
	URL             sql.NullString  `db:"url" json:"url"`
	Date            SQLiteDate      `db:"date" json:"date"`
	Rating          sql.NullString  `db:"rating" json:"rating"`
	Tags            sql.NullString  `db:"tags" json:"tags"`
	Models          sql.NullString  `db:"models" json:"models"`
	Episode         sql.NullInt64   `db:"episode" json:"episode"`
	GalleryFilename sql.NullString  `db:"gallery_filename" json:"gallery_filename"`
	GalleryURL      sql.NullString  `db:"gallery_url" json:"gallery_url"`
	VideoFilename   sql.NullString  `db:"video_filename" json:"video_filename"`
	VideoURL        sql.NullString  `db:"video_url" json:"video_url"`
	StudioID        sql.NullInt64   `db:"studio_id,omitempty" json:"studio_id"`
	CreatedAt       SQLiteTimestamp `db:"created_at" json:"created_at"`
	UpdatedAt       SQLiteTimestamp `db:"updated_at" json:"updated_at"`
}

type ScrapedPerformer struct {
	Name         *string `graphql:"name" json:"name"`
	URL          *string `graphql:"url" json:"url"`
	PhotoUrl     *string `graphql:"photo_url" json:"photo_url"`
	Twitter      *string `graphql:"twitter" json:"twitter"`
	Instagram    *string `graphql:"instagram" json:"instagram"`
	Birthdate    *string `graphql:"birthdate" json:"birthdate"`
	Ethnicity    *string `graphql:"ethnicity" json:"ethnicity"`
	Country      *string `graphql:"country" json:"country"`
	EyeColor     *string `graphql:"eye_color" json:"eye_color"`
	Height       *string `graphql:"height" json:"height"`
	Measurements *string `graphql:"measurements" json:"measurements"`
	FakeTits     *string `graphql:"fake_tits" json:"fake_tits"`
	CareerLength *string `graphql:"career_length" json:"career_length"`
	Tattoos      *string `graphql:"tattoos" json:"tattoos"`
	Piercings    *string `graphql:"piercings" json:"piercings"`
	Aliases      *string `graphql:"aliases" json:"aliases"`
}

type ScrapedScene struct {
	Title      *string                  `graphql:"title" json:"title"`
	Details    *string                  `graphql:"details" json:"details"`
	URL        *string                  `graphql:"url" json:"url"`
	Date       *string                  `graphql:"date" json:"date"`
	File       *SceneFileType           `graphql:"file" json:"file"`
	Studio     *ScrapedSceneStudio      `graphql:"studio" json:"studio"`
	Tags       []*ScrapedSceneTag       `graphql:"tags" json:"tags"`
	Performers []*ScrapedScenePerformer `graphql:"performers" json:"performers"`
}

type ScrapedScenePerformer struct {
	// Set if performer matched
	ID           *string `graphql:"id" json:"id"`
	Name         string  `graphql:"name" json:"name"`
	URL          *string `graphql:"url" json:"url"`
	PhotoUrl     *string `graphql:"photo_url" json:"photo_url"`
	Twitter      *string `graphql:"twitter" json:"twitter"`
	Instagram    *string `graphql:"instagram" json:"instagram"`
	Birthdate    *string `graphql:"birthdate" json:"birthdate"`
	Ethnicity    *string `graphql:"ethnicity" json:"ethnicity"`
	Country      *string `graphql:"country" json:"country"`
	EyeColor     *string `graphql:"eye_color" json:"eye_color"`
	Height       *string `graphql:"height" json:"height"`
	Measurements *string `graphql:"measurements" json:"measurements"`
	FakeTits     *string `graphql:"fake_tits" json:"fake_tits"`
	CareerLength *string `graphql:"career_length" json:"career_length"`
	Tattoos      *string `graphql:"tattoos" json:"tattoos"`
	Piercings    *string `graphql:"piercings" json:"piercings"`
	Aliases      *string `graphql:"aliases" json:"aliases"`
}

type ScrapedSceneStudio struct {
	// Set if studio matched
	ID   *string `graphql:"id" json:"id"`
	Name string  `graphql:"name" json:"name"`
	URL  *string `graphql:"url" json:"url"`
}

type ScrapedSceneTag struct {
	// Set if tag matched
	ID   *string `graphql:"id" json:"id"`
	Name string  `graphql:"name" json:"name"`
}

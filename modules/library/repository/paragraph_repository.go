package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/internal/database"
	"pirate-lang-go/modules/library/entity"
)

func (r *LibraryRepository) CreateParagraph(ctx context.Context, paragraph *entity.Paragraph) error {
	var (
		Title         sql.NullString
		ParagraphType sql.NullString
		AudioUrl      sql.NullString
		ImageUrl      sql.NullString
	)

	if paragraph.Title != "" {
		Title = sql.NullString{String: paragraph.Title, Valid: true}
	}
	if paragraph.ParagraphType != "" {
		ParagraphType = sql.NullString{String: paragraph.ParagraphType, Valid: true}
	}
	if paragraph.AudioUrl != "" {
		AudioUrl = sql.NullString{String: paragraph.AudioUrl, Valid: true}
	}
	if paragraph.ImageUrl != "" {
		ImageUrl = sql.NullString{String: paragraph.ImageUrl, Valid: true}
	}

	_, err := r.Queries.CreateParagraph(ctx, database.CreateParagraphParams{
		ParagraphContent: paragraph.ParagraphContent,
		Title:            Title,
		PartID:           paragraph.PartID,
		ParagraphOrder:   paragraph.ParagraphOrder,
		ParagraphType:    ParagraphType,
		AudioUrl:         AudioUrl,
		ImageUrl:         ImageUrl,
	})
	if err != nil {
		logger.Error("LibraryRepository.CreateParagraph: failed to create paragraph", "error", err)
		return err
	}
	return nil
}

func (r *LibraryRepository) UpdateParagraph(ctx context.Context, paragraph *entity.Paragraph, paragraphId uuid.UUID) error {
	var (
		Title         sql.NullString
		ParagraphType sql.NullString
	)

	if paragraph.Title != "" {
		Title = sql.NullString{String: paragraph.Title, Valid: true}
	}
	if paragraph.ParagraphType != "" {
		ParagraphType = sql.NullString{String: paragraph.ParagraphType, Valid: true}
	}

	err := r.Queries.UpdateParagraph(ctx, database.UpdateParagraphParams{
		ParagraphContent: paragraph.ParagraphContent,
		Title:            Title,
		PartID:           paragraph.PartID,
		ParagraphOrder:   paragraph.ParagraphOrder,
		ParagraphType:    ParagraphType,
		ParagraphID:      paragraphId,
	})
	if err != nil {
		logger.Error("LibraryRepository.UpdateParagraph: failed to update paragraph",
			"paragraph_id", paragraphId, "error", err)
		return err
	}
	return nil
}

func (r *LibraryRepository) GetParagraph(ctx context.Context, paragraphId uuid.UUID) (*entity.Paragraph, error) {
	dbParagraph, err := r.Queries.GetParagraphByID(ctx, paragraphId)
	if err != nil {
		logger.Error("LibraryRepository.GetParagraph: failed to retrieve paragraph",
			"paragraph_id", paragraphId, "error", err)
		return nil, err
	}

	getString := func(ns sql.NullString) string {
		if ns.Valid {
			return ns.String
		}
		return ""
	}

	return &entity.Paragraph{
		ParagraphID:      dbParagraph.ParagraphID,
		ParagraphContent: dbParagraph.ParagraphContent,
		Title:            getString(dbParagraph.Title),
		PartID:           dbParagraph.PartID,
		ParagraphOrder:   dbParagraph.ParagraphOrder,
		ParagraphType:    getString(dbParagraph.ParagraphType),
		AudioUrl:         getString(dbParagraph.AudioUrl),
		ImageUrl:         getString(dbParagraph.ImageUrl),
		CreatedAt:        dbParagraph.CreatedAt.Time,
		UpdatedAt:        dbParagraph.UpdatedAt.Time,
	}, nil
}

func (r *LibraryRepository) GetParagraphsByPartId(ctx context.Context, partId uuid.UUID) ([]*entity.Paragraph, error) {

	dbParagraphs, err := r.Queries.GetParagraphByPartId(ctx, partId)
	if err != nil {
		logger.Error("LibraryRepository.GetParagraphs: failed to retrieve paginated paragraphs",
			"error", err)
		return nil, err
	}

	var paragraphs []*entity.Paragraph

	for _, dbParagraph := range dbParagraphs {
		paragraph := &entity.Paragraph{
			ParagraphID:      dbParagraph.ParagraphID,
			ParagraphContent: dbParagraph.ParagraphContent,
			Title:            dbParagraph.Title.String,
			PartID:           dbParagraph.PartID,
			ParagraphOrder:   dbParagraph.ParagraphOrder,
			ParagraphType:    dbParagraph.ParagraphType.String,
			AudioUrl:         dbParagraph.AudioUrl.String,
			ImageUrl:         dbParagraph.ImageUrl.String,
			CreatedAt:        dbParagraph.CreatedAt.Time,
			UpdatedAt:        dbParagraph.UpdatedAt.Time,
		}
		paragraphs = append(paragraphs, paragraph)
	}

	return paragraphs, nil
}
func (r *LibraryRepository) UpdateAudioParagraph(ctx context.Context, audioUrl *string, paragraphId uuid.UUID) error {
	var (
		Url sql.NullString
	)
	if audioUrl != nil {
		Url = sql.NullString{String: *audioUrl, Valid: true}
	}
	err := r.Queries.UpdateParagraphAudioURL(ctx, database.UpdateParagraphAudioURLParams{
		AudioUrl:    Url,
		ParagraphID: paragraphId,
	})
	if err != nil {
		logger.Error("LibraryRepository:UpdateAudioParagraph: failed to update audio content for group",
			"group_id", paragraphId,
			"audio_url_attempted", audioUrl,
			"error", err)
		return err
	}
	return nil
}

func (r *LibraryRepository) UpdateImageParagraph(ctx context.Context, imageUrl *string, paragraphId uuid.UUID) error {
	var (
		Url sql.NullString
	)
	if imageUrl != nil {
		Url = sql.NullString{String: *imageUrl, Valid: true}
	}
	err := r.Queries.UpdateParagraphImageURL(ctx, database.UpdateParagraphImageURLParams{
		ImageUrl:    Url,
		ParagraphID: paragraphId,
	})
	if err != nil {
		logger.Error("LibraryRepository.UpdateImageGroup: failed to update image content for group",
			"group_id", paragraphId,
			"image_url_attempted", imageUrl,
			"error", err)
		return err
	}
	return nil
}

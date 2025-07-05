package repository

//
//func (r *LibraryRepository) CreateGroupGroup(ctx context.Context, group *entity.QuestionGroup) (*uuid.UUID, error) {
//	var (
//		Name      string
//		Desc      sql.NullString
//		PartId    uuid.UUID
//		PlanType  string
//		GroupType string
//	)
//	Name = group.Name
//	if group.Description != "" {
//		Desc = sql.NullString{String: group.Description, Valid: true}
//	}
//	PlanType = group.PlanType
//	GroupType = group.GroupType
//	PartId = group.PartID
//	params := database.CreateQuestionGroupParams{
//		Name:        Name,
//		Description: Desc,
//		PartID:      PartId,
//		PlanType:    PlanType,
//		GroupType:   GroupType,
//	}
//	groupId, err := r.Queries.CreateQuestionGroup(ctx, params)
//	if err != nil {
//		logger.Error("LibraryRepository.CreateGroupGroup: failed to create question group",
//			"group_name", Name,
//			"part_id", PartId,
//			"error", err)
//		return nil, err
//	}
//	return &groupId, err
//}
//
//func (r *LibraryRepository) UpdateAudioGroup(ctx context.Context, audioUrl *string, groupId uuid.UUID) error {
//	var (
//		Url sql.NullString
//	)
//	if audioUrl != nil {
//		Url = sql.NullString{String: *audioUrl, Valid: true}
//	}
//	err := r.Queries.UpdateAudioContentQuestionGroup(ctx, database.UpdateAudioContentQuestionGroupParams{
//		ContextAudioUrl: Url,
//		QuestionGroupID: groupId,
//	})
//	if err != nil {
//		logger.Error("LibraryRepository.UpdateAudioGroup: failed to update audio content for group",
//			"group_id", groupId,
//			"audio_url_attempted", audioUrl,
//			"error", err)
//		return err
//	}
//	return nil
//}
//
//func (r *LibraryRepository) UpdateImageGroup(ctx context.Context, imageUrl *string, groupId uuid.UUID) error {
//	var (
//		Url sql.NullString
//	)
//	if imageUrl != nil {
//		Url = sql.NullString{String: *imageUrl, Valid: true}
//	}
//	err := r.Queries.UpdateImageContentQuestionGroup(ctx, database.UpdateImageContentQuestionGroupParams{
//		ContextImageUrl: Url,
//		QuestionGroupID: groupId,
//	})
//	if err != nil {
//		logger.Error("LibraryRepository.UpdateImageGroup: failed to update image content for group",
//			"group_id", groupId,
//			"image_url_attempted", imageUrl,
//			"error", err)
//		return err
//	}
//	return nil
//}
//
//func (r *LibraryRepository) GetAudioGroup(ctx context.Context, groupId uuid.UUID) (string, error) {
//	objectName, err := r.Queries.GetAudioUrlGroup(ctx, groupId)
//	if err != nil {
//		logger.Error("LibraryRepository.GetAudioGroup: failed to retrieve audio URL for group",
//			"group_id", groupId,
//			"error", err)
//		return "", err
//	}
//	return objectName.String, nil
//}
//
//func (r *LibraryRepository) GroupExists(ctx context.Context, groupId uuid.UUID) (bool, error) {
//	exists, err := r.Queries.QuestionGroupExists(ctx, groupId)
//	if err != nil {
//		logger.Error("LibraryRepository.GroupExists: failed to check existence of group",
//			"group_id", groupId,
//			"error", err)
//		return false, err
//	}
//	return exists, nil
//}
//
//func (r *LibraryRepository) UpdateQuestionGroup(ctx context.Context, group *entity.QuestionGroup, groupId uuid.UUID) error {
//	var (
//		Name        string
//		Desc        sql.NullString
//		PartId      uuid.UUID
//		PlanType    string
//		GroupType   string
//		TextContent sql.NullString
//	)
//	Name = group.Name
//	if group.Description != "" {
//		Desc = sql.NullString{String: group.Description, Valid: true}
//	}
//	PlanType = group.PlanType
//	GroupType = group.GroupType
//	PartId = group.PartID
//	if group.ContextTextContent != "" {
//		TextContent = sql.NullString{String: group.ContextTextContent, Valid: true}
//	}
//	params := database.UpdateQuestionGroupParams{
//		Name:               Name,
//		Description:        Desc,
//		PartID:             PartId,
//		PlanType:           PlanType,
//		GroupType:          GroupType,
//		ContextTextContent: TextContent,
//		QuestionGroupID:    groupId,
//	}
//	err := r.Queries.UpdateQuestionGroup(ctx, params)
//	if err != nil {
//		logger.Error("LibraryRepository.UpdateQuestionGroup: failed to update question group details",
//			"group_id", groupId,
//			"group_name_attempted", Name,
//			"error", err)
//		return err
//	}
//	return nil
//}
//
//func (r *LibraryRepository) GetQuestionGroups(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedQuestionGroup, error) {
//	// Get total count
//	totalItems, err := r.Queries.GetPartsCount(ctx)
//	if err != nil {
//		logger.Error("LibraryRepository.GetQuestionGroups: failed to get total count of groups",
//			"page_number", pageNumber,
//			"page_size", pageSize,
//			"error", err)
//		return nil, err
//	}
//
//	offset := (pageNumber - 1) * pageSize
//	// Get paginated Parts
//	listParams := database.GetPaginatedQuestionGroupsParams{
//		Limit:  int32(pageSize),
//		Offset: int32(offset),
//	}
//
//	dbParts, err := r.Queries.GetPaginatedQuestionGroups(ctx, listParams)
//	if err != nil {
//		logger.Error("LibraryRepository.GetQuestionGroups: failed to retrieve paginated question groups",
//			"page_number", pageNumber,
//			"page_size", pageSize,
//			"offset", offset,
//			"error", err)
//		return nil, err
//	}
//	var groups []*entity.QuestionGroup
//	for _, dbPart := range dbParts {
//		part := &entity.QuestionGroup{
//			QuestionGroupID: dbPart.QuestionGroupID,
//			Name:            dbPart.Name,
//			Description:     dbPart.Description.String,
//			PartID:          dbPart.PartID,
//			PlanType:        dbPart.PlanType,
//			GroupType:       dbPart.GroupType,
//		}
//		groups = append(groups, part)
//	}
//	totalPages := (totalItems + int64(pageSize) - 1) / int64(pageSize)
//
//	return &entity.PaginatedQuestionGroup{
//		Items:       groups,
//		TotalItems:  totalItems,
//		TotalPages:  totalPages,
//		CurrentPage: pageNumber,
//		PageSize:    pageSize,
//	}, nil
//}
//func (r *LibraryRepository) GetQuestionsByGroups(ctx context.Context, groupId uuid.UUID) ([]*entity.Question, error) {
//	questionDBs, err := r.Queries.GetGroupQuestionList(ctx, uuid.NullUUID{UUID: groupId})
//	if err != nil {
//		logger.Error("LibraryRepository:UpdateQuestionGroup: failed to get questions from group",
//			"group_id", groupId,
//			"error", err)
//		return nil, err
//	}
//	var questions []*entity.Question
//	for _, questionDB := range questionDBs {
//		question := &entity.Question{
//			QuestionID:           questionDB.QuestionID,
//			QuestionOrder:        questionDB.QuestionOrder,
//			TextContent:          questionDB.TextContent.String,
//			AudioUrl:             questionDB.AudioUrl.String,
//			ImageUrl:             questionDB.ImageUrl.String,
//			ExpectedAnswerFormat: questionDB.ExpectedAnswerFormat.String,
//		}
//		questions = append(questions, question)
//	}
//	return questions, nil
//}
//func (r *LibraryRepository) CreateQuestion(ctx context.Context, questionRequest *entity.Question, groupId uuid.UUID) (*entity.Question, error) {
//
//	var (
//		QuestionGroupID      uuid.NullUUID
//		QuestionOrder        int32
//		TextContent          sql.NullString
//		AudioUrl             sql.NullString
//		ImageUrl             sql.NullString
//		ExpectedAnswerFormat sql.NullString
//	)
//	QuestionGroupID = uuid.NullUUID{UUID: groupId}
//	QuestionOrder = questionRequest.QuestionOrder
//	TextContent = sql.NullString{String: questionRequest.TextContent, Valid: true}
//	AudioUrl = sql.NullString{String: questionRequest.AudioUrl, Valid: true}
//	questionDB, err := r.Queries.CreateQuestion(ctx, database.CreateQuestionParams{
//		QuestionGroupID:      QuestionGroupID,
//		QuestionOrder:        QuestionOrder,
//		TextContent:          TextContent,
//		AudioUrl:             AudioUrl,
//		ImageUrl:             ImageUrl,
//		ExpectedAnswerFormat: ExpectedAnswerFormat,
//	})
//	if err != nil {
//		logger.Error("LibraryRepository:CreateQuestion: failed to create question",
//			"group_id", groupId,
//			"error", err)
//		return nil, err
//	}
//
//	question := &entity.Question{
//		QuestionID:           questionDB.QuestionID,
//		QuestionOrder:        questionDB.QuestionOrder,
//		TextContent:          questionDB.TextContent.String,
//		AudioUrl:             questionDB.AudioUrl.String,
//		ImageUrl:             questionDB.ImageUrl.String,
//		ExpectedAnswerFormat: questionDB.ExpectedAnswerFormat.String,
//	}
//	return question, nil
//}
//
//func (r *LibraryRepository) UpdateQuestion(ctx context.Context, questionRequest *entity.Question, questionId uuid.UUID) error {
//
//	var (
//		QuestionOrder        int32
//		TextContent          sql.NullString
//		ExpectedAnswerFormat sql.NullString
//	)
//	QuestionOrder = questionRequest.QuestionOrder
//	TextContent = sql.NullString{String: questionRequest.TextContent, Valid: true}
//	err := r.Queries.UpdateQuestion(ctx, database.UpdateQuestionParams{
//		QuestionOrder:        QuestionOrder,
//		TextContent:          TextContent,
//		ExpectedAnswerFormat: ExpectedAnswerFormat,
//		QuestionID:           questionId,
//	})
//	if err != nil {
//		logger.Error("LibraryRepository:CreateQuestion: failed to create question")
//		return err
//	}
//	return nil
//}

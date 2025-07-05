package service

//
//func (s *LibraryService) CreateQuestionGroup(ctx context.Context, req *dto.CreateQuestionGroupRequest) (string, *errors.AppError) {
//	groupEntity := mapper.ToQuestionGroupEntity(req)
//	groupId, err := s.repo.CreateGroupGroup(ctx, groupEntity)
//	if err != nil {
//		logger.Error("LibraryService:CreateQuestionGroup:Failed to create question group", "error", err)
//		return "", errors.NewAppError(errors.ErrInternal, "Service:CreateQuestionGroup:Failed to create question group", err)
//	}
//	return groupId.String(), nil
//}
//func (s *LibraryService) UpdateQuestionGroup(ctx context.Context, groupId uuid.UUID, req *dto.UpdateQuestionGroupRequest) *errors.AppError {
//
//	groupEntity := mapper.ToQuestionGroupEntityForUpdate(req)
//	err := s.repo.UpdateQuestionGroup(ctx, groupEntity, groupId)
//	if err != nil {
//		logger.Error("LibraryService:UpdateQuestionGroup:Failed to update question group", "error", err)
//		return errors.NewAppError(errors.ErrInternal, "Service:UpdateQuestionGroup:Failed to update question group", err)
//	}
//	return nil
//}
//func (s *LibraryService) UploadAudioGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError) {
//	src, err := file.Open()
//	if err != nil {
//		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
//		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to read audio file", err)
//	}
//	defer src.Close()
//	objectName, objectURL, err := s.storage.UploadAudio(ctx, groupId, src, file.Size, file.Filename, GroupFolder)
//	if err != nil {
//		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
//		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to upload audio file", err)
//	}
//	err = s.repo.UpdateAudioGroup(ctx, &objectURL, groupId)
//	if err != nil {
//		logger.Error("LibraryService:UploadAudioGroup:Failed to update audio URL in database", "error", err, "groupId", groupId.String())
//		return nil, errors.NewAppError(errors.ErrInternal, "Service:UploadAudioGroup:Failed to persist audio information in database", err)
//	}
//	response := &dto.UpdateContentFileResponse{
//		Filename:  objectName,
//		ObjectURL: objectURL,
//	}
//	return response, nil
//}
//func (s *LibraryService) UploadTranscriptAudioGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID, language string) (*dto.UpdateContentFileResponse, *errors.AppError) {
//	src, err := file.Open()
//	if err != nil {
//		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
//		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to read audio file", err)
//	}
//	defer src.Close()
//	objectName, objectURL, err := s.storage.UploadTranscriptAudio(ctx, groupId, src, file.Size, file.Filename, TranscriptFolder, language)
//	if err != nil {
//		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
//		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to upload audio file", err)
//	}
//	err = s.repo.UpdateAudioGroup(ctx, &objectName, groupId)
//	if err != nil {
//		logger.Error("LibraryService:UploadAudioGroup:Failed to update audio URL in database", "error", err, "groupId", groupId.String())
//		return nil, errors.NewAppError(errors.ErrInternal, "Service:UploadAudioGroup:Failed to persist audio information in database", err)
//	}
//	response := &dto.UpdateContentFileResponse{
//		Filename:  objectName,
//		ObjectURL: objectURL,
//	}
//	return response, nil
//}
//func (s *LibraryService) UploadImageGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError) {
//	src, err := file.Open()
//	if err != nil {
//		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
//		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to read audio file", err)
//	}
//	defer src.Close()
//	objectName, objectURL, err := s.storage.UploadImage(ctx, groupId, src, file.Size, file.Filename, ImageGroupFolder)
//	if err != nil {
//		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
//		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to upload audio file", err)
//	}
//	err = s.repo.UpdateAudioGroup(ctx, &objectName, groupId)
//	if err != nil {
//		logger.Error("LibraryService:UploadAudioGroup:Failed to update audio URL in database", "error", err, "groupId", groupId.String())
//		return nil, errors.NewAppError(errors.ErrInternal, "Service:UploadAudioGroup:Failed to persist audio information in database", err)
//	}
//	response := &dto.UpdateContentFileResponse{
//		Filename:  objectName,
//		ObjectURL: objectURL,
//	}
//	return response, nil
//}
//func (s *LibraryService) DeleteAudioGroup(ctx context.Context, groupId uuid.UUID) *errors.AppError {
//	objectName := ""
//	err := s.repo.UpdateAudioGroup(ctx, &objectName, groupId)
//	if err != nil {
//		logger.Error("LibraryService:UploadAudioGroup:Failed to update audio URL in database", "error", err, "groupId", groupId.String())
//		return errors.NewAppError(errors.ErrInternal, "LibraryService:UploadAudioGroup:Failed to persist audio information in database", err)
//	}
//	return nil
//}
//func (s *LibraryService) GetQuestionGroups(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedGroupResponse, *errors.AppError) {
//
//	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
//	defer cancel()
//
//	getQuestionGroups, err := s.repo.GetQuestionGroups(ctx, pageNumber, pageSize)
//	if err != nil {
//		logger.Error("LibraryService:GetParts:Failed to get parts", "error", err)
//		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetQuestionGroups:Failed to Get Question group", err)
//	}
//	// Convert to DTO
//	groupDTOs := mapper.ToPaginatedGroupsResponse(getQuestionGroups)
//	return groupDTOs, nil
//}
//func (s *LibraryService) GetQuestionsByGroups(ctx context.Context, groupId uuid.UUID) ([]*dto.QuestionResponse, error) {
//	questionDBs, err := s.repo.GetQuestionsByGroups(ctx, groupId)
//	if err != nil {
//		logger.Error("LibraryService:GetQuestionGroup:Failed to get questions from group", err)
//		return nil, err
//	}
//	var questions []*dto.QuestionResponse
//	for _, questionDB := range questionDBs {
//		question := mapper.ToQuestionResponse(questionDB)
//		questions = append(questions, question)
//	}
//	return questions, nil
//}
//func (s *LibraryService) CreateQuestion(ctx context.Context, request *dto.CreateQuestionRequest, groupId uuid.UUID) (*dto.QuestionResponse, error) {
//	questionEntity := mapper.ToQuestionEntityFromCreate(request)
//	question, err := s.repo.CreateQuestion(ctx, questionEntity, groupId)
//	if err != nil {
//		logger.Error("LibraryService:CreateQuestion: failed to create question", err)
//		return nil, err
//	}
//	response := mapper.ToQuestionResponse(question)
//	return response, nil
//}
//func (s *LibraryService) UpdateQuestion(ctx context.Context, request *dto.UpdateQuestionRequest, questionId uuid.UUID) error {
//	questionEntity := mapper.ToQuestionEntityFromUpdate(request)
//	err := s.repo.UpdateQuestion(ctx, questionEntity, questionId)
//	if err != nil {
//		logger.Error("LibraryService:CreateQuestion: failed to create question", err)
//		return err
//	}
//	return nil
//}

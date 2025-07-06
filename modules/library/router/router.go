package router

import (
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/middleware"
	"pirate-lang-go/modules/library/controller"
)

type LibraryRouter struct {
	controller *controller.LibraryController
}

func NewLibraryRouter(controller *controller.LibraryController) *LibraryRouter {
	return &LibraryRouter{
		controller: controller,
	}
}
func (r *LibraryRouter) Setup(e *echo.Echo, middleware *middleware.Middleware) {
	// API v1 group
	v1 := e.Group("/v1")
	//public group - no middleware needed
	public := v1.Group("/public")
	publicExams := public.Group("/exams")
	publicExams.GET("", r.controller.GetExams)
	// Admin routes
	admin := v1.Group("/admin")
	// Exam routes
	examsAdmin := admin.Group("/exams")
	examsAdmin.GET("", r.controller.GetExams)
	examsAdmin.POST("", r.controller.CreateExam)
	examsAdmin.GET("/:examId", r.controller.GetExam)
	examsAdmin.PUT("/:examId", r.controller.UpdateExam)
	examsAdmin.GET("/:examId/parts", r.controller.GetExamPartsByExam)

	examPartsAdmin := admin.Group("/parts")

	examPartsAdmin.POST("", r.controller.CreateExamPart)
	examPartsAdmin.GET("/:partId", r.controller.GetExamPart)
	examPartsAdmin.PUT("/:partId", r.controller.UpdateExamPart)
	examPartsAdmin.GET("/:partId/paragraphs", r.controller.GetParagraphsByPart)
	paragraphsAdmin := admin.Group("/paragraphs")
	paragraphsAdmin.POST("", r.controller.CreateParagraph)
	paragraphsAdmin.GET("/:paragraphId", r.controller.GetParagraph)
	paragraphsAdmin.PUT("/:paragraphId", r.controller.UpdateParagraph)

	paragraphsAdmin.POST("/:paragraphId/audio", r.controller.UploadAudioParagraph)
	paragraphsAdmin.POST("/:paragraphId/image", r.controller.UploadImageParagraph)
	paragraphsAdmin.POST("/:paragraphId/transcript", r.controller.UploadTranscriptAudioParagraph)
	// Paragraph Routes
	practicePartsAdmin := admin.Group("/practice-parts")
	practicePartsAdmin.GET("", r.controller.GetPracticeParts)
	practicePartsAdmin.POST("", r.controller.CreateExamPart)
	practicePartsAdmin.GET("/:partId", r.controller.GetExamPart)
	practicePartsAdmin.PUT("/:partId", r.controller.UpdateExamPart)
	practicePartsAdmin.GET("/:partId/paragraphs", r.controller.GetParagraphsByPart)
	//questions := admin.Group("/questions")
	//groups := questions.Group("/groups")
	//groups.GET("", r.controller.GetQuestionGroups)
	//groups.POST("", r.controller.CreateQuestionGroup)
	//groups.PUT("/:groupId", r.controller.UpdateQuestionGroup)
	//groups.POST("/:groupId/audio", r.controller.UploadAudioGroup)
	//groups.POST("/:groupId/image", r.controller.UploadImageGroup)
	//groups.POST("/:groupId/transcript", r.controller.UploadTranscriptAudioGroup)
	test := v1.Group("/test2")
	test.GET("/hello", r.controller.HelloWorld)

}

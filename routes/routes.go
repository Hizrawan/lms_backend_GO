package routes

import (
	"lms-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// =======================
		// Auth
		// =======================
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginUser)

		// =======================
		// Users
		// =======================
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUserByID)

		// =======================
		// Courses
		// =======================
		api.GET("/courses", controllers.GetCourses)
		api.POST("/courses", controllers.CreateCourse)
		api.GET("/courses/:id", controllers.GetCourseByID)
		api.PUT("/courses/:id", controllers.UpdateCourse)
		api.DELETE("/courses/:id", controllers.DeleteCourse)

		// =======================
		// Modules & Lessons
		// =======================
		api.POST("/courses/:id/modules", controllers.CreateModule)
		api.GET("/modules/:id/lessons", controllers.GetLessons)
		api.POST("/modules/:id/lessons", controllers.CreateLesson)

		// =======================
		// Enrollments
		// =======================
		api.POST("/courses/:id/enroll", controllers.EnrollCourse)
		api.GET("/users/:id/enrollments", controllers.GetUserEnrollments)

		// =======================
		// Discussion Forum
		// =======================
		api.GET("/threads", controllers.GetThreads)
		api.GET("/threads/:id", controllers.GetThreadByID)
		api.POST("/threads", controllers.CreateThread)
		api.POST("/threads/:id/replies", controllers.CreateReply)
		api.GET("/threads/:id/replies", controllers.GetReplies)
		// api.POST("/threads/:id/like", controllers.LikeThread)
		// api.POST("/replies/:id/like", controllers.LikeReply)
		// api.POST("/threads/:id/subscribe", controllers.SubscribeThread)

		// =======================
		// Attachments
		// =======================
		api.POST("/attachments", controllers.UploadAttachment)

		// =======================
		// Assignments
		// =======================
		api.GET("/assignments", controllers.GetAssignments)
		api.POST("/assignments", controllers.CreateAssignment)
		api.POST("/assignments/:id/submit", controllers.SubmitAssignment)
		api.GET("/assignments/:id/submissions", controllers.GetAssignmentSubmissions)

		// =======================
		// Quizzes
		// =======================
		api.GET("/quizzes", controllers.GetQuizzes)
		api.POST("/quizzes", controllers.CreateQuiz)
		api.POST("/quizzes/:id/answer", controllers.SubmitAnswer)
		api.GET("/quizzes/:id/results", controllers.GetQuizResults)

		// =======================
		// Certificates
		// =======================
		api.GET("/certificates", controllers.GetCertificates)
		api.POST("/certificates", controllers.IssueCertificate)

		// =======================
		// Tags & Categories
		// =======================
		api.GET("/tags", controllers.GetTags)
		api.POST("/tags", controllers.CreateTag)
		api.GET("/categories", controllers.GetCategories)
		api.POST("/categories", controllers.CreateCategory)

		// =======================
		// Feedback
		// =======================
		api.GET("/feedback/:course_id", controllers.GetCourseFeedback)
		api.POST("/feedback/:course_id", controllers.PostFeedback)
	}

	return r
}

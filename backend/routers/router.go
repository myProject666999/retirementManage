package routers

import (
	"retirementManage/controllers"
	"retirementManage/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		api.POST("/login", controllers.Login)

		auth := api.Group("")
		auth.Use(middleware.JWTAuth())
		{
			auth.GET("/user/info", controllers.GetUserInfo)
			auth.POST("/logout", controllers.Logout)

			users := auth.Group("/users")
			{
				users.GET("", controllers.GetUsers)
				users.GET("/:id", controllers.GetUser)
				users.POST("", controllers.CreateUser)
				users.PUT("/:id", controllers.UpdateUser)
				users.DELETE("/:id", controllers.DeleteUser)
				users.POST("/:id/reset-password", controllers.ResetPassword)
			}

			channels := auth.Group("/channels")
			{
				channels.GET("/all", controllers.GetAllChannels)
				channels.GET("", controllers.GetChannels)
				channels.GET("/:id", controllers.GetChannel)
				channels.POST("", controllers.CreateChannel)
				channels.PUT("/:id", controllers.UpdateChannel)
				channels.DELETE("/:id", controllers.DeleteChannel)
			}

			roomTypes := auth.Group("/room-types")
			{
				roomTypes.GET("/all", controllers.GetAllRoomTypes)
				roomTypes.GET("", controllers.GetRoomTypes)
				roomTypes.GET("/:id", controllers.GetRoomType)
				roomTypes.POST("", controllers.CreateRoomType)
				roomTypes.PUT("/:id", controllers.UpdateRoomType)
				roomTypes.DELETE("/:id", controllers.DeleteRoomType)
			}

			buildings := auth.Group("/buildings")
			{
				buildings.GET("/all", controllers.GetAllBuildings)
				buildings.GET("", controllers.GetBuildings)
				buildings.GET("/:id", controllers.GetBuilding)
				buildings.POST("", controllers.CreateBuilding)
				buildings.PUT("/:id", controllers.UpdateBuilding)
				buildings.DELETE("/:id", controllers.DeleteBuilding)
			}

			consultations := auth.Group("/consultations")
			{
				consultations.GET("", controllers.GetConsultations)
				consultations.GET("/:id", controllers.GetConsultation)
				consultations.POST("", controllers.CreateConsultation)
				consultations.PUT("/:id", controllers.UpdateConsultation)
				consultations.DELETE("/:id", controllers.DeleteConsultation)
			}

			reservations := auth.Group("/reservations")
			{
				reservations.GET("", controllers.GetReservations)
				reservations.GET("/:id", controllers.GetReservation)
				reservations.POST("", controllers.CreateReservation)
				reservations.PUT("/:id", controllers.UpdateReservation)
				reservations.DELETE("/:id", controllers.DeleteReservation)
			}

			elders := auth.Group("/elders")
			{
				elders.GET("/all", controllers.GetAllElders)
				elders.GET("", controllers.GetElders)
				elders.GET("/:id", controllers.GetElder)
				elders.POST("", controllers.CreateElder)
				elders.PUT("/:id", controllers.UpdateElder)
				elders.DELETE("/:id", controllers.DeleteElder)
			}

			employees := auth.Group("/employees")
			{
				employees.GET("", controllers.GetEmployees)
				employees.GET("/:id", controllers.GetEmployee)
				employees.POST("", controllers.CreateEmployee)
				employees.PUT("/:id", controllers.UpdateEmployee)
				employees.DELETE("/:id", controllers.DeleteEmployee)
			}

			serviceItems := auth.Group("/service-items")
			{
				serviceItems.GET("/all", controllers.GetAllServiceItems)
				serviceItems.GET("", controllers.GetServiceItems)
				serviceItems.GET("/:id", controllers.GetServiceItem)
				serviceItems.POST("", controllers.CreateServiceItem)
				serviceItems.PUT("/:id", controllers.UpdateServiceItem)
				serviceItems.DELETE("/:id", controllers.DeleteServiceItem)
			}

			careLevels := auth.Group("/care-levels")
			{
				careLevels.GET("/all", controllers.GetAllCareLevels)
				careLevels.GET("", controllers.GetCareLevels)
				careLevels.GET("/:id", controllers.GetCareLevel)
				careLevels.POST("", controllers.CreateCareLevel)
				careLevels.PUT("/:id", controllers.UpdateCareLevel)
				careLevels.DELETE("/:id", controllers.DeleteCareLevel)
			}

			serviceReservations := auth.Group("/service-reservations")
			{
				serviceReservations.GET("", controllers.GetServiceReservations)
				serviceReservations.GET("/:id", controllers.GetServiceReservation)
				serviceReservations.POST("", controllers.CreateServiceReservation)
				serviceReservations.PUT("/:id", controllers.UpdateServiceReservation)
				serviceReservations.DELETE("/:id", controllers.DeleteServiceReservation)
			}

			dishes := auth.Group("/dishes")
			{
				dishes.GET("/all", controllers.GetAllDishes)
				dishes.GET("", controllers.GetDishes)
				dishes.GET("/:id", controllers.GetDish)
				dishes.POST("", controllers.CreateDish)
				dishes.PUT("/:id", controllers.UpdateDish)
				dishes.DELETE("/:id", controllers.DeleteDish)
			}

			mealPackages := auth.Group("/meal-packages")
			{
				mealPackages.GET("/all", controllers.GetAllMealPackages)
				mealPackages.GET("", controllers.GetMealPackages)
				mealPackages.GET("/:id", controllers.GetMealPackage)
				mealPackages.POST("", controllers.CreateMealPackage)
				mealPackages.PUT("/:id", controllers.UpdateMealPackage)
				mealPackages.DELETE("/:id", controllers.DeleteMealPackage)
			}

			cateringOrders := auth.Group("/catering-orders")
			{
				cateringOrders.GET("", controllers.GetCateringOrders)
				cateringOrders.GET("/:id", controllers.GetCateringOrder)
				cateringOrders.PUT("/:id/status", controllers.UpdateCateringOrderStatus)
			}

			recharges := auth.Group("/recharges")
			{
				recharges.GET("", controllers.GetRecharges)
				recharges.GET("/:id", controllers.GetRecharge)
				recharges.POST("", controllers.CreateRecharge)
				recharges.POST("/:id/confirm", controllers.ConfirmRecharge)
				recharges.POST("/:id/cancel", controllers.CancelRecharge)
			}

			expenses := auth.Group("/expenses")
			{
				expenses.GET("", controllers.GetExpenseRecords)
				expenses.GET("/:id", controllers.GetExpenseRecord)
			}

			contracts := auth.Group("/contracts")
			{
				contracts.GET("", controllers.GetContracts)
				contracts.GET("/:id", controllers.GetContract)
				contracts.POST("", controllers.CreateContract)
				contracts.PUT("/:id", controllers.UpdateContract)
			}

			outings := auth.Group("/outings")
			{
				outings.GET("", controllers.GetOutings)
				outings.GET("/:id", controllers.GetOuting)
				outings.POST("", controllers.CreateOuting)
				outings.PUT("/:id", controllers.UpdateOuting)
			}

			visits := auth.Group("/visits")
			{
				visits.GET("", controllers.GetVisits)
				visits.GET("/:id", controllers.GetVisit)
				visits.POST("", controllers.CreateVisit)
				visits.PUT("/:id", controllers.UpdateVisit)
			}

			accidents := auth.Group("/accidents")
			{
				accidents.GET("", controllers.GetAccidents)
				accidents.GET("/:id", controllers.GetAccident)
				accidents.POST("", controllers.CreateAccident)
				accidents.PUT("/:id", controllers.UpdateAccident)
			}

			checkouts := auth.Group("/checkouts")
			{
				checkouts.GET("", controllers.GetCheckouts)
				checkouts.GET("/:id", controllers.GetCheckout)
				checkouts.POST("/:id/approve", controllers.ApproveCheckout)
				checkouts.POST("/:id/reject", controllers.RejectCheckout)
			}
		}
	}

	return r
}

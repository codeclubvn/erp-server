package repo

// func GenerateToken(ctx *gin.Context) {
// 	var request model.TokenRequest
// 	var user model.User
// 	if err := ctx.ShouldBindJSON(&request); err != nil {
// 		ctx.JSON(400, gin.H{
// 			"error": err.Error(),
// 		})
// 		ctx.Abort()
// 		return
// 	}
//
// 	// Kiểm tra xem email có tồn tại trong database không và password có đúng
// 	// không
// 	record := config.DbDefault.Where("email = ?", request.Email).First(&user)
// 	if record.Error != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": record.Error.Error(),
// 		})
// 		ctx.Abort()
// 		return
// 	}
// 	credentialError := handler.IUser.CheckPassword(user.Password)
// 	if credentialError != nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{
// 			"error": credentialError.Error(),
// 		})
// 		ctx.Abort()
// 		return
// 	}
//
// 	tokenString, err := middleware.GenerateJWT(user.Email, user.Password)
//
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		ctx.Abort()
// 		return
// 	}
//
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"token": tokenString,
// 	})
// }

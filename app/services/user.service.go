package services

/*
 ユーザー登録
*/
// func CreateUser(w http.ResponseWriter, createUser *models.User, signUpRequestParam models.SignUpRequest) error {
// 	hashPassword := logic.ChangeHashPassword(signUpRequestParam.Password)
// 	// 登録データを作成
// 	createUser.Name = signUpRequestParam.Name
// 	createUser.Email = signUpRequestParam.Email
// 	createUser.Password = string(hashPassword)

// 	if err := repositories.CreateUser(createUser); err != nil {
// 		logic.SendResponse(w, logic.CreateErrorStringResponse("ユーザー登録処理に失敗"), http.StatusInternalServerError)
// 		return err
// 	}

// 	return nil
// }
package controller

import (
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"strings"
	"time"
	"todo-togo/model"
	"todo-togo/service"
)

//TODO move this as env var or separated config
const JWTKeyString = "p@ssw0rd@layStr!ng"

type UserController struct {
	uService service.IUserService
}

func NewUserController(uService *service.IUserService) UserController  {
	return	UserController{uService: *uService}
}

func (u *UserController) UserRoute(app *fiber.App)  {
	//app.Get("/users", u.usr.)
	//app.Get("/user/:id", controller.GetAll)

	app.Put("/user", u.Update)
	app.Post("/user", u.NewUser)
	app.Delete("/user/:id", u.Delete)

	app.Post("/user/login", u.Login)
}

// Login godoc
// @Description Get token for authentication header [NOT IMPLEMENTED]
// @Summary Get login token
// @Tags User
// @Accept json
// @Produce json
// @Param Login body model.LoginInput true "Login details"
// @Success 200 {object} model.JsonSuccessLoginReturn
// @Failure 500 {object} model.JsonInternalErrorReturn
// @Failure 404 {object} model.JsonBadRequestErrorReturn
// @Router /user/login [post]
func (u *UserController) Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		log.Println("error to parse login details,",err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Wrong login details",
		})

	}
	req := model.UserReqResponse{
		Email:    strings.ToLower(input.Email),
		Password: input.Password,
	}

	res, err := u.uService.GetUser(req)
	if res == nil {
		log.Println("Login error, user not found")
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"error": "Wrong login details",
		})
	}


	password := input.Password
	password += res.Salt
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))

	if res.Password != hashedPassword {
		log.Println("Login error, incorrect password")
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"error": "Password is incorrect",
		})
	}

	res.Password = ""
	res.Salt = ""

	token := jwt.New(jwt.SigningMethodHS256)
	tokenClaim := jwt.MapClaims{}
	tokenClaim["user_id"] = res.UserID
	tokenClaim["user_name"] = res.Name
	tokenClaim["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token.Claims = tokenClaim

	tokenString, err := token.SignedString([]byte(JWTKeyString))
	if err != nil {
		log.Println("Error token signing,",err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "User on login process",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"token": tokenString,
	})

}

// PostUser godoc
// @Description Add new user, "salt" and "id" should be empty
// @Summary Add new user
// @Tags User
// @Accept json
// @Produce json
// @Param NewUser body model.UserReqResponse true "User details"
// @Success 200 {object} model.JsonSuccessUserReturn
// @Failure 500 {object} model.JsonInternalErrorReturn
// @Failure 404 {object} model.JsonBadRequestErrorReturn
// @Router /user [post]
func (u *UserController) NewUser(c *fiber.Ctx) error {
	var input model.UserReqResponse

	if err := c.BodyParser(&input); err != nil {
		log.Println("error to parse user details,",err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Wrong input details",
		})

	}

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	input.Salt = func (n int) string {
		b := make([]byte, n)
		for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
		return string(b)
	}(32)

	password:= input.Password
	password += input.Salt

	h := sha256.New()
	h.Write([]byte(password))
	password = fmt.Sprintf("%x", h.Sum(nil))

	//set back password with encrypted version
	input.Password=password

	input.Email=strings.ToLower(input.Email)

	res, err := u.uService.AddUser(input)
	if res == nil {
		log.Println("Error on add user,",err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Unable to add user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"user": res,
	})

}


// PutUser godoc
// @Description Update user data, "salt" should be empty
// @Summary Update user data
// @Tags User
// @Accept json
// @Produce json
// @Param User body model.UserReqResponse true "User details"
// @Success 200 {object} model.JsonSuccessUserReturn
// @Failure 500 {object} model.JsonInternalErrorReturn
// @Failure 404 {object} model.JsonBadRequestErrorReturn
// @Router /user [put]
func (u *UserController) Update(c *fiber.Ctx) error  {
	//TODO add validation for mandatory input params
	var req model.UserReqResponse
	if err := c.BodyParser(&req); err != nil {
		log.Println("error to parse body-request updated request,",err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Error on user details",
		})
	}

	//check if ID exist
	old, err := u.uService.GetUser(req)
	if err != nil {
		log.Println("error id not found,",err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Error on update, record not found",
		})
	}

	if req.Password!=""{
		password:= req.Password
		password += old.Salt

		h := sha256.New()
		h.Write([]byte(password))
		password = fmt.Sprintf("%x", h.Sum(nil))

		req.Password=password
	} else {
		req.Password=old.Password
	}

	if req.Email!=""{
		req.Email=strings.ToLower(req.Email)
	} else {
		req.Email=old.Email
	}

	user, err := u.uService.UpdateUser(req)
	if err != nil {
		log.Println("error on modify user record,",err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Error when modify user",
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"user":   user,
	})
}

// DeleteUser godoc
// @Description Delete user data
// @Summary Delete user data
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.JsonDeleteUserReturn
// @Failure 500 {object} model.JsonInternalErrorReturn
// @Failure 404 {object} model.JsonBadRequestErrorReturn
// @Router /user/{id} [delete]
func (u *UserController) Delete(c *fiber.Ctx) error  {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("error to parse requested id,",err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Error to get deleted user",
		})
	}

	deletedID := model.UserReqResponse{
		UserID: id,
	}

	err = u.uService.DeleteUser(deletedID)
	if err != nil {
		log.Println("error to delete user record,",err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Error when delete user",
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"user":   nil,
	})
}
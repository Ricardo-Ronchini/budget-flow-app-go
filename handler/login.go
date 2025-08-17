package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/ricardo-ronchini/budget-flow-app-go/db"
)

var jwtSecret = []byte("sua-chave-secreta-super-segura") // ideal usar env

type LoginCredentials struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}
	defer db.Close()

	var creds LoginCredentials

	// Faz o bind do JSON para a struct
	if err := c.Bind(&creds); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Formato do JSON inválido",
		})
	}

	// Validação fictícia (exemplo simples) | mock
	if creds.Email != "teste@exemplo.com" || creds.Password != "123456" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"erro": "Credenciais inválidas",
		})
	}

	// Criação do token JWT | mock
	claims := jwt.MapClaims{
		"user_id": "abc123",
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"erro": "Erro ao gerar o token",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenStr,
	})
}

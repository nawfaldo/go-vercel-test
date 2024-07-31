package user

import (
	"fmt"
	"net/http"
	"vercer/types"
	"vercer/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var session *sessions.CookieStore

func RegisterRoutes(router *mux.Router, store *sessions.CookieStore) {
	session = store

	router.HandleFunc("/login", utils.Cors(handleLogin)).Methods("POST")
	router.HandleFunc("/register", utils.Cors(handleRegister)).Methods("POST")
	router.HandleFunc("/auth", utils.Cors(handleAuth)).Methods("GET")
	router.HandleFunc("/logout", utils.Cors(handleLogout)).Methods("POST")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	user := GetUserByName(payload.Name)

	cookie, _ := session.Get(r, "kukis")
	cookie.Options = &sessions.Options{
		MaxAge:   3600 * 24, // 1 day
		SameSite: http.SameSiteStrictMode,
	}

	cookie.Values["user"] = user.ID
	cookie.Save(r, w)

	utils.WriteJSON(w, http.StatusAccepted, nil)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	userId := uuid.New().String()

	err := CreateUser(types.User{
		ID:       userId,
		Name:     payload.Name,
		Password: payload.Password,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	cookie, _ := session.Get(r, "kukis")
	cookie.Options = &sessions.Options{
		MaxAge:   3600 * 24, // 1 day
		SameSite: http.SameSiteStrictMode,
	}

	cookie.Values["user"] = userId
	cookie.Save(r, w)

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	cookie, _ := session.Get(r, "kukis")

	userId := cookie.Values["user"]

	if userId == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not authorized"))
		return
	}

	user := GetUserById(userId.(string))

	utils.WriteJSON(w, http.StatusAccepted, user)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := session.Get(r, "kukis")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	cookie.Values["user"] = nil
	cookie.Options = &sessions.Options{
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
	}

	if err := cookie.Save(r, w); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

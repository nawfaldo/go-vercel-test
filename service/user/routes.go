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

	router.HandleFunc("/login", handleLogin).Methods("POST")
	router.HandleFunc("/register", handleRegister).Methods("POST")
	router.HandleFunc("/auth", handleAuth).Methods("GET")
	router.HandleFunc("/logout", handleLogout).Methods("POST")
	router.HandleFunc("/chats", handleChats).Methods("GET")
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
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	}

	cookie.Values["user"] = user.ID
	if err := cookie.Save(r, w); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Log cookie details for debugging
	fmt.Printf("Cookie set: %+v\n", cookie)

	utils.WriteJSON(w, http.StatusAccepted, nil)
}

func handleChats(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusAccepted, "hallo")
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
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
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
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	}

	if err := cookie.Save(r, w); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

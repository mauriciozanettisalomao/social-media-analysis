package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	factories "github.com/mauriciozanettisalomao/social-media-analysis/internal/app/factory"
	socialMedia "github.com/mauriciozanettisalomao/social-media-analysis/pkg/analysis"

	"github.com/gorilla/mux"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.Header().Set("Version", "1.2")

	w.WriteHeader(code)
	w.Write(response)
}

// SearchMention gets the mentions as requested
func SearchMention(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	socialMediaParam := vars["socialmedia"]
	param := vars["param"]
	targetTranslation := vars["targetTranslation"]

	clientSocialMedia := factories.NewSocialMedia(socialMediaParam)
	if clientSocialMedia == nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("social media %s is not implemented yet", socialMediaParam))
	}

	filters := make(map[string]string, 0)

	values := r.URL.Query()
	for k, v := range values {
		for _, filter := range v {
			filters[k] = filter
		}
	}

	socialMediaResp, err := socialMedia.Search(socialMediaParam, param, targetTranslation, filters, clientSocialMedia)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("could not search mentions: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, socialMediaResp)

}

// SearchUser gets the user data as requested
func SearchUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	socialMediaParam := vars["socialmedia"]
	username := vars["username"]

	clientSocialMedia := factories.NewSocialMedia(socialMediaParam)
	if clientSocialMedia == nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("social media %s is not implemented yet", socialMediaParam))
	}

	filters := make(map[string]string, 0)

	values := r.URL.Query()
	for k, v := range values {
		for _, filter := range v {
			filters[k] = filter
		}
	}

	socialMediaResp, err := socialMedia.SearchUser(username, filters, clientSocialMedia)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("could not search user data: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, socialMediaResp)

}

// SearchFollowers gets the followers of a user as requested
func SearchFollowers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	socialMediaParam := vars["socialmedia"]
	username := vars["username"]

	clientSocialMedia := factories.NewSocialMedia(socialMediaParam)
	if clientSocialMedia == nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("social media %s is not implemented yet", socialMediaParam))
	}

	filters := make(map[string]string, 0)

	values := r.URL.Query()
	for k, v := range values {
		for _, filter := range v {
			filters[k] = filter
		}
	}

	socialMediaResp, err := socialMedia.SearchFollowers(username, filters, clientSocialMedia)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("could not search followers of a user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, socialMediaResp)

}

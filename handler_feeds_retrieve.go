package main

import "net/http"

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	dbFeeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get the feeds")
	}

	respondWithJson(w, http.StatusOK, databaseFeedsToFeeds(dbFeeds))
}

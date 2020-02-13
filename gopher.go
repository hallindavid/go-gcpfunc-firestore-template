package gopher

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
)

// FirestoreLookup should query firestore
func FirestoreLookup(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("object_id")
	if id == "" || len(id) < 1 {
		http.Error(w, "Invalid object ID", http.StatusUnprocessableEntity)
		return
	}

	projectID := "ENTER-YOUR-FIRESTORE-PROJECT-ID-HERE"
	collectionName := "ENTER-YOUR-COLLECTION-NAME-HERE"

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Close client when done.
	defer client.Close()

	objectDoc, err := client.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		http.Error(w, "Object Not Found", http.StatusUnprocessableEntity)
		return
	}

	js, err := json.Marshal(objectDoc.Data())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

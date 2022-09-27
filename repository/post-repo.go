package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/hamza-starcevic/goRest/entity"
)

// * Implements the PostRepository interface
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}

// * save post to firestore
func (m *repo) Save(post *entity.Post) (*entity.Post, error) {
	firestoreClient, ctx := createClient()

	_, _, err := firestoreClient.Collection("posts").Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Text":  post.Text,
		"Title": post.Title,
	})
	if err != nil {
		log.Printf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

// * find all posts from firestore
func (m *repo) FindAll() ([]entity.Post, error) {

	firestoreClient, ctx := createClient()
	defer firestoreClient.Close()

	var posts []entity.Post
	iterator := firestoreClient.Collection("posts").Documents(ctx)

	//?Find all posts from firestore
	for {
		var post entity.Post
		doc, err := iterator.Next()
		if err != nil {
			log.Printf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		doc.DataTo(&post)
		posts = append(posts, post)
		if doc == nil {
			break
		}
	}

	return posts, nil
}

// ! Establish a connection to the firestore database
func createClient() (*firestore.Client, context.Context) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "bookstoreproject-209f2")
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
	}
	return client, ctx
}

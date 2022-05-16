package main

import (
	"fmt"
	"log"

	database "github.com/pyumz/social/internal"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database was created and/or already exists!")

	user, err := c.CreateUser("jane@test.com", "123", "jane", 20)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	updateUser, err := c.UpdateUser("jane@test.com", "123", "janet", 20)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User updated", updateUser)

	gotUser, err := c.GetUser("jane@test.com")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Got user", gotUser)

	err = c.DeleteUser("jane@test.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User was deleted from Users database")

	_, err = c.GetUser("test@test.com")
	if err == nil {
		log.Fatal("User shouldn't exist")
	}

	fmt.Println("Confirmed that user is deleted")

	// Test for Posts

	user, err = c.CreateUser("test@test.com", "pass", "joe", 22)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User created to write a post", user)

	post, err := c.CreatePost("test@test.com", "My first post is about my cat!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Post created", post)

	secondPost, err := c.CreatePost("test@test.com", "My second post is about my second cat, of course")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Second post created", secondPost)

	posts, err := c.GetPosts("test@test.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found posts", posts)

	err = c.DeletePost(post.ID)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("First post was deleted")

	posts, err = c.GetPosts("test@test.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found posts", posts)

	err = c.DeletePost(secondPost.ID)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Second post was deleted")

	err = c.DeleteUser("test@test.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted from database, again")
}

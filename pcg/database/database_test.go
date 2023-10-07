package database

import (
	"fmt"
	"testing"

	"GoNews/pcg/typeStruct"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndReadFromDB(t *testing.T) {
	InitDB()

	defer DB.Close()

	// Создаем тестовый пост
	testPost := typeStruct.Post{
		Title:   "Test Title 2",
		Content: "Test Content",
		PubTime: 1692645688,
		Link:    "http://example.com/test",
	}

	// Сохраняем тестовый пост в базу данных
	err := SaveToDB(testPost)
	if err != nil {
		t.Fatalf("Failed to save post to DB: %v", err)
	}

	// Читаем пост из базы данных по названию
	readPost, err := ReadFromDB("Test Title 2")
	if err != nil {
		t.Fatalf("Failed to read post from DB: %v", err)
	}

	// Сравниваем значения
	if readPost.Title != testPost.Title ||
		readPost.Content != testPost.Content ||
		readPost.PubTime != testPost.PubTime ||
		readPost.Link != testPost.Link {
		t.Errorf("Saved data doesn't match expected data.")
	}
}

func TestDeleteByTitle(t *testing.T) {
	InitDB()

	defer DB.Close()

	// Создаем тестовый пост
	testPost := typeStruct.Post{
		Title:   "Test Title 1",
		Content: "Test Content",
		PubTime: 1692645688,
		Link:    "http://example.com/test",
	}

	// Сохраняем тестовый пост в базу данных
	err := SaveToDB(testPost)
	assert.NoError(t, err, "Failed to save post to DB")

	// Удаляем пост по названию
	err = DeleteByTitle("Test Title 3")
	assert.NoError(t, err, "Failed to delete post by title")

	// Пытаемся прочитать пост с удаленным названием
	_, err = ReadFromDB("Test Title 3")
	assert.Error(t, err, "Expected an error when trying to read deleted post")
}

func TestSearchPostsByKeyword(t *testing.T) {
	// Инициализация базы данных и выполнение схемы
	db := InitDB()
	defer db.Close()

	// Вставка тестовых данных
	post1 := typeStruct.Post{
		Title:   "aa24 f=f2 +++ 56ty",
		Content: "Test Description 1",
		PubTime: 1234567890,
		Link:    "http://example.com/test1",
	}

	// Сохранение тестовых данных в базе данных
	if err := SaveToDB(post1); err != nil {
		t.Fatalf("Failed to save post to DB: %v", err)
	}

	// Вызов функции для поиска
	keyword := "F=F2"
	posts, err := SearchPostsByKeyword(keyword)
	if err != nil {
		t.Fatalf("SearchPostsByKeyword failed: %v", err)
	}

	// Проверка результатов
	if len(posts) != 1 {
		t.Fatalf("Expected 1 post, but got %d", len(posts))
	} else {
		fmt.Println(posts)
	}
	if posts[0].Title != post1.Title {
		t.Fatalf("Expected post with title '%s', but got '%s'", post1.Title, posts[0].Title)
	}

	keyword = " "
	posts, err = SearchPostsByKeyword(keyword)
	if err != nil {
		t.Fatalf("SearchPostsByKeyword failed: %v", err)
	}

	// Проверка результатов
	if len(posts) < 1 {
		t.Fatalf("Expected 1 post, but got %d", len(posts))
	}

	keyword = ""
	posts, err = SearchPostsByKeyword(keyword)
	if err != nil {
		t.Fatalf("SearchPostsByKeyword failed: %v", err)
	}

	// Проверка результатов
	if len(posts) < 1 {
		t.Fatalf("Expected 1 post, but got %d", len(posts))
	}

	// Удаление тестовых записей из базы данных
	if err := DeleteByTitle(post1.Title); err != nil {
		t.Fatalf("Failed to delete test record: %v", err)
	}
}

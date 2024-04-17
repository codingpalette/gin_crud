package models

import (
	"database/sql"
	"fmt"
)

type Post struct {
	Id int `json:"id"`
	// title 타입은 string 또는 null 일 수 있음
	Title   sql.NullString `json:"title"`
	Content sql.NullString `json:"content"`
}

// 게시글 리스트 조회
func GetPostList(db *sql.DB) ([]Post, error) {
	// "SELECT * FROM posts" SQL 쿼리를 실행하여 결과 rows를 받음
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Println(err)
		return nil, err // 에러가 발생하면 nil과 에러를 반환
	}
	defer rows.Close() // 함수 종료 시 rows를 닫음

	var posts []Post  // Post 타입의 슬라이스를 생성하여 결과를 담을 변수 posts 선언
	for rows.Next() { // rows.Next() 메서드를 사용하여 다음 레코드가 있는지 확인하는 반복문
		var post Post // Post 구조체 변수 post 선언
		// 현재 레코드의 값을 post 변수에 스캔하여 가져옴
		if err := rows.Scan(&post.Id, &post.Title, &post.Content); err != nil {
			fmt.Println(err)
			return nil, err // 스캔 에러가 발생하면 nil과 에러를 반환
		}
		// 가져온 레코드 정보를 posts 슬라이스에 추가
		posts = append(posts, post)
	}

	return posts, nil // 모든 레코드 처리가 완료되면 posts 슬라이스를 반환
}

// 게시글 생성
func PostCreate(db *sql.DB, title, content string) (Post, error) {
	result, err := db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}

	newPost := Post{
		Id:      int(lastId),
		Title:   sql.NullString{String: title, Valid: true},
		Content: sql.NullString{String: content, Valid: true},
	}

	return newPost, nil
}

// 게시글 하나 조회
func GetPost(db *sql.DB, id string) (Post, error) {
	result, err := db.Query("SELECT * FROM posts WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}

	var post Post
	for result.Next() {
		if err := result.Scan(&post.Id, &post.Title, &post.Content); err != nil {
			fmt.Println(err)
			return Post{}, err
		}
	}

	return post, nil
}

// 게시글 수정
func PostUpdate(db *sql.DB, id int, title, content string) (Post, error) {
	result, err := db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", title, content, id)
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}

	if affected == 0 {
		return Post{}, fmt.Errorf("No post found")
	}

	updatedPost := Post{
		Id:      int(id),
		Title:   sql.NullString{String: title, Valid: true},
		Content: sql.NullString{String: content, Valid: true},
	}

	return updatedPost, nil
}

// 게시글 삭제
func PostDelete(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if affected == 0 {
		return fmt.Errorf("No post found")
	}

	return nil
}

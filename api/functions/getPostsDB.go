package functions

import (
	"context"
	"fmt"

	"github.com/BR7T/Blog/config"
	"github.com/BR7T/Blog/structs"
)

func GetPostDatabase(page int, size int) ([]structs.GetPost, error) {
    var list []structs.GetPost

    conn, err := config.Connection()
    if err != nil {
		fmt.Println(err)
        return nil, err
    }
    defer conn.Close(context.Background())

    rows, err := conn.Query(
        context.Background(),
        "SELECT id, title , content, created_at, updated_at FROM pages LIMIT $1 OFFSET $2",size,size*(page - 1),
    )
    if err != nil {
		fmt.Println(err)
        return nil, err
    }
    defer rows.Close() 

    for rows.Next() {
        var row structs.GetPost
        if err := rows.Scan(&row.ID, &row.Title , &row.Content, &row.CreatedAt, &row.UpdatedAt); err != nil {
            return nil, fmt.Errorf("scanning row: %w", err) 
        }
        list = append(list, row)
    }

    if err := rows.Err(); err != nil {
		fmt.Println(err)
        return nil, fmt.Errorf("iterating rows: %w", err)
    }

    return list, nil
}

func GetPostIDDatabase(idPost int) (*structs.GetPost , error){
    conn, err := config.Connection()
    if err != nil {
		fmt.Println(err)
        return nil, err
    }
    defer conn.Close(context.Background())

    rowQ := conn.QueryRow(
        context.Background(),
        "SELECT id, title  , content, created_at, updated_at FROM pages WHERE id = $1", idPost,
    )

    var row structs.GetPost
        if err := rowQ.Scan(&row.ID, &row.Title , &row.Content, &row.CreatedAt, &row.UpdatedAt); err != nil {
            return nil, fmt.Errorf("scanning row: %w", err) 
    }

    return &row, nil
}
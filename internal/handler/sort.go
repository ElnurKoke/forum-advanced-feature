package handler

import "forum/internal/models"

type ByCreatedAt []models.Post

func (a ByCreatedAt) Len() int           { return len(a) }
func (a ByCreatedAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreatedAt) Less(i, j int) bool { return a[i].CreateAt.After(a[j].CreateAt) }

type ByCreatedAtMes []models.Message

func (a ByCreatedAtMes) Len() int      { return len(a) }
func (a ByCreatedAtMes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByCreatedAtMes) Less(i, j int) bool {
	return a[i].CreateAt.After(a[j].CreateAt)
}

type ByCreatedAtCom []models.Comment

func (a ByCreatedAtCom) Len() int           { return len(a) }
func (a ByCreatedAtCom) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreatedAtCom) Less(i, j int) bool { return a[i].Created_at.After(a[j].Created_at) }

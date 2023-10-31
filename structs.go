package main

import (
	"gorm.io/gorm"
	"time"
)

type Config struct {
	Token          string
	Timeout        int
	SetCommands    bool
	Database       string
	MotivationPath string `yaml:"motivationPath"`
	Commands       map[string]string
	Ranks          map[string]Rank
	Tasks          map[int]TaskData
	Motivations    map[string]Motivation
}

type Rank struct {
	ID     string
	Name   string
	Score  int
	Levels map[int]string
}

type TaskData struct {
	Points int
	Task   string
}

type Motivation struct {
	UUID      string `gorm:"primaryKey"`
	Pack      string
	PackPlace int
	ID        string
	Category  string
	Language  string
	Extension string
	Path      string
}

type User struct {
	gorm.Model
	ID            int64  `gorm:"primaryKey"`
	Username      string
}

type Journey struct {
	gorm.Model           `yaml:"-"`
	UserID     int64     `yaml:"-"`
	RankSystem string
	Start      time.Time
	End        time.Time
	Text       string
}

type Entry struct {
	gorm.Model           `yaml:"-"`
	UserID       int64   `yaml:"-"`
	IsPublic     bool
	CreatedAtStr string  `yaml:"createdat"`
	Note         int
	Text         string  `gorm:"size:4096"`
}

type Task struct {
	gorm.Model           `yaml:"-"`
	UserID    int64		 `yaml:"-"`
	ChatID    int64		 `yaml:"-"`
	MessageID string	 `yaml:"-"`
	TaskID    int		 `yaml:"-"`
	Date      time.Time  `gorm:"autoCreateTime"`
	Done      time.Time  `gorm:"autoUpdateTime"`
	Text      string
	IsDone    bool
}

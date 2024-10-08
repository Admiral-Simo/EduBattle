// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createExam = `-- name: CreateExam :one
INSERT INTO exams (lesson_id, title, max_score, student_score, exam_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, lesson_id, title, max_score, student_score, exam_url
`

type CreateExamParams struct {
	LessonID     pgtype.Int4
	Title        string
	MaxScore     pgtype.Int4
	StudentScore pgtype.Int4
	ExamUrl      pgtype.Text
}

func (q *Queries) CreateExam(ctx context.Context, arg CreateExamParams) (Exam, error) {
	row := q.db.QueryRow(ctx, createExam,
		arg.LessonID,
		arg.Title,
		arg.MaxScore,
		arg.StudentScore,
		arg.ExamUrl,
	)
	var i Exam
	err := row.Scan(
		&i.ID,
		&i.LessonID,
		&i.Title,
		&i.MaxScore,
		&i.StudentScore,
		&i.ExamUrl,
	)
	return i, err
}

const createExercise = `-- name: CreateExercise :one
INSERT INTO exercises (lesson_id, title, max_score, student_score, exercise_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, lesson_id, title, max_score, student_score, exercise_url
`

type CreateExerciseParams struct {
	LessonID     pgtype.Int4
	Title        string
	MaxScore     pgtype.Int4
	StudentScore pgtype.Int4
	ExerciseUrl  pgtype.Text
}

func (q *Queries) CreateExercise(ctx context.Context, arg CreateExerciseParams) (Exercise, error) {
	row := q.db.QueryRow(ctx, createExercise,
		arg.LessonID,
		arg.Title,
		arg.MaxScore,
		arg.StudentScore,
		arg.ExerciseUrl,
	)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.LessonID,
		&i.Title,
		&i.MaxScore,
		&i.StudentScore,
		&i.ExerciseUrl,
	)
	return i, err
}

const createFile = `-- name: CreateFile :one
INSERT INTO files (group_id, uploaded_by, file_path)
VALUES ($1, $2, $3)
RETURNING id, group_id, uploaded_by, file_path, uploaded_at
`

type CreateFileParams struct {
	GroupID    pgtype.Int4
	UploadedBy pgtype.Int4
	FilePath   string
}

func (q *Queries) CreateFile(ctx context.Context, arg CreateFileParams) (File, error) {
	row := q.db.QueryRow(ctx, createFile, arg.GroupID, arg.UploadedBy, arg.FilePath)
	var i File
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.UploadedBy,
		&i.FilePath,
		&i.UploadedAt,
	)
	return i, err
}

const createGroup = `-- name: CreateGroup :one
INSERT INTO groups (name, trophies)
VALUES ($1, $2)
RETURNING id, name, trophies
`

type CreateGroupParams struct {
	Name     string
	Trophies pgtype.Int4
}

func (q *Queries) CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error) {
	row := q.db.QueryRow(ctx, createGroup, arg.Name, arg.Trophies)
	var i Group
	err := row.Scan(&i.ID, &i.Name, &i.Trophies)
	return i, err
}

const createLesson = `-- name: CreateLesson :one
INSERT INTO lessons (title, subject_id, group_id, completed, lesson_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, title, subject_id, group_id, completed, lesson_url
`

type CreateLessonParams struct {
	Title     string
	SubjectID pgtype.Int4
	GroupID   pgtype.Int4
	Completed pgtype.Bool
	LessonUrl pgtype.Text
}

func (q *Queries) CreateLesson(ctx context.Context, arg CreateLessonParams) (Lesson, error) {
	row := q.db.QueryRow(ctx, createLesson,
		arg.Title,
		arg.SubjectID,
		arg.GroupID,
		arg.Completed,
		arg.LessonUrl,
	)
	var i Lesson
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.SubjectID,
		&i.GroupID,
		&i.Completed,
		&i.LessonUrl,
	)
	return i, err
}

const createMessage = `-- name: CreateMessage :one
INSERT INTO messages (group_id, sender_id, message)
VALUES ($1, $2, $3)
RETURNING id, group_id, sender_id, message, created_at
`

type CreateMessageParams struct {
	GroupID  pgtype.Int4
	SenderID pgtype.Int4
	Message  string
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.db.QueryRow(ctx, createMessage, arg.GroupID, arg.SenderID, arg.Message)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.SenderID,
		&i.Message,
		&i.CreatedAt,
	)
	return i, err
}

const createSubject = `-- name: CreateSubject :one
INSERT INTO subjects (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateSubject(ctx context.Context, name string) (Subject, error) {
	row := q.db.QueryRow(ctx, createSubject, name)
	var i Subject
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createTrophy = `-- name: CreateTrophy :one
INSERT INTO trophies (group_id, lesson_id, amount)
VALUES ($1, $2, $3)
RETURNING id, group_id, lesson_id, earned_at, amount
`

type CreateTrophyParams struct {
	GroupID  pgtype.Int4
	LessonID pgtype.Int4
	Amount   pgtype.Int4
}

func (q *Queries) CreateTrophy(ctx context.Context, arg CreateTrophyParams) (Trophy, error) {
	row := q.db.QueryRow(ctx, createTrophy, arg.GroupID, arg.LessonID, arg.Amount)
	var i Trophy
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.LessonID,
		&i.EarnedAt,
		&i.Amount,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, email, password, role, group_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, email, role, group_id
`

type CreateUserParams struct {
	Name     pgtype.Text
	Email    string
	Password string
	Role     pgtype.Text
	GroupID  pgtype.Int4
}

type CreateUserRow struct {
	ID      int32
	Name    pgtype.Text
	Email   string
	Role    pgtype.Text
	GroupID pgtype.Int4
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Role,
		arg.GroupID,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Role,
		&i.GroupID,
	)
	return i, err
}

const deleteExamByID = `-- name: DeleteExamByID :exec
DELETE FROM exams
WHERE id = $1
`

func (q *Queries) DeleteExamByID(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteExamByID, id)
	return err
}

const deleteExerciseByID = `-- name: DeleteExerciseByID :exec
DELETE FROM exercises
WHERE id = $1
`

func (q *Queries) DeleteExerciseByID(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteExerciseByID, id)
	return err
}

const deleteFileByID = `-- name: DeleteFileByID :exec
DELETE FROM files
WHERE id = $1
`

func (q *Queries) DeleteFileByID(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteFileByID, id)
	return err
}

const deleteGroupByID = `-- name: DeleteGroupByID :exec
DELETE FROM groups
WHERE id = $1
`

func (q *Queries) DeleteGroupByID(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteGroupByID, id)
	return err
}

const deleteLessonByID = `-- name: DeleteLessonByID :exec
DELETE FROM lessons
WHERE id = $1
`

func (q *Queries) DeleteLessonByID(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteLessonByID, id)
	return err
}

const deleteMessageByID = `-- name: DeleteMessageByID :exec
DELETE FROM messages
WHERE id = $1
`

func (q *Queries) DeleteMessageByID(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteMessageByID, id)
	return err
}

const deleteSubjectByID = `-- name: DeleteSubjectByID :exec
DELETE FROM subjects
WHERE id = $1
`

func (q *Queries) DeleteSubjectByID(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteSubjectByID, id)
	return err
}

const deleteTrophyByID = `-- name: DeleteTrophyByID :exec
DELETE FROM trophies
WHERE id = $1
`

func (q *Queries) DeleteTrophyByID(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteTrophyByID, id)
	return err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUserByID, id)
	return err
}

const getAllGroups = `-- name: GetAllGroups :many

SELECT id, name, trophies
FROM groups
`

// Group Queries
func (q *Queries) GetAllGroups(ctx context.Context) ([]Group, error) {
	rows, err := q.db.Query(ctx, getAllGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Group
	for rows.Next() {
		var i Group
		if err := rows.Scan(&i.ID, &i.Name, &i.Trophies); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSubjects = `-- name: GetAllSubjects :many

SELECT id, name
FROM subjects
`

// Subject Queries
func (q *Queries) GetAllSubjects(ctx context.Context) ([]Subject, error) {
	rows, err := q.db.Query(ctx, getAllSubjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Subject
	for rows.Next() {
		var i Subject
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUsers = `-- name: GetAllUsers :many

SELECT id, name, email, role, group_id
FROM users
`

type GetAllUsersRow struct {
	ID      int32
	Name    pgtype.Text
	Email   string
	Role    pgtype.Text
	GroupID pgtype.Int4
}

// User Queries
func (q *Queries) GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUsersRow
	for rows.Next() {
		var i GetAllUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Role,
			&i.GroupID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExamByID = `-- name: GetExamByID :one
SELECT id, lesson_id, title, max_score, student_score, exam_url
FROM exams
WHERE id = $1
`

func (q *Queries) GetExamByID(ctx context.Context, id int32) (Exam, error) {
	row := q.db.QueryRow(ctx, getExamByID, id)
	var i Exam
	err := row.Scan(
		&i.ID,
		&i.LessonID,
		&i.Title,
		&i.MaxScore,
		&i.StudentScore,
		&i.ExamUrl,
	)
	return i, err
}

const getExamsByLessonID = `-- name: GetExamsByLessonID :many

SELECT id, lesson_id, title, max_score, student_score, exam_url
FROM exams
WHERE lesson_id = $1
`

// Exam Queries
func (q *Queries) GetExamsByLessonID(ctx context.Context, lessonID pgtype.Int4) ([]Exam, error) {
	rows, err := q.db.Query(ctx, getExamsByLessonID, lessonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Exam
	for rows.Next() {
		var i Exam
		if err := rows.Scan(
			&i.ID,
			&i.LessonID,
			&i.Title,
			&i.MaxScore,
			&i.StudentScore,
			&i.ExamUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExerciseByID = `-- name: GetExerciseByID :one
SELECT id, lesson_id, title, max_score, student_score, exercise_url
FROM exercises
WHERE id = $1
`

func (q *Queries) GetExerciseByID(ctx context.Context, id int32) (Exercise, error) {
	row := q.db.QueryRow(ctx, getExerciseByID, id)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.LessonID,
		&i.Title,
		&i.MaxScore,
		&i.StudentScore,
		&i.ExerciseUrl,
	)
	return i, err
}

const getExercisesByLessonID = `-- name: GetExercisesByLessonID :many

SELECT id, lesson_id, title, max_score, student_score, exercise_url
FROM exercises
WHERE lesson_id = $1
`

// Exercise Queries
func (q *Queries) GetExercisesByLessonID(ctx context.Context, lessonID pgtype.Int4) ([]Exercise, error) {
	rows, err := q.db.Query(ctx, getExercisesByLessonID, lessonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Exercise
	for rows.Next() {
		var i Exercise
		if err := rows.Scan(
			&i.ID,
			&i.LessonID,
			&i.Title,
			&i.MaxScore,
			&i.StudentScore,
			&i.ExerciseUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFilesByGroupID = `-- name: GetFilesByGroupID :many

SELECT id, group_id, uploaded_by, file_path, uploaded_at
FROM files
WHERE group_id = $1
`

// File Sharing Queries
func (q *Queries) GetFilesByGroupID(ctx context.Context, groupID pgtype.Int4) ([]File, error) {
	rows, err := q.db.Query(ctx, getFilesByGroupID, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []File
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.GroupID,
			&i.UploadedBy,
			&i.FilePath,
			&i.UploadedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupByID = `-- name: GetGroupByID :one
SELECT id, name, trophies
FROM groups
WHERE id = $1
`

func (q *Queries) GetGroupByID(ctx context.Context, id int32) (Group, error) {
	row := q.db.QueryRow(ctx, getGroupByID, id)
	var i Group
	err := row.Scan(&i.ID, &i.Name, &i.Trophies)
	return i, err
}

const getLeaderboard = `-- name: GetLeaderboard :many

SELECT g.id AS group_id, g.name AS group_name, SUM(t.amount) AS total_trophies
FROM groups g
LEFT JOIN trophies t ON g.id = t.group_id
GROUP BY g.id
ORDER BY total_trophies DESC
`

type GetLeaderboardRow struct {
	GroupID       int32
	GroupName     string
	TotalTrophies int64
}

// Leaderboard Queries
func (q *Queries) GetLeaderboard(ctx context.Context) ([]GetLeaderboardRow, error) {
	rows, err := q.db.Query(ctx, getLeaderboard)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLeaderboardRow
	for rows.Next() {
		var i GetLeaderboardRow
		if err := rows.Scan(&i.GroupID, &i.GroupName, &i.TotalTrophies); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLessonByID = `-- name: GetLessonByID :one
SELECT id, title, subject_id, group_id, completed, lesson_url
FROM lessons
WHERE id = $1
`

func (q *Queries) GetLessonByID(ctx context.Context, id int32) (Lesson, error) {
	row := q.db.QueryRow(ctx, getLessonByID, id)
	var i Lesson
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.SubjectID,
		&i.GroupID,
		&i.Completed,
		&i.LessonUrl,
	)
	return i, err
}

const getLessonsBySubjectID = `-- name: GetLessonsBySubjectID :many

SELECT id, title, subject_id, group_id, completed, lesson_url
FROM lessons
WHERE subject_id = $1
`

// Lesson Queries
func (q *Queries) GetLessonsBySubjectID(ctx context.Context, subjectID pgtype.Int4) ([]Lesson, error) {
	rows, err := q.db.Query(ctx, getLessonsBySubjectID, subjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Lesson
	for rows.Next() {
		var i Lesson
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.SubjectID,
			&i.GroupID,
			&i.Completed,
			&i.LessonUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMessagesByGroupID = `-- name: GetMessagesByGroupID :many

SELECT id, group_id, sender_id, message, created_at
FROM messages
WHERE group_id = $1
`

// Message Queries
func (q *Queries) GetMessagesByGroupID(ctx context.Context, groupID pgtype.Int4) ([]Message, error) {
	rows, err := q.db.Query(ctx, getMessagesByGroupID, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.GroupID,
			&i.SenderID,
			&i.Message,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubjectByID = `-- name: GetSubjectByID :one
SELECT id, name
FROM subjects
WHERE id = $1
`

func (q *Queries) GetSubjectByID(ctx context.Context, id int32) (Subject, error) {
	row := q.db.QueryRow(ctx, getSubjectByID, id)
	var i Subject
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getTrophiesByGroupID = `-- name: GetTrophiesByGroupID :many

SELECT id, group_id, lesson_id, earned_at, amount
FROM trophies
WHERE group_id = $1
`

// Trophies Queries
func (q *Queries) GetTrophiesByGroupID(ctx context.Context, groupID pgtype.Int4) ([]Trophy, error) {
	rows, err := q.db.Query(ctx, getTrophiesByGroupID, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Trophy
	for rows.Next() {
		var i Trophy
		if err := rows.Scan(
			&i.ID,
			&i.GroupID,
			&i.LessonID,
			&i.EarnedAt,
			&i.Amount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTrophyByID = `-- name: GetTrophyByID :one
SELECT id, group_id, lesson_id, earned_at, amount
FROM trophies
WHERE id = $1
`

func (q *Queries) GetTrophyByID(ctx context.Context, id int32) (Trophy, error) {
	row := q.db.QueryRow(ctx, getTrophyByID, id)
	var i Trophy
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.LessonID,
		&i.EarnedAt,
		&i.Amount,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, email, role, group_id
FROM users
WHERE id = $1
`

type GetUserByIDRow struct {
	ID      int32
	Name    pgtype.Text
	Email   string
	Role    pgtype.Text
	GroupID pgtype.Int4
}

func (q *Queries) GetUserByID(ctx context.Context, id int32) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Role,
		&i.GroupID,
	)
	return i, err
}

const getUsersByGroupID = `-- name: GetUsersByGroupID :many
SELECT id, name, email, role, group_id
FROM users
WHERE group_id = $1
`

type GetUsersByGroupIDRow struct {
	ID      int32
	Name    pgtype.Text
	Email   string
	Role    pgtype.Text
	GroupID pgtype.Int4
}

func (q *Queries) GetUsersByGroupID(ctx context.Context, groupID pgtype.Int4) ([]GetUsersByGroupIDRow, error) {
	rows, err := q.db.Query(ctx, getUsersByGroupID, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersByGroupIDRow
	for rows.Next() {
		var i GetUsersByGroupIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Role,
			&i.GroupID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExamScore = `-- name: UpdateExamScore :exec
UPDATE exams
SET student_score = $2
WHERE id = $1
`

type UpdateExamScoreParams struct {
	ID           int32
	StudentScore pgtype.Int4
}

func (q *Queries) UpdateExamScore(ctx context.Context, arg UpdateExamScoreParams) error {
	_, err := q.db.Exec(ctx, updateExamScore, arg.ID, arg.StudentScore)
	return err
}

const updateExerciseScore = `-- name: UpdateExerciseScore :exec
UPDATE exercises
SET student_score = $2
WHERE id = $1
`

type UpdateExerciseScoreParams struct {
	ID           int32
	StudentScore pgtype.Int4
}

func (q *Queries) UpdateExerciseScore(ctx context.Context, arg UpdateExerciseScoreParams) error {
	_, err := q.db.Exec(ctx, updateExerciseScore, arg.ID, arg.StudentScore)
	return err
}

const updateGroupTrophies = `-- name: UpdateGroupTrophies :exec
UPDATE groups
SET trophies = $2
WHERE id = $1
`

type UpdateGroupTrophiesParams struct {
	ID       int32
	Trophies pgtype.Int4
}

func (q *Queries) UpdateGroupTrophies(ctx context.Context, arg UpdateGroupTrophiesParams) error {
	_, err := q.db.Exec(ctx, updateGroupTrophies, arg.ID, arg.Trophies)
	return err
}

const updateLessonCompletion = `-- name: UpdateLessonCompletion :exec
UPDATE lessons
SET completed = $2
WHERE id = $1
`

type UpdateLessonCompletionParams struct {
	ID        int32
	Completed pgtype.Bool
}

func (q *Queries) UpdateLessonCompletion(ctx context.Context, arg UpdateLessonCompletionParams) error {
	_, err := q.db.Exec(ctx, updateLessonCompletion, arg.ID, arg.Completed)
	return err
}

const updateUserGroup = `-- name: UpdateUserGroup :exec
UPDATE users
SET group_id = $2
WHERE id = $1
`

type UpdateUserGroupParams struct {
	ID      int32
	GroupID pgtype.Int4
}

func (q *Queries) UpdateUserGroup(ctx context.Context, arg UpdateUserGroupParams) error {
	_, err := q.db.Exec(ctx, updateUserGroup, arg.ID, arg.GroupID)
	return err
}

-- User Queries

-- name: GetAllUsers :many
SELECT id, name, email, role, group_id
FROM users;

-- name: GetUserByID :one
SELECT id, name, email, role, group_id
FROM users
WHERE id = $1;

-- name: GetUsersByGroupID :many
SELECT id, name, email, role, group_id
FROM users
WHERE group_id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email, password, role, group_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, email, role, group_id;

-- name: UpdateUserGroup :exec
UPDATE users
SET group_id = $2
WHERE id = $1;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1;

-- Group Queries

-- name: GetAllGroups :many
SELECT id, name, trophies
FROM groups;

-- name: GetGroupByID :one
SELECT id, name, trophies
FROM groups
WHERE id = $1;

-- name: CreateGroup :one
INSERT INTO groups (name, trophies)
VALUES ($1, $2)
RETURNING id, name, trophies;

-- name: UpdateGroupTrophies :exec
UPDATE groups
SET trophies = $2
WHERE id = $1;

-- name: DeleteGroupByID :exec
DELETE FROM groups
WHERE id = $1;

-- Subject Queries

-- name: GetAllSubjects :many
SELECT id, name
FROM subjects;

-- name: GetSubjectByID :one
SELECT id, name
FROM subjects
WHERE id = $1;

-- name: CreateSubject :one
INSERT INTO subjects (name)
VALUES ($1)
RETURNING id, name;

-- name: DeleteSubjectByID :exec
DELETE FROM subjects
WHERE id = $1;

-- Lesson Queries

-- name: GetLessonsBySubjectID :many
SELECT id, title, subject_id, group_id, completed, lesson_url
FROM lessons
WHERE subject_id = $1;

-- name: GetLessonByID :one
SELECT id, title, subject_id, group_id, completed, lesson_url
FROM lessons
WHERE id = $1;

-- name: CreateLesson :one
INSERT INTO lessons (title, subject_id, group_id, completed, lesson_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, title, subject_id, group_id, completed, lesson_url;

-- name: UpdateLessonCompletion :exec
UPDATE lessons
SET completed = $2
WHERE id = $1;

-- name: DeleteLessonByID :exec
DELETE FROM lessons
WHERE id = $1;

-- Exercise Queries

-- name: GetExercisesByLessonID :many
SELECT id, lesson_id, title, max_score, student_score, exercise_url
FROM exercises
WHERE lesson_id = $1;

-- name: GetExerciseByID :one
SELECT id, lesson_id, title, max_score, student_score, exercise_url
FROM exercises
WHERE id = $1;

-- name: CreateExercise :one
INSERT INTO exercises (lesson_id, title, max_score, student_score, exercise_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, lesson_id, title, max_score, student_score, exercise_url;

-- name: UpdateExerciseScore :exec
UPDATE exercises
SET student_score = $2
WHERE id = $1;

-- name: DeleteExerciseByID :exec
DELETE FROM exercises
WHERE id = $1;

-- Exam Queries

-- name: GetExamsByLessonID :many
SELECT id, lesson_id, title, max_score, student_score, exam_url
FROM exams
WHERE lesson_id = $1;

-- name: GetExamByID :one
SELECT id, lesson_id, title, max_score, student_score, exam_url
FROM exams
WHERE id = $1;

-- name: CreateExam :one
INSERT INTO exams (lesson_id, title, max_score, student_score, exam_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, lesson_id, title, max_score, student_score, exam_url;

-- name: UpdateExamScore :exec
UPDATE exams
SET student_score = $2
WHERE id = $1;

-- name: DeleteExamByID :exec
DELETE FROM exams
WHERE id = $1;

-- Trophies Queries

-- name: GetTrophiesByGroupID :many
SELECT id, group_id, lesson_id, earned_at, amount
FROM trophies
WHERE group_id = $1;

-- name: GetTrophyByID :one
SELECT id, group_id, lesson_id, earned_at, amount
FROM trophies
WHERE id = $1;

-- name: CreateTrophy :one
INSERT INTO trophies (group_id, lesson_id, amount)
VALUES ($1, $2, $3)
RETURNING id, group_id, lesson_id, earned_at, amount;

-- name: DeleteTrophyByID :exec
DELETE FROM trophies
WHERE id = $1;

-- Message Queries

-- name: GetMessagesByGroupID :many
SELECT id, group_id, sender_id, message, created_at
FROM messages
WHERE group_id = $1;

-- name: CreateMessage :one
INSERT INTO messages (group_id, sender_id, message)
VALUES ($1, $2, $3)
RETURNING id, group_id, sender_id, message, created_at;

-- name: DeleteMessageByID :exec
DELETE FROM messages
WHERE id = $1;

-- File Sharing Queries

-- name: GetFilesByGroupID :many
SELECT id, group_id, uploaded_by, file_path, uploaded_at
FROM files
WHERE group_id = $1;

-- name: CreateFile :one
INSERT INTO files (group_id, uploaded_by, file_path)
VALUES ($1, $2, $3)
RETURNING id, group_id, uploaded_by, file_path, uploaded_at;

-- name: DeleteFileByID :exec
DELETE FROM files
WHERE id = $1;

-- Leaderboard Queries

-- name: GetLeaderboard :many
SELECT g.id AS group_id, g.name AS group_name, SUM(t.amount) AS total_trophies
FROM groups g
LEFT JOIN trophies t ON g.id = t.group_id
GROUP BY g.id
ORDER BY total_trophies DESC;

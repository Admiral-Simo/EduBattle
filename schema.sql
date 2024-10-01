-- User Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'student', -- 'admin' or 'student'
    group_id INT REFERENCES groups(id) ON DELETE SET NULL
);

-- Group Table
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    trophies INT DEFAULT 0 -- Total trophies earned by the group
);

-- Subject Table
CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Lesson Table (Linked to a Subject)
CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    subject_id INT REFERENCES subjects(id) ON DELETE CASCADE,
    group_id INT REFERENCES groups(id) ON DELETE CASCADE,
    completed BOOLEAN DEFAULT FALSE,
    lesson_url VARCHAR(255) -- URL attached to the lesson (e.g., a video or reading material)
);

-- Exercise Table (Linked to a Lesson)
CREATE TABLE exercises (
    id SERIAL PRIMARY KEY,
    lesson_id INT REFERENCES lessons(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    max_score INT DEFAULT 100,
    student_score INT,
    exercise_url VARCHAR(255) -- URL to the exercise resource (e.g., problem set or interactive quiz)
);

-- Exam Table (Linked to a Lesson)
CREATE TABLE exams (
    id SERIAL PRIMARY KEY,
    lesson_id INT REFERENCES lessons(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    max_score INT DEFAULT 100,
    student_score INT,
    exam_url VARCHAR(255) -- URL to the exam resource (e.g., exam sheet or quiz platform)
);

-- Trophies Table (Logs Trophies Earned by Groups)
CREATE TABLE trophies (
    id SERIAL PRIMARY KEY,
    group_id INT REFERENCES groups(id) ON DELETE CASCADE,
    lesson_id INT REFERENCES lessons(id) ON DELETE CASCADE,
    earned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    amount INT DEFAULT 1 -- Number of trophies earned
);

-- Messages Table (For Intra-group Communication)
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    group_id INT REFERENCES groups(id) ON DELETE CASCADE,
    sender_id INT REFERENCES users(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- File Sharing Table (To Share Photos/Documents Within a Group)
CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    group_id INT REFERENCES groups(id) ON DELETE CASCADE,
    uploaded_by INT REFERENCES users(id) ON DELETE CASCADE,
    file_path VARCHAR(255) NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Leaderboard View (Dynamic View to Track Group Rankings)
CREATE VIEW leaderboard AS
SELECT
    g.id AS group_id,
    g.name AS group_name,
    SUM(t.amount) AS total_trophies
FROM groups g
LEFT JOIN trophies t ON g.id = t.group_id
GROUP BY g.id
ORDER BY total_trophies DESC;

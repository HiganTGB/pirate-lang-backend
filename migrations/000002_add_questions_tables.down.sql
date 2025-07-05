-- ======================
-- Trigger
-- ======================
DROP TRIGGER IF EXISTS update_questions_updated_at ON questions;
DROP TRIGGER IF EXISTS update_exams_updated_at ON exams;
DROP TRIGGER IF EXISTS update_paragraphs_updated_at ON paragraphs;
DROP TRIGGER IF EXISTS update_exam_parts_updated_at ON exam_parts;
-- ======================
-- Table
-- ======================
DROP TABLE IF EXISTS questions;

DROP TABLE IF EXISTS paragraphs;

DROP TABLE IF EXISTS exam_parts;

DROP TABLE IF EXISTS exams;
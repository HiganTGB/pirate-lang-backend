-- ========================
-- DROP TRIGGER
-- ========================
DROP TRIGGER IF EXISTS set_question_options_update_at ON question_options;
DROP TRIGGER IF EXISTS set_questions_update_at ON questions;
DROP TRIGGER IF EXISTS set_question_group_update_at ON question_groups;
DROP TRIGGER IF EXISTS set_question_group_version ON question_groups;
DROP TRIGGER IF EXISTS set_parts_update_at ON parts;

-- ========================
-- DROP TRIGGER FUNCTION
-- ========================
DROP FUNCTION IF EXISTS update_version_column();

-- ========================
-- DROP TABLE
-- ========================
DROP TABLE IF EXISTS question_options;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS question_groups;
DROP TABLE IF EXISTS parts;
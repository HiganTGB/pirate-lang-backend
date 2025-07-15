-- name: GetUserByEmailOrUserNameOrId :one
-- GetUserByEmailOrUserNameOrId retrieves a user by email, user_name, or id.
SELECT id, user_name, email, password, created_at, updated_at
FROM users
WHERE
    (sqlc.narg(email)::text IS NULL OR email = sqlc.narg(email)::text) AND
    (sqlc.narg(user_name)::text IS NULL OR user_name = sqlc.narg(user_name)::text) AND
    (sqlc.narg(id)::uuid IS NULL OR id = sqlc.narg(id)::uuid)
LIMIT 1;
-- name: CreateAccount :one
-- CreateAccount creates a new user and returns selected fields.
INSERT INTO users (user_name, email, password)
VALUES ($1, $2, $3)
RETURNING id, user_name, email;

-- name: UpdatePassword :execresult
-- UpdatePassword updates the password for a given user ID.
UPDATE users
SET password = $1, updated_at = NOW()
WHERE id = $2;

-- name: GetUsersCount :one
-- GetUsersCount returns the total number of users.
SELECT COUNT(*) FROM users;

-- name: GetPaginatedUsers :many
-- GetPaginatedUsers retrieves a list of users with pagination.
SELECT id, user_name, email, created_at, updated_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateRole :exec
-- CreateRole creates a new role.
INSERT INTO roles (name, description)
VALUES ($1, $2);
-- name: GetRoles :many
-- GetRoles retrieves all roles.
SELECT id, name, description, created_at, updated_at
FROM roles;

-- name: CreatePermission :exec
-- CreatePermission creates a new permission.
INSERT INTO permissions (name, description)
VALUES ($1, $2);

-- name: GetPermissions :many
-- GetPermissions retrieves all permissions.
SELECT id, name, description, created_at, updated_at
FROM permissions;

-- name: AssignPermissionToRole :exec
-- AssignPermissionToRole assigns a permission to a role.
INSERT INTO role_permissions (role_id, permission_id)
VALUES ($1, $2);

-- name: AssignRoleToUser :exec
-- AssignRoleToUser assigns a role to a user.
INSERT INTO user_roles (user_id, role_id)
VALUES ($1, $2);

-- name: RoleExists :one
-- RoleExists checks if a role with the given ID exists.
SELECT EXISTS(SELECT 1 FROM roles WHERE id = $1);

-- name: PermissionExists :one
-- PermissionExists checks if a permission with the given ID exists.
SELECT EXISTS(SELECT 1 FROM permissions WHERE id = $1);

-- name: DeleteRole :exec
-- DeleteRole deletes a role by its ID.
DELETE FROM roles WHERE id = $1;

-- name: DeletePermission :exec
-- DeletePermission deletes a permission by its ID.
DELETE FROM permissions WHERE id = $1;

-- name: HasPermission :one
-- HasPermission checks if a user has a specific permission.
SELECT EXISTS(
    SELECT 1 FROM user_roles ur
                      JOIN role_permissions rp ON ur.role_id = rp.role_id
                      JOIN permissions p ON rp.permission_id = p.id
    WHERE ur.user_id = $1 AND p.id = $2
);
-- name: GetRole :one
SELECT
    r.id AS role_id,
    r.name AS role_name,
    r.description AS role_description,
    p.id AS permission_id,
    p.name AS permission_name,
    p.description AS permission_description
FROM
    roles AS r
        JOIN
    role_permissions AS rp ON r.id = rp.role_id
        JOIN
    permissions AS p ON rp.permission_id = p.id;
-- name: LockUser :execresult
-- LockUser to lock user account
UPDATE users
set is_locked=true,lock_reason=$1,locked_at=now()
where id=$2;

-- name: UnlockUser :execresult
-- UnlockUser to unlock user account
UPDATE users
set is_locked=false,unlock_reason=$1,unlocked_at=now()
where id=$2;
-- 00002

-- name: CreateUserProfile :exec
-- CreateUserProfile creates a new Userprofile.
INSERT INTO user_profiles(user_id, full_name, birthday, gender, phone_number, address, bio)
VALUES($1,$2,$3,$4,$5,$6,$7);
-- name: UpdateUserProfile :exec
Update user_profiles
set full_name = $1,birthday=$2,gender=$3,phone_number=$4,address=$5,bio=$6
where user_id =$7;
-- name: UpdateUserAvatar :exec
Update user_profiles
set avatar_url=$1
where user_id =$2;
-- name: GetUserAvatar :one
SELECT avatar_url
FROM  user_profiles
where user_id =$1;
-- name: GetUserProfile :one
SELECT
    user_id,u.email,u.user_name,full_name,birthday,gender,phone_number,address,avatar_url,bio
FROM
    user_profiles p join users u on p.user_id = u.id
WHERE
    user_id = $1;

-- ========================
-- 002
-- ========================

-- name: CreateExam :one
INSERT INTO Exams (
    exam_title,
    description,
    duration_minutes,
    exam_type,
    max_listening_score,
    max_reading_score,
    max_speaking_score,
    max_writing_score,
    total_score
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9
         ) RETURNING exam_id;

-- name: GetExam :one
SELECT
    exam_id,
    exam_title,
    description,
    duration_minutes,
    exam_type,
    max_listening_score,
    max_reading_score,
    max_speaking_score,
    max_writing_score,
    total_score,
    created_at,
    updated_at
FROM
    Exams
WHERE
    exam_id = $1;

-- name: GetPaginatedExams :many
SELECT
    exam_id,
    exam_title,
    description,
    duration_minutes,
    exam_type,
    max_listening_score,
    max_reading_score,
    max_speaking_score,
    max_writing_score,
    total_score,
    created_at,
    updated_at
FROM
    Exams
LIMIT $1 OFFSET $2;
-- name: UpdateExam :exec
UPDATE Exams
SET
    exam_title = $2,
    description = $3,
    duration_minutes = $4,
    exam_type = $5,
    max_listening_score = $6,
    max_reading_score = $7,
    max_speaking_score = $8,
    max_writing_score = $9,
    total_score = $10
WHERE
    exam_id = $1;

-- name: DeleteExam :exec
DELETE FROM Exams
WHERE
    exam_id = $1;
-- name: GetExamsCount :one
SELECT COUNT(*) FROM exams;

-- name: CreateExamPart :one
INSERT INTO exam_parts (
    exam_id,
    part_title,
    part_order,
    description,
    is_practice_component,
    plan_type,
    toeic_part_number
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         ) RETURNING part_id;

-- name: GetExamPartByID :one
SELECT
    part_id,
    exam_id,
    part_title,
    part_order,
    description,
    is_practice_component,
    plan_type,
    created_at,
    updated_at,
    toeic_part_number
FROM
    exam_parts
WHERE
    part_id = $1;

-- name: GetPaginatedPracticeExamParts :many
SELECT
    part_id,
    exam_id,
    part_title,
    part_order,
    description,
    is_practice_component,
    plan_type,
    created_at,
    updated_at,
    toeic_part_number
FROM
    exam_parts
WHERE
    is_practice_component == 'TRUE'
LIMIT $1 OFFSET $2;
-- name: GetPracticeExamPartCount :one
SELECT COUNT(*) FROM exam_parts where is_practice_component == 'TRUE';
-- name: GetExamPartsByExamId :many
SELECT
    part_id,
    exam_id,
    part_title,
    part_order,
    description,
    is_practice_component,
    plan_type,
    created_at,
    updated_at,
    toeic_part_number
FROM
    exam_parts
WHERE
    exam_id = $1
ORDER BY
    part_order;

-- name: UpdateExamPart :exec
UPDATE exam_parts
SET
    exam_id = $2,
    part_title = $3,
    part_order = $4,
    description = $5,
    is_practice_component = $6,
    plan_type = $7,
    toeic_part_number = $8
WHERE
    part_id = $1;

-- name: DeleteExamPart :exec
DELETE FROM exam_parts
WHERE
    part_id = $1;

---
-- Paragraphs Queries
---

-- name: CreateParagraph :one
INSERT INTO Paragraphs (
    paragraph_content,
    title,
    part_id,
    paragraph_order,
    paragraph_type,
    audio_url,
    image_url
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         ) RETURNING paragraph_id;

-- name: GetParagraphByID :one
SELECT
    paragraph_id,
    paragraph_content,
    title,
    part_id,
    paragraph_order,
    paragraph_type,
    audio_url,
    image_url,
    created_at,
    updated_at
FROM
    Paragraphs
WHERE
    paragraph_id = $1;
-- name: GetParagraphByPartId :many
SELECT
    paragraph_id,
    paragraph_content,
    title,
    part_id,
    paragraph_order,
    paragraph_type,
    audio_url,
    image_url,
    created_at,
    updated_at
FROM
    Paragraphs
WHERE
    part_id = $1;
-- name: ListParagraphs :many
SELECT
    paragraph_id,
    paragraph_content,
    title,
    part_id,
    paragraph_order,
    paragraph_type,
    audio_url,
    image_url,
    created_at,
    updated_at
FROM
    Paragraphs;

-- name: ListParagraphsByPartID :many
SELECT
    paragraph_id,
    paragraph_content,
    title,
    part_id,
    paragraph_order,
    paragraph_type,
    audio_url,
    image_url,
    created_at,
    updated_at
FROM
    Paragraphs
WHERE
    part_id = $1
ORDER BY
    paragraph_order;

-- name: UpdateParagraph :exec
UPDATE Paragraphs
SET
    paragraph_content = $2,
    title = $3,
    part_id = $4,
    paragraph_order = $5,
    paragraph_type = $6
WHERE
    paragraph_id = $1;

-- name: UpdateParagraphAudioURL :exec
UPDATE Paragraphs
SET
    audio_url = $2
WHERE
    paragraph_id = $1;

-- name: UpdateParagraphImageURL :exec
UPDATE Paragraphs
SET
    image_url = $2
WHERE
    paragraph_id = $1;

-- name: DeleteParagraph :exec
DELETE FROM Paragraphs
WHERE
    paragraph_id = $1;

---
-- Questions Queries
---

-- name: CreateQuestion :one
INSERT INTO Questions (
    question_content,
    question_type,
    part_id,
    paragraph_id,
    question_order,
    audio_url,
    image_url,
    toeic_question_section,
    question_number_in_part,
    answer_option,
    correct_answer
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
         ) RETURNING question_id,question_content,question_type,part_id,paragraph_id,question_order,audio_url,image_url,toeic_question_section,question_number_in_part;

-- name: GetQuestionByID :one
SELECT
    question_id,
    question_content,
    question_type,
    part_id,
    paragraph_id,
    question_order,
    audio_url,
    image_url,
    toeic_question_section,
    question_number_in_part,
    answer_option,
    correct_answer,
    created_at,
    updated_at
FROM
    Questions
WHERE
    question_id = $1;

-- name: ListQuestions :many
SELECT
    question_id,
    question_content,
    question_type,
    part_id,
    paragraph_id,
    question_order,
    audio_url,
    image_url,
    toeic_question_section,
    question_number_in_part,
    answer_option,
    correct_answer,
    created_at,
    updated_at
FROM
    Questions;

-- name: ListQuestionsByPartID :many
SELECT
    question_id,
    question_content,
    question_type,
    part_id,
    paragraph_id,
    question_order,
    audio_url,
    image_url,
    toeic_question_section,
    question_number_in_part,
    answer_option,
    correct_answer,
    created_at,
    updated_at
FROM
    Questions
WHERE
    part_id = $1
ORDER BY
    question_order;
-- name: ListQuestionsByParagraphID :many
SELECT
    question_id,
    question_content,
    question_type,
    part_id,
    paragraph_id,
    question_order,
    audio_url,
    image_url,
    toeic_question_section,
    question_number_in_part,
    answer_option,
    correct_answer,
    created_at,
    updated_at
FROM
    Questions
WHERE
    paragraph_id = $1
Order By
    question_order ASC,
    question_number_in_part ASC,
    question_id ASC;
-- name: GetCountSeparateQuestionsByPartID :one
SELECT
    count(*)
FROM
    Questions
WHERE
    part_id = $1 and paragraph_id ISNULL;
-- name: GetPaginatedSeparateQuestionsByPartID :many
SELECT
    question_id,
    question_content,
    question_type,
    part_id,
    paragraph_id,
    question_order,
    audio_url,
    image_url,
    toeic_question_section,
    question_number_in_part,
    answer_option,
    correct_answer,
    created_at,
    updated_at
FROM
    Questions
WHERE
    part_id = $1 and paragraph_id ISNULL
Order By
    question_order ASC,
    question_number_in_part ASC,
    question_id ASC
Limit $2 OFFSET $3;

-- name: UpdateQuestion :exec
UPDATE Questions
SET
    question_content = $2,
    question_type = $3,
    part_id = $4,
    paragraph_id = $5,
    question_order = $6,
    toeic_question_section = $7,
    question_number_in_part = $8,
    answer_option = $9,
    correct_answer = $10
WHERE
    question_id = $1;

-- name: UpdateQuestionAudioURL :exec
UPDATE Questions
SET
    audio_url = $2
WHERE
    question_id = $1;

-- name: UpdateQuestionImageURL :exec
UPDATE Questions
SET
    image_url = $2
WHERE
    question_id = $1;

-- name: DeleteQuestion :exec
DELETE FROM Questions
WHERE
    question_id = $1;
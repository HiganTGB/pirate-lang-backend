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

-- name: CreatePart :exec
INSERT INTO parts(skill, name, description, sequence)
VALUES($1,$2,$3,$4);
-- name: UpdatePart :exec
UPDATE parts
SET skill=$1,name=$2,description=$3,sequence=$4
WHERE part_id=$5;
-- name: GetPartsCount :one
SELECT COUNT(*) FROM parts;

-- name: GetPaginatedParts :many
SELECT part_id, skill, name, description, sequence, created_at, updated_at
FROM parts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetPart :one
SELECT part_id, skill, name, description, sequence, created_at, updated_at
FROM parts
WHERE part_id=$1;


-- name: CreateQuestionGroup :one
INSERT INTO  question_groups
    (name, description, part_id,plan_type, group_type)
VALUES($1,$2,$3,$4,$5)
RETURNING question_group_id;
-- name: UpdateQuestionGroup :exec
UPDATE question_groups
SET name=$1,description=$2,part_id=$3,plan_type=$4,group_type=$5
WHERE question_group_id=$6;
-- name: UpdateTextContentQuestionGroup :exec
UPDATE question_groups
SET context_text_content=$1
WHERE question_group_id=$2;
-- name: UpdateAudioContentQuestionGroup :exec
UPDATE question_groups
SET context_audio_url=$1
WHERE question_group_id=$2;
-- name: UpdateImageContentQuestionGroup :exec
UPDATE question_groups
SET context_image_url=$1
WHERE question_group_id=$2;
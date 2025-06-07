-- Enable uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ========================
-- USERS
-- ========================
CREATE TABLE IF NOT EXISTS users (
                                     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                     user_name VARCHAR(255) NOT NULL UNIQUE DEFAULT '',
                                     email VARCHAR(255) NOT NULL UNIQUE DEFAULT '',
                                     password VARCHAR(255) NOT NULL DEFAULT '',
                                     is_social_login BOOLEAN DEFAULT FALSE,
                                     is_locked BOOLEAN DEFAULT FALSE,
                                     locked_at TIMESTAMPTZ,
                                     lock_reason TEXT DEFAULT '',
                                     unlocked_at TIMESTAMPTZ,
                                     unlock_reason TEXT DEFAULT '',
                                     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- ========================
-- USER PROFILES
-- ========================
CREATE TABLE IF NOT EXISTS user_profiles (
                                             user_id UUID PRIMARY KEY,
                                             full_name VARCHAR(255) DEFAULT '',
                                             birthday DATE,
                                             gender VARCHAR(10) DEFAULT '',
                                             phone_number VARCHAR(20) DEFAULT '',
                                             address TEXT DEFAULT '',
                                             avatar_url TEXT DEFAULT '',
                                             bio TEXT DEFAULT '',
                                             created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                             updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                             CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ========================
-- ROLES
-- ========================
CREATE TABLE IF NOT EXISTS roles (
                                     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                     name VARCHAR(100) NOT NULL UNIQUE DEFAULT '',
                                     description TEXT DEFAULT '',
                                     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- ========================
-- PERMISSIONS
-- ========================
CREATE TABLE IF NOT EXISTS permissions (
                                           id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                           name VARCHAR(100) NOT NULL UNIQUE DEFAULT '',
                                           description TEXT DEFAULT '',
                                           created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                           updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- ========================
-- ROLE-PERMISSIONS
-- ========================
CREATE TABLE IF NOT EXISTS role_permissions (
                                                role_id UUID NOT NULL,
                                                permission_id UUID NOT NULL,
                                                PRIMARY KEY (role_id, permission_id),
                                                FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
                                                FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);

-- ========================
-- USER-ROLES
-- ========================
CREATE TABLE IF NOT EXISTS user_roles (
                                          user_id UUID NOT NULL,
                                          role_id UUID NOT NULL,
                                          PRIMARY KEY (user_id, role_id),
                                          FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                                          FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

-- ========================
-- USER PROVIDERS (OAuth)
-- ========================
CREATE TABLE IF NOT EXISTS user_providers (
                                              id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                              user_id UUID NOT NULL,
                                              provider VARCHAR(50) NOT NULL DEFAULT '',              -- e.g. 'google'
                                              provider_user_id VARCHAR(255) NOT NULL DEFAULT '',     -- ID from Google
                                              email VARCHAR(255) DEFAULT '',
                                              access_token TEXT DEFAULT '',
                                              refresh_token TEXT DEFAULT '',
                                              avatar_url TEXT DEFAULT '',
                                              full_name VARCHAR(255) DEFAULT '',
                                              created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                              updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                              UNIQUE (provider, provider_user_id),
                                              FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ========================
-- TRIGGER FUNCTION: update_updated_at_column
-- ========================
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- ========================
-- TRIGGERS
-- ========================
CREATE TRIGGER set_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_user_profiles_updated_at
    BEFORE UPDATE ON user_profiles
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_roles_updated_at
    BEFORE UPDATE ON roles
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_permissions_updated_at
    BEFORE UPDATE ON permissions
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_user_providers_updated_at
    BEFORE UPDATE ON user_providers
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();


-- ========================
-- Parts
-- ========================
CREATE TABLE parts (
                       part_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
                       skill varchar(20) NOT NULL,
                       name text NOT NULL,
                       description text,
                       sequence int not null,
                       created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       unique (skill, sequence),
                       CONSTRAINT chk_skill CHECK (skill IN ('Listening', 'Reading', 'Speaking', 'Writing'))
);


-- ========================
-- Question_Groups
-- ========================
CREATE TABLE question_groups (
                                 question_group_id uuid primary key DEFAULT uuid_generate_v4(),
                                 name text NOT NULL,
                                 description text,
                                 context_text_content text,
                                 context_audio_url text,
                                 context_image_url text,
                                 part_id uuid not null,
                                 group_type varchar(30) NOT NULL,
                                 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                 updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                 plan_type varchar(30) not null ,
                                 version int default 1 not null,
                                 is_locked BOOLEAN DEFAULT FALSE,
                                 locked_at TIMESTAMPTZ,
                                 lock_reason TEXT DEFAULT '',
                                 unlocked_at TIMESTAMPTZ,
                                 unlock_reason TEXT DEFAULT '',
                                 FOREIGN KEY (part_id) references parts(part_id),
                                 CONSTRAINT chk_group_type CHECK (group_type IN (
                                                                                 'MULTIPLE_CHOICE',
                                                                                 'MULTIPLE_CHOICE_HIDDEN',
                                                                                 'ESSAY'
                                     )),
                                 CONSTRAINT chk_plan_type_group CHECK (plan_type IN (
                                                                                     'SUBSCRIPTION',
                                                                                     'FREE'
                                     ))
);

---

-- ========================
-- Questions
-- ========================
CREATE TABLE questions (
                           question_id uuid primary key DEFAULT uuid_generate_v4(),
                           question_group_id uuid,
                           question_order int NOT NULL,
                           text_content text,
                           audio_url text,
                           image_url text,
                           expected_answer_format varchar(20),
                           created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                           FOREIGN KEY (question_group_id) references question_groups(question_group_id) on delete cascade,
                           CONSTRAINT chk_expected_answer_format CHECK (expected_answer_format IN ('TEXT', 'AUDIO', 'MULTIPLE_CHOICE'))
);
-- ========================
-- Question_Options
-- ========================
CREATE TABLE question_options (
                                  option_id uuid primary key DEFAULT uuid_generate_v4(),
                                  question_id uuid not null,
                                  option_name varchar(10) not null,
                                  option_value text not null,
                                  is_correct_option boolean default false,
                                  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                  FOREIGN KEY (question_id) references questions(question_id) on delete cascade,
                                  unique (question_id, option_name),
                                  CONSTRAINT chk_option_name CHECK (option_name IN ('A', 'B', 'C', 'D'))
);
-- ========================
-- TRIGGER FUNCTION: update_version_column
-- ========================
CREATE OR REPLACE FUNCTION update_version_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.version = OLD.version +1;
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';
-- ========================
-- TRIGGERS
-- ========================
CREATE TRIGGER set_parts_update_at
    BEFORE UPDATE ON parts
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER set_question_group_version
    BEFORE UPDATE ON question_groups
    FOR EACH ROW
EXECUTE FUNCTION update_version_column();
CREATE TRIGGER set_question_group_update_at
    BEFORE UPDATE ON question_groups
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER set_questions_update_at
    BEFORE UPDATE ON Questions
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER set_question_options_update_at
    BEFORE UPDATE ON question_options
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();



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
                                 version int default 1 not null,
                                 plan_type varchar(30) not null ,
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
-- TRIGGER FUNCTION
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

-- ========================
-- Exams
-- ========================
CREATE TABLE exams (
                       exam_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       exam_title VARCHAR(255) NOT NULL,
                       description TEXT,
                       duration_minutes INT,

                       exam_type VARCHAR(50) NOT NULL,
                       max_listening_score INT,
                       max_reading_score INT,
                       max_speaking_score INT,
                       max_writing_score INT,
                       total_score INT,

                       created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

                       CONSTRAINT CHK_ExamType CHECK (
                           exam_type IN (
                                         'TOEIC L&R',
                                         'TOEIC S&W',
                                         'TOEIC Bridge',
                                         'General'
                               )
                           )
);
-- ========================
-- ExamParts
-- ========================
CREATE TABLE exam_parts (
                           part_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                           exam_id UUID, -- Can be NULL for standalone practice components
                           part_title TEXT NOT NULL,
                           part_order INT, -- Can be NULL for standalone practice components
                           description TEXT,

                           is_practice_component BOOLEAN DEFAULT FALSE,
                           plan_type VARCHAR(20) NOT NULL, -- e.g., 'SUBSCRIPTION', 'FREE'

                           created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                           toeic_part_number INT,

                           CONSTRAINT FK_part_exam FOREIGN KEY (exam_id) REFERENCES exams (exam_id),
                           CONSTRAINT chk_plan_type_parts CHECK (plan_type IN (
                                                                               'SUBSCRIPTION',
                                                                               'FREE'
                               ))
);


-- ========================
-- Paragraphs
-- ========================
CREATE TABLE paragraphs (
                            paragraph_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Added default UUID generation
                            paragraph_content TEXT NOT NULL,
                            title VARCHAR(255),
                            part_id UUID NOT NULL,
                            paragraph_order INT NOT NULL,

                            paragraph_type VARCHAR(50), -- e.g., 'Reading Passage', 'Audio Script', 'General Context'
                            audio_url VARCHAR(255),
                            image_url VARCHAR(255),

                            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- Will be updated by trigger

                            FOREIGN KEY (part_id) REFERENCES exam_parts (part_id),
                            CONSTRAINT CHK_ParagraphType CHECK (
                                paragraph_type IN (
                                                   'Reading Passage',
                                                   'Audio Script',
                                                   'General Context',
                                                   'Image Context'
                                    )
                                )
);
-- ========================
-- Questions
-- ========================
CREATE TABLE questions (
                           question_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                           question_content TEXT NOT NULL,
                           question_type VARCHAR(50) NOT NULL,
                           part_id UUID NOT NULL,
                           paragraph_id UUID,

                           question_order INT NOT NULL,
                           audio_url VARCHAR(255),
                           image_url VARCHAR(255),

                           toeic_question_section VARCHAR(20) NOT NULL,
                           question_number_in_part INT,
                           answer_option JSON, -- Storing answer choices as JSON {"A" : TEXT}
                           correct_answer TEXT,

                           created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

                           FOREIGN KEY (part_id) REFERENCES exam_parts (part_id),
                           FOREIGN KEY (paragraph_id) REFERENCES paragraphs(paragraph_id) ON DELETE SET NULL,
                           CONSTRAINT chk_toeic_section CHECK (toeic_question_section IN ('Listening', 'Reading', 'Speaking', 'Writing')),
                           CONSTRAINT chk_question_type CHECK (
                               question_type IN (
                                                 'MultipleChoice',      -- Standard A, B, C, D choices
                                                 'TrueFalse',           -- Simple True/False questions
                                                 'FillInTheBlank',      -- Single word or short phrase fill-in-the-blank
                                                 'Essay',               -- Open-ended, longer text response
                                                 'ShortAnswer',         -- Open-ended, but expecting a concise text response


                                                 'PhotoDescription',    -- For TOEIC L&R Part 1 (describing a picture with options)
                                                 'QuestionResponse',    -- For TOEIC L&R Part 2 (listening to a question and choosing the best response)


                                                 'ReadAloud',           -- For TOEIC S&W Part 1 (reading a text aloud)
                                                 'PictureDescription',  -- For TOEIC S&W Part 2 (describing a picture orally)
                                                 'OpenResponse',        -- For TOEIC S&W Part 3, 4, 5 (responding to questions or expressing opinions orally)


                                                 'Matching',            -- Matching items from two lists
                                                 'Ordering',            -- Arranging items in a specific order
                                                 'Instruction'          -- A question that serves as an instruction for a group of sub-questions (though you removed ParentQuestionID, this type can still be useful for visual grouping)
                                   )
                               )
);
-- ======================
-- Trigger
-- ======================
CREATE TRIGGER update_questions_updated_at
    BEFORE UPDATE ON questions
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_exams_updated_at
    BEFORE UPDATE ON exams
    FOR EACH ROW
EXECUTE  FUNCTION update_updated_at_column();
CREATE TRIGGER update_paragraphs_updated_at
    BEFORE UPDATE ON paragraphs
    FOR EACH ROW
EXECUTE  FUNCTION update_updated_at_column();
CREATE TRIGGER update_exam_parts_updated_at
    BEFORE UPDATE ON exam_parts
    FOR EACH ROW
EXECUTE  FUNCTION update_updated_at_column();


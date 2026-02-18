-- ENUMS
CREATE TYPE user_role          AS ENUM ('admin', 'user');
CREATE TYPE document_status    AS ENUM ('uploading', 'uploaded', 'summarizing', 'ready', 'error');
CREATE TYPE job_type           AS ENUM ('summarize', 'generate_questions');
CREATE TYPE job_status         AS ENUM ('pending', 'running', 'completed', 'failed');
CREATE TYPE question_type      AS ENUM ('multiple_choice', 'multiple_choice_image', 'essay');
CREATE TYPE difficulty_level   AS ENUM ('easy', 'medium', 'hard');
CREATE TYPE qset_status        AS ENUM ('generating', 'completed', 'error');
CREATE TYPE image_source       AS ENUM ('ai_generated', 'extracted_from_pdf', 'user_uploaded');

-- ── USERS ─────────────────────────────────────────────────────
CREATE TABLE users (
  id              UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
  email           VARCHAR(255) NOT NULL UNIQUE,
  name            VARCHAR(255) NOT NULL,
  password_hash   TEXT         NOT NULL,
  role            user_role    NOT NULL DEFAULT 'user',
  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  last_login_at   TIMESTAMPTZ
);

-- ── DOCUMENTS ─────────────────────────────────────────────────
CREATE TABLE documents (
  id                UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id           UUID         NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  title             VARCHAR(255) NOT NULL,
  original_filename VARCHAR(500) NOT NULL,
  cf_object_key     TEXT         NOT NULL,         
  cf_public_url     TEXT         NOT NULL UNIQUE,  
  file_size_bytes   BIGINT       NOT NULL,
  page_count        INT,                            
  status            document_status NOT NULL DEFAULT 'uploading',
  created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_documents_user_id  ON documents(user_id);
CREATE INDEX idx_documents_status   ON documents(status);

-- ── DOCUMENT SUMMARIES ────────────────────────────────────────
CREATE TABLE document_summaries (
  id           UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  document_id  UUID        NOT NULL UNIQUE REFERENCES documents(id) ON DELETE CASCADE,
  user_id      UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  summary_data JSONB       NOT NULL,  
  version      INT         NOT NULL DEFAULT 1,
  is_edited    BOOLEAN     NOT NULL DEFAULT FALSE,
  ai_model     VARCHAR(100) NOT NULL, 
  token_usage  JSONB,                 
  created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  edited_at    TIMESTAMPTZ             
);
CREATE INDEX idx_summaries_document_id ON document_summaries(document_id);
CREATE INDEX idx_summaries_data        ON document_summaries USING GIN(summary_data);

-- ── AI JOBS ───────────────────────────────────────────────────
CREATE TABLE ai_jobs (
  id            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  document_id   UUID        NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
  user_id       UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  job_type      job_type    NOT NULL,
  status        job_status  NOT NULL DEFAULT 'pending',
  error_message TEXT,
  retry_count   INT         NOT NULL DEFAULT 0,
  started_at    TIMESTAMPTZ,
  finished_at   TIMESTAMPTZ,
  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_jobs_status      ON ai_jobs(status) WHERE status = 'pending';
CREATE INDEX idx_jobs_document_id ON ai_jobs(document_id);

-- ── QUESTION IMAGES ───────────────────────────────────────────
CREATE TABLE question_images (
  id               UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  question_set_id  UUID        NOT NULL REFERENCES question_sets(id) ON DELETE CASCADE,
  user_id          UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  cf_object_key    TEXT        NOT NULL,
  cf_public_url    TEXT        NOT NULL UNIQUE,
  source_type      image_source NOT NULL,
  alt_text         TEXT,
  file_size_bytes  INT         NOT NULL DEFAULT 0,
  created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ── QUESTION SETS ─────────────────────────────────────────────
CREATE TABLE question_sets (
  id               UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  document_id      UUID        NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
  summary_id       UUID        NOT NULL REFERENCES document_summaries(id),
  user_id          UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  title            VARCHAR(255) NOT NULL,
  config           JSONB       NOT NULL,  
  total_questions  INT         NOT NULL DEFAULT 0,
  status           qset_status NOT NULL DEFAULT 'generating',
  created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_qsets_user_id     ON question_sets(user_id);
CREATE INDEX idx_qsets_document_id ON question_sets(document_id);

-- ── QUESTIONS ─────────────────────────────────────────────────
CREATE TABLE questions (
  id                UUID             PRIMARY KEY DEFAULT gen_random_uuid(),
  question_set_id   UUID             NOT NULL REFERENCES question_sets(id) ON DELETE CASCADE,
  question_image_id UUID             REFERENCES question_images(id) ON DELETE SET NULL,
  order_num         INT              NOT NULL,
  type              question_type    NOT NULL,
  difficulty        difficulty_level NOT NULL DEFAULT 'medium',
  question_text     TEXT             NOT NULL,
  options           JSONB,           
  correct_answer    TEXT             NOT NULL,
  explanation       TEXT,
  topic_tag         VARCHAR(100),
  created_at        TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
  UNIQUE(question_set_id, order_num)
);
CREATE INDEX idx_questions_set_id ON questions(question_set_id);
CREATE INDEX idx_questions_type   ON questions(type);

-- ── AUDIT LOGS ────────────────────────────────────────────────
CREATE TABLE audit_logs (
  id           BIGSERIAL    PRIMARY KEY,
  user_id      UUID         REFERENCES users(id) ON DELETE SET NULL,
  action       VARCHAR(100) NOT NULL,  
  entity_type  VARCHAR(50)  NOT NULL,
  entity_id    UUID,
  meta         JSONB,
  created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_audit_user_id   ON audit_logs(user_id);
CREATE INDEX idx_audit_entity    ON audit_logs(entity_type, entity_id);
CREATE INDEX idx_audit_created   ON audit_logs(created_at DESC);
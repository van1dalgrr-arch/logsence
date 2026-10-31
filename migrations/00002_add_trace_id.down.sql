DROP INDEX IF EXISTS idx_logs_trace_id;

ALTER TABLE logs DROP COLUMN IF EXISTS trace_id;

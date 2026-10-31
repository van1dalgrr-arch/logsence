ALTER TABLE logs ADD COLUMN IF NOT EXISTS trace_id TEXT;

CREATE INDEX IF NOT EXISTS idx_logs_trace_id ON logs (trace_id);

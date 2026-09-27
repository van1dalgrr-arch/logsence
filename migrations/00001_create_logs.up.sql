CREATE TABLE logs (
    id BIGSERIAL PRIMARY KEY,
    service TEXT NOT NULL,
    level TEXT NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    received_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_logs_service_created ON logs (service, created_at DESC);

CREATE TABLE IF NOT EXISTS planner_entries (
    id TEXT PRIMARY KEY,
    task_id TEXT,
    title TEXT NOT NULL,
    description TEXT,
    start_time DATETIME NOT NULL,
    end_time DATETIME,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_planner_entries_task_id
    ON planner_entries(task_id);

CREATE INDEX IF NOT EXISTS idx_planner_entries_start_time
    ON planner_entries(start_time);
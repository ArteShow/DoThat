CREATE TABLE IF NOT EXISTS learning_packs (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    language TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS learning_words (
    id TEXT PRIMARY KEY,
    pack_id TEXT NOT NULL,
    word TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,

    FOREIGN KEY (pack_id)
        REFERENCES learning_packs(id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS learning_translations (
    id TEXT PRIMARY KEY,
    word_id TEXT NOT NULL,
    language TEXT NOT NULL,
    translation TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,

    FOREIGN KEY (word_id)
        REFERENCES learning_words(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_learning_words_pack_id
    ON learning_words(pack_id);

CREATE INDEX IF NOT EXISTS idx_learning_translations_word_id
    ON learning_translations(word_id);
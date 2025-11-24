-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS answers (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    question_id INT NOT NULL,
    user_id UUID NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_answers_question
        FOREIGN KEY (question_id)
        REFERENCES question(id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE answers
-- +goose StatementEnd

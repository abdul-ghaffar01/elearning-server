CREATE TABLE IF NOT EXISTS correct_choice(
    mcqId uuid REFERENCES mcq(id) ON DELETE CASCADE,
    choiceId uuid REFERENCES choice(id) ON DELETE CASCADE,
    reason TEXT,
    PRIMARY KEY (mcqId, choiceId)
)
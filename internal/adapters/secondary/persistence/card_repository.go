package persistence

import (
	"database/sql"
	"time"

	"github.com/gcaldasl/srs-cli/internal/core/domain"
	"github.com/gcaldasl/srs-cli/internal/core/ports"
)

type SQLiteRepository struct {
    db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) ports.CardRepository {
    return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Create(card *domain.Card) error {
    query := `INSERT INTO cards (front_side, back_side, last_reviewed, next_review, interval, ease_factor) 
              VALUES (?, ?, ?, ?, ?, ?)`
    _, err := r.db.Exec(query, card.FrontSide, card.BackSide, card.LastReviewed, card.NextReview, card.Interval, card.EaseFactor)
    return err
}

func (r *SQLiteRepository) Get(id int64) (*domain.Card, error) {
    query := `SELECT id, front_side, back_side, last_reviewed, next_review, interval, ease_factor FROM cards WHERE id = ?`
    row := r.db.QueryRow(query, id)

    card := &domain.Card{}
    err := row.Scan(&card.ID, &card.FrontSide, &card.BackSide, &card.LastReviewed, &card.NextReview, &card.Interval, &card.EaseFactor)
    if err != nil {
        return nil, err
    }

    return card, nil
}

func (r *SQLiteRepository) Update(card *domain.Card) error {
    query := `UPDATE cards SET front_side = ?, back_side = ?, last_reviewed = ?, next_review = ?, interval = ?, ease_factor = ? WHERE id = ?`
    _, err := r.db.Exec(query, card.FrontSide, card.BackSide, card.LastReviewed, card.NextReview, card.Interval, card.EaseFactor, card.ID)
    return err
}

func (r *SQLiteRepository) Delete(id int64) error {
    query := `DELETE FROM cards WHERE id = ?`
    _, err := r.db.Exec(query, id)
    return err
}

func (r *SQLiteRepository) ListDue() ([]*domain.Card, error) {
    query := `SELECT id, front_side, back_side, last_reviewed, next_review, interval, ease_factor FROM cards WHERE next_review <= ?`
    rows, err := r.db.Query(query, time.Now())
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cards []*domain.Card
    for rows.Next() {
        card := &domain.Card{}
        err := rows.Scan(&card.ID, &card.FrontSide, &card.BackSide, &card.LastReviewed, &card.NextReview, &card.Interval, &card.EaseFactor)
        if err != nil {
            return nil, err
        }
        cards = append(cards, card)
    }

    return cards, nil
}

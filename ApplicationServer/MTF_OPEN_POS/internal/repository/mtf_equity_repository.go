// ApplicationServer\Equity_Pos_Service\internal\repository\equity_repository.go
package repository

import (
	"gorm.io/gorm"
	"github.com/SudarshanZone/MTF_OPEN_POS/internal/models"
)

type EquityRepository interface {
	GetEquityPositions(account string) ([]models.MtfEquityPosition, error)
}

type equityRepository struct {
	db *gorm.DB
}

func NewEquityRepository(db *gorm.DB) EquityRepository {
	return &equityRepository{db: db}
}

func (r *equityRepository) GetEquityPositions(account string) ([]models.MtfEquityPosition, error) {
	query := `
		SELECT * FROM epb_em_pstn_book WHERE epb_clm_mtch_accnt = ?`
	
	var positions []models.MtfEquityPosition
	result := r.db.Raw(query, account).Scan(&positions)
	if result.Error != nil {
		return nil, result.Error
	}
	return positions, nil
}

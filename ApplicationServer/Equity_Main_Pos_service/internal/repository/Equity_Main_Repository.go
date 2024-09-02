// repository/position_repository.go
package repository

import (
	"gorm.io/gorm"
	"github.com/SudarshanZone/Equity_Main_Open_Pos/internal/models"
)

type PositionRepository interface {
	GetEquityPositions(account string) ([]models.OtpEquityPosition, error)
}

type positionRepository struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) PositionRepository {
	return &positionRepository{db: db}
}

func (r *positionRepository) GetEquityPositions(account string) ([]models.OtpEquityPosition, error) {
	query := `
		SELECT * FROM otp_trd_pstn WHERE otp_clm_mtch_acct = ?`
	
	var positions []models.OtpEquityPosition
	result := r.db.Raw(query, account).Scan(&positions)
	if result.Error != nil {
		return nil, result.Error
	}
	return positions, nil
}

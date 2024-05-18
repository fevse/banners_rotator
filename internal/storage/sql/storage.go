package sqlstorage

import (
	"fmt"

	"github.com/fevse/banners_rotator/internal/bandit"
	_ "github.com/jackc/pgx/stdlib" // driver
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

type Storage struct {
	db *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(dsn string) (err error) {
	s.db, err = sqlx.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("db connection error: %w", err)
	}
	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) Migrate(dir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("migration, set dialect error: %w", err)
	}
	if err := goose.Up(s.db.DB, dir); err != nil {
		return fmt.Errorf("migration up error: %w", err)
	}
	return nil
}

func (s *Storage) Add(slot, banner int) error {
	groups := make([]int, 0)
	err := s.db.Select(&groups, `SELECT group_id FROM sgroup`)
	if err != nil {
		return err
	}
	for _, group := range groups {
		_, err = s.db.Exec(`INSERT INTO storage(banner_id, slot_id, group_id)
		VALUES($1, $2, $3);`,
			banner, slot, group)
	}
	return err
}

func (s *Storage) Delete(slot, banner int) error {
	_, err := s.db.Exec(`DELETE FROM storage WHERE banner_id=$1 AND slot_id=$2`, banner, slot)
	return err
}

func (s *Storage) Click(banner, slot, group int) error {
	reward := make([]float64, 0)
	view := make([]int, 0)
	click := make([]int, 0)
	id := make([]int, 0)
	err := s.db.Select(&reward, `SELECT reward FROM storage
		WHERE banner_id=$1 AND slot_id=$2 AND group_id=$3;`, banner, slot, group)
	if err != nil {
		return err
	}

	err = s.db.Select(&click, `SELECT click FROM storage
		WHERE banner_id=$1 AND slot_id=$2 AND group_id=$3;`, banner, slot, group)
	if err != nil {
		return err
	}

	err = s.db.Select(&view, `SELECT view FROM storage
		WHERE banner_id=$1 AND slot_id=$2 AND group_id=$3;`, banner, slot, group)
	if err != nil {
		return err
	}

	err = s.db.Select(&id, `SELECT id FROM storage
		WHERE banner_id=$1 AND slot_id=$2 AND group_id=$3;`, banner, slot, group)
	if err != nil {
		return err
	}

	ban, err := bandit.New(1)
	if err != nil {
		return err
	}

	ban.Clicks[id[0]] = click[0]
	ban.Views[id[0]] = view[0]
	ban.Rewards[id[0]] = reward[0]
	ban.Update(id[0], 1.1)

	_, err = s.db.Exec(`UPDATE storage 
		SET click=$1, reward=$2
		WHERE banner_id=$3 AND slot_id=$4 AND group_id=$5;`, ban.Clicks[id[0]], ban.Rewards[id[0]], banner, slot, group)
	return err
}

func (s *Storage) Choose(slot, group int) (banner int, err error) {
	view := make([]int, 0)
	click := make([]int, 0)
	reward := make([]float64, 0)

	views := make(map[int]int)
	clicks := make(map[int]int)
	rewards := make(map[int]float64)

	banners := make([]int, 0)

	err = s.db.Select(&banners, `SELECT banner_id FROM storage
		WHERE slot_id=$1 AND group_id=$2;`, slot, group)
	if err != nil {
		return 0, err
	}

	for _, banner := range banners {
		err = s.db.Select(&view, `SELECT view FROM storage
			WHERE banner_id=$1 AND slot_id=$2 AND group_id=$3;`, banner, slot, group)
		if err != nil {
			return 0, err
		}

		err = s.db.Select(&click, `SELECT click FROM storage
			WHERE banner_id=$1 AND slot_id=$2 AND group_id=$3;`, banner, slot, group)
		if err != nil {
			return 0, err
		}

		err = s.db.Select(&reward, `SELECT reward FROM storage
			WHERE banner_id=$1 AND slot_id=$2 AND group_id=$3;`, banner, slot, group)
		if err != nil {
			return 0, err
		}
		views[banner] = view[0]
		clicks[banner] = click[0]
		rewards[banner] = reward[0]
	}

	ban, err := bandit.New(len(banners))
	if err != nil {
		return 0, err
	}

	ban.Views = views
	ban.Clicks = clicks
	ban.Rewards = rewards

	res := ban.SelectArm()

	_, err = s.db.Exec(`UPDATE storage 
		SET view=$1
		WHERE banner_id=$2 AND slot_id=$3 AND group_id=$4;`, ban.Views[res], res, slot, group)

	return res, err
}

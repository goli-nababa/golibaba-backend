package storage

import (
	"errors"
	"transportation/internal/common/domain"

	"gorm.io/gorm"
)

// Retrieve multiple records with filtering, sorting, joining, and preloading
func GetRecords[T any](db *gorm.DB, request *domain.RepositoryRequest) ([]T, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	query := db.Model(new(T))

	// Apply preloads
	for _, preload := range request.Preloads {
		if preload != "" {
			query = query.Preload(preload)
		}
	}

	// Apply filters
	for _, filter := range request.Filters {
		switch filter.Operator {
		case "=":
			query = query.Where(filter.Field+" = ?", filter.Value)
		case "!=":
			query = query.Where(filter.Field+" != ?", filter.Value)
		case ">":
			query = query.Where(filter.Field+" > ?", filter.Value)
		case "<":
			query = query.Where(filter.Field+" < ?", filter.Value)
		case ">=":
			query = query.Where(filter.Field+" >= ?", filter.Value)
		case "<=":
			query = query.Where(filter.Field+" <= ?", filter.Value)
		case "LIKE":
			query = query.Where(filter.Field+" LIKE ?", "%"+filter.Value+"%")
		case "IS":
			query = query.Where(filter.Field + " IS " + filter.Value)
		default:
			return nil, errors.New("unsupported filter operator: " + filter.Operator)
		}
	}

	// Apply sorts
	for _, sort := range request.Sorts {
		if sort.SortType != "asc" && sort.SortType != "desc" {
			return nil, errors.New("invalid sort type: " + sort.SortType)
		}
		query = query.Order(sort.Field + " " + sort.SortType)
	}

	// Apply limit and offset
	if request.Limit > 0 {
		query = query.Limit(int(request.Limit))
	}
	query = query.Offset(int(request.Offset))

	// Execute the query
	var records []T
	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Create a new record
func CreateRecord[T any](db *gorm.DB, record *T) error {
	if db == nil {
		return errors.New("database connection is nil")
	}

	if err := db.Create(record).Error; err != nil {
		return err
	}

	return nil
}

// Update an existing record
func UpdateRecord[T any](db *gorm.DB, id any, updates T) error {
	if db == nil {
		return errors.New("database connection is nil")
	}

	if err := db.Model(new(T)).Where("id = ?", id).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

// Retrieve a record by ID
func GetRecordByID[T any](db *gorm.DB, id any, preloads []string) (*T, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	query := db.Model(new(T))

	// Apply preloads
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	var record T
	if err := query.First(&record, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if not found
		}
		return nil, err
	}

	return &record, nil
}

// Delete a record by ID
func DeleteRecordByID[T any](db *gorm.DB, id any) error {
	if db == nil {
		return errors.New("database connection is nil")
	}

	if err := db.Delete(new(T), "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

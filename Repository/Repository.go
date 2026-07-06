package Repository

import (
	"regexp"
	"strings"

	"gorm.io/gorm"
)

var validIdentifier = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

type Condition struct {
	Field    string
	Operator string
	Value    interface{}
}

type QueryParams struct {
	Top          int
	Skip         int
	SortBy       string
	Order        string
	Filter       map[string]string
	Conditions   []Condition
	Search       string
	SearchFields []string
}

type PaginationResult[T any] struct {
	Data       []T   `json:"Data"`
	Total      int64 `json:"Total"`
	Top        int   `json:"Top"`
	Skip       int   `json:"Skip"`
	TotalPages int   `json:"TotalPages"`
}

type Repository[T any] interface {
	FindAll(params QueryParams) (*PaginationResult[T], error)
	FindByID(id uint) (T, error)
	Create(entity *T) error
	Update(id uint, entity *T) error
	Delete(id uint) error
}

type GormRepository[T any] struct {
	DB *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{DB: db}
}

func (r *GormRepository[T]) FindAll(params QueryParams) (*PaginationResult[T], error) {
	var entities []T
	var total int64

	if params.Top < 1 {
		params.Top = 10
	}
	if params.Skip < 0 {
		params.Skip = 0
	}

	query := r.DB.Model(new(T))

	for field, value := range params.Filter {
		if validIdentifier.MatchString(field) {
			query = query.Where("\""+field+"\" = ?", value)
		}
	}

	for _, cond := range params.Conditions {
		if validIdentifier.MatchString(cond.Field) {
			query = query.Where("\""+cond.Field+"\" "+cond.Operator+" ?", cond.Value)
		}
	}

	if params.Search != "" && len(params.SearchFields) > 0 {
		conditions := make([]string, 0, len(params.SearchFields))
		args := make([]interface{}, 0, len(params.SearchFields))
		for _, field := range params.SearchFields {
			if validIdentifier.MatchString(field) {
				conditions = append(conditions, "\""+field+"\" LIKE ?")
				args = append(args, "%"+params.Search+"%")
			}
		}
		if len(conditions) > 0 {
			query = query.Where(strings.Join(conditions, " OR "), args...)
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	if params.SortBy != "" && validIdentifier.MatchString(params.SortBy) {
		order := "asc"
		if params.Order == "desc" {
			order = "desc"
		}
		query = query.Order("\"" + params.SortBy + "\" " + order)
	}

	if err := query.Offset(params.Skip).Limit(params.Top).Find(&entities).Error; err != nil {
		return nil, err
	}

	totalPages := int((total + int64(params.Top) - 1) / int64(params.Top))

	return &PaginationResult[T]{
		Data:       entities,
		Total:      total,
		Top:        params.Top,
		Skip:       params.Skip,
		TotalPages: totalPages,
	}, nil
}

func (r *GormRepository[T]) FindByID(id uint) (T, error) {
	var entity T
	return entity, r.DB.First(&entity, id).Error
}

func (r *GormRepository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

func (r *GormRepository[T]) Update(id uint, entity *T) error {
	return r.DB.Model(entity).Where("\"ID\" = ?", id).Updates(entity).Error
}

func (r *GormRepository[T]) Delete(id uint) error {
	var entity T
	return r.DB.Delete(&entity, id).Error
}

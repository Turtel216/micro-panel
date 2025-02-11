package handler

import "time"

type (
	ProductReq struct {
		ID           int64   `json:"id"`
		Name         string  `json:"name"`
		Image        string  `json:"image"`
		Category     string  `json:"category"`
		Description  string  `json:"description"`
		Rating       int64   `json:"rating"`
		NumReviews   int64   `json:"num_reviews"`
		Price        float32 `json:"price"`
		CountInStock int64   `json:"count_in_stock"`
	}

	ProductRes struct {
		ID           int64      `json:"id"`
		Name         string     `json:"name"`
		Image        string     `json:"image"`
		Category     string     `json:"category"`
		Description  string     `json:"description"`
		Rating       int64      `json:"rating"`
		NumReviews   int64      `json:"num_reviews"`
		Price        float32    `json:"price"`
		CountInStock int64      `json:"count_in_stock"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    *time.Time `json:"updated_at"`
	}

	OrderItem struct {
		Name      string  `json:"name"`
		Quantity  int64   `json:"quantity"`
		Image     string  `json:"image"`
		Price     float32 `json:"price"`
		ProductID int64   `json:"product_id"`
	}

	OrderReq struct {
		ID            int64       `json:"id"`
		Items         []OrderItem `json:"items"`
		PaymentMethod string      `json:"payment_method"`
		TaxPrice      float32     `json:"tax_price"`
		ShippingPrice float32     `json:"shipping_price"`
		TotalPrice    float32     `json:"total_price"`
		Status        string      `json:"status"`
	}

	OrderRes struct {
		ID            int64       `json:"id"`
		Items         []OrderItem `json:"items"`
		PaymentMethod string      `json:"payment_method"`
		TaxPrice      float32     `json:"tax_price"`
		ShippingPrice float32     `json:"shipping_price"`
		TotalPrice    float32     `json:"total_price"`
		Status        string      `json:"status"`
		CreatedAt     time.Time   `json:"created_at"`
		UpdatedAt     *time.Time  `json:"updated_at"`
	}

	UserReq struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}

	UserRes struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}

	ListUserRes struct {
		Users []UserRes `json:"users"`
	}

	LoginUserReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginUserRes struct {
		SessionID             string    `json:"session_id"`
		RefreshToken          string    `json:"refresh_token"`
		AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
		RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
		AccessToken           string    `json:"access_token"`
		User                  UserRes   `json:"user"`
	}

	RenewAccessTokenReq struct {
		RefreshToken string `json:"refresh_token"`
	}

	RenewAccessTokenRes struct {
		AccessToken          string    `json:"access_token"`
		AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	}
)

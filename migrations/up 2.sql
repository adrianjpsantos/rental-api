-- ENUMS
CREATE DATABASE rental_db;   
CREATE TYPE user_role AS ENUM (
    'admin',
    'lessor',
    'lessee'
);

CREATE TYPE rental_status AS ENUM (
    'pending',
    'approved',
    'active',
    'completed',
    'cancelled',
    'rejected'
);

CREATE TYPE payment_status AS ENUM (
    'pending',
    'completed',
    'refunded',
    'failed'
);

CREATE TYPE delivery_method AS ENUM (
    'pickup',
    'delivery',
    'shipping'
);

CREATE TYPE review_type AS ENUM (
    'as_lessor',
    'as_lessee'
);

CREATE TYPE availability_type AS ENUM (
    'blocked',
    'available'
);

CREATE TYPE availability_reason AS ENUM (
    'maintenance',
    'rented',
    'owner_block'
);

-- USERS
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    cpf TEXT UNIQUE,
    phone TEXT,
    birth_date TIMESTAMP,
    avatar_url TEXT,
    is_verified BOOLEAN DEFAULT FALSE,
    role user_role NOT NULL,
    reputation REAL DEFAULT 0,
    total_rentals INT DEFAULT 0,
    total_items_rented INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE sessions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    token_hash TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    actived BOOLEAN DEFAULT TRUE,

    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- CATEGORIES
CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    slug TEXT UNIQUE,
    description TEXT,
    icon TEXT,
    position INT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- ITEMS
CREATE TABLE items (
    id UUID PRIMARY KEY,
    owner_id UUID NOT NULL,
    category_id UUID NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    brand TEXT,
    model TEXT,
    year INT,
    condition TEXT,
    price_per_day NUMERIC(10,2),
    price_per_hour NUMERIC(10,2),
    min_rental_days INT,
    max_rental_days INT,
    quantity INT,
    location TEXT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    photos TEXT[],
    rules TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    FOREIGN KEY (owner_id) REFERENCES users(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- RENTALS
CREATE TABLE rentals (
    id UUID PRIMARY KEY,
    item_id UUID NOT NULL,
    lessee_id UUID NOT NULL,
    lessor_id UUID NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    total_amount NUMERIC(10,2),
    status rental_status NOT NULL,
    payment_status payment_status NOT NULL,
    delivery_method delivery_method NOT NULL,
    notes TEXT,
    cancellation_reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    started_at TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    cancelled_at TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (lessee_id) REFERENCES users(id),
    FOREIGN KEY (lessor_id) REFERENCES users(id)
);

-- AVAILABILITY SLOTS
CREATE TABLE availability_slots (
    id UUID PRIMARY KEY,
    item_id UUID NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    type availability_type NOT NULL,
    reason availability_reason,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    FOREIGN KEY (item_id) REFERENCES items(id)
);

-- REVIEWS
CREATE TABLE reviews (
    id UUID PRIMARY KEY,
    rental_id UUID NOT NULL,
    reviewer_id UUID NOT NULL,
    reviewed_id UUID NOT NULL,
    item_id UUID NOT NULL,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    review_type review_type NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    FOREIGN KEY (rental_id) REFERENCES rentals(id),
    FOREIGN KEY (reviewer_id) REFERENCES users(id),
    FOREIGN KEY (reviewed_id) REFERENCES users(id),
    FOREIGN KEY (item_id) REFERENCES items(id),

    UNIQUE (rental_id, reviewer_id)
);

-- INDEXES
CREATE INDEX idx_items_owner ON items(owner_id);
CREATE INDEX idx_items_category ON items(category_id);

CREATE INDEX idx_rentals_item ON rentals(item_id);
CREATE INDEX idx_rentals_lessee ON rentals(lessee_id);
CREATE INDEX idx_rentals_lessor ON rentals(lessor_id);

CREATE INDEX idx_availability_item_dates 
ON availability_slots(item_id, start_date, end_date);

CREATE INDEX idx_reviews_item ON reviews(item_id);
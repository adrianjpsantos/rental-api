CREATE TYPE system_role AS ENUM (
    'user',
    'admin'
);

CREATE TYPE auth_provider AS ENUM (
    'local',
    'google',
    'facebook',
    'github',
    'apple',
    'microsoft'
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

--------------------------------------------------------
-- USERS
--------------------------------------------------------

CREATE TABLE users (
    id UUID PRIMARY KEY,

    role system_role NOT NULL DEFAULT 'user',

    reputation REAL NOT NULL DEFAULT 0,

    total_rentals INT NOT NULL DEFAULT 0,

    total_items_rented INT NOT NULL DEFAULT 0,

    active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    deleted_at TIMESTAMP WITH TIME ZONE
);

--------------------------------------------------------
-- PROFILES
--------------------------------------------------------

CREATE TABLE profiles (
    user_id UUID PRIMARY KEY,

    first_name TEXT NOT NULL,

    last_name TEXT,

    cpf TEXT UNIQUE,

    phone TEXT,

    birth_date DATE,

    avatar_url TEXT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

--------------------------------------------------------
-- AUTH ACCOUNTS
--------------------------------------------------------

CREATE TABLE auth_accounts (

    id UUID PRIMARY KEY,

    user_id UUID NOT NULL,

    provider auth_provider NOT NULL,

    provider_user_id TEXT NOT NULL,

    email TEXT,

    password_hash TEXT,

    email_verified BOOLEAN NOT NULL DEFAULT FALSE,

    is_primary BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    UNIQUE(provider, provider_user_id)
);

--------------------------------------------------------
-- SESSIONS
--------------------------------------------------------

CREATE TABLE sessions (

    id UUID PRIMARY KEY,

    auth_account_id UUID NOT NULL,

    token_hash TEXT NOT NULL UNIQUE,

    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,

    revoked_at TIMESTAMP WITH TIME ZONE,

    last_used_at TIMESTAMP WITH TIME ZONE,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    FOREIGN KEY (auth_account_id)
        REFERENCES auth_accounts(id)
        ON DELETE CASCADE
);

--------------------------------------------------------
-- CATEGORIES
--------------------------------------------------------

CREATE TABLE categories (

    id UUID PRIMARY KEY,

    name TEXT NOT NULL,

    slug TEXT UNIQUE,

    description TEXT,

    icon TEXT,

    position INT,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

--------------------------------------------------------
-- ITEMS
--------------------------------------------------------

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

    quantity INT NOT NULL DEFAULT 1,

    location TEXT,

    latitude DOUBLE PRECISION,

    longitude DOUBLE PRECISION,

    rules TEXT,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    FOREIGN KEY (owner_id)
        REFERENCES users(id),

    FOREIGN KEY (category_id)
        REFERENCES categories(id)
);

--------------------------------------------------------
-- ITEM IMAGES
--------------------------------------------------------

CREATE TABLE item_images (

    id UUID PRIMARY KEY,

    item_id UUID NOT NULL,

    image_url TEXT NOT NULL,

    position INT NOT NULL DEFAULT 0,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    FOREIGN KEY (item_id)
        REFERENCES items(id)
        ON DELETE CASCADE
);

--------------------------------------------------------
-- RENTALS
--------------------------------------------------------

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

    started_at TIMESTAMP WITH TIME ZONE,

    completed_at TIMESTAMP WITH TIME ZONE,

    cancelled_at TIMESTAMP WITH TIME ZONE,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    FOREIGN KEY (item_id)
        REFERENCES items(id),

    FOREIGN KEY (lessee_id)
        REFERENCES users(id),

    FOREIGN KEY (lessor_id)
        REFERENCES users(id)
);

--------------------------------------------------------
-- AVAILABILITY
--------------------------------------------------------

CREATE TABLE availability_slots (

    id UUID PRIMARY KEY,

    item_id UUID NOT NULL,

    start_date TIMESTAMP WITH TIME ZONE NOT NULL,

    end_date TIMESTAMP WITH TIME ZONE NOT NULL,

    type availability_type NOT NULL,

    reason availability_reason,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    FOREIGN KEY (item_id)
        REFERENCES items(id)
);

--------------------------------------------------------
-- REVIEWS
--------------------------------------------------------

CREATE TABLE reviews (

    id UUID PRIMARY KEY,

    rental_id UUID NOT NULL,

    reviewer_id UUID NOT NULL,

    reviewed_id UUID NOT NULL,

    item_id UUID NOT NULL,

    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),

    comment TEXT,

    review_type review_type NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    FOREIGN KEY (rental_id)
        REFERENCES rentals(id),

    FOREIGN KEY (reviewer_id)
        REFERENCES users(id),

    FOREIGN KEY (reviewed_id)
        REFERENCES users(id),

    FOREIGN KEY (item_id)
        REFERENCES items(id),

    UNIQUE (rental_id, reviewer_id)
);

--------------------------------------------------------
-- INDEXES
--------------------------------------------------------

CREATE INDEX idx_auth_accounts_user
ON auth_accounts(user_id);

CREATE INDEX idx_auth_accounts_email
ON auth_accounts(email);

CREATE INDEX idx_sessions_auth_account
ON sessions(auth_account_id);

CREATE INDEX idx_items_owner
ON items(owner_id);

CREATE INDEX idx_items_category
ON items(category_id);

CREATE INDEX idx_item_images_item
ON item_images(item_id);

CREATE INDEX idx_rentals_item
ON rentals(item_id);

CREATE INDEX idx_rentals_lessee
ON rentals(lessee_id);

CREATE INDEX idx_rentals_lessor
ON rentals(lessor_id);

CREATE INDEX idx_availability_item_dates
ON availability_slots(item_id, start_date, end_date);

CREATE INDEX idx_reviews_item
ON reviews(item_id);
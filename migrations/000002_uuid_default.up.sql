--------------------------------------------------------
-- ENABLE UUID GENERATION
--------------------------------------------------------

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--------------------------------------------------------
-- USERS
--------------------------------------------------------

ALTER TABLE users
ALTER COLUMN id SET DEFAULT gen_random_uuid();

--------------------------------------------------------
-- AUTH ACCOUNTS
--------------------------------------------------------

ALTER TABLE auth_accounts
ALTER COLUMN id SET DEFAULT gen_random_uuid();

--------------------------------------------------------
-- SESSIONS
--------------------------------------------------------

ALTER TABLE sessions
ALTER COLUMN id SET DEFAULT gen_random_uuid();

--------------------------------------------------------
-- CATEGORIES
--------------------------------------------------------

ALTER TABLE categories
ALTER COLUMN id SET DEFAULT gen_random_uuid();

--------------------------------------------------------
-- ITEMS
--------------------------------------------------------

ALTER TABLE items
ALTER COLUMN id SET DEFAULT gen_random_uuid();

--------------------------------------------------------
-- ITEM IMAGES
--------------------------------------------------------

ALTER TABLE item_images
ALTER COLUMN id SET DEFAULT gen_random_uuid();

--------------------------------------------------------
-- RENTALS
--------------------------------------------------------

ALTER TABLE rentals
ALTER COLUMN id SET DEFAULT gen_random_uuid();

--------------------------------------------------------
-- AVAILABILITY SLOTS
--------------------------------------------------------

ALTER TABLE availability_slots
ALTER COLUMN id SET DEFAULT gen_random_uuid();

--------------------------------------------------------
-- REVIEWS
--------------------------------------------------------

ALTER TABLE reviews
ALTER COLUMN id SET DEFAULT gen_random_uuid();
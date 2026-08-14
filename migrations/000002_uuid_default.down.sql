--------------------------------------------------------
-- USERS
--------------------------------------------------------

ALTER TABLE users
ALTER COLUMN id DROP DEFAULT;

--------------------------------------------------------
-- AUTH ACCOUNTS
--------------------------------------------------------

ALTER TABLE auth_accounts
ALTER COLUMN id DROP DEFAULT;

--------------------------------------------------------
-- SESSIONS
--------------------------------------------------------

ALTER TABLE sessions
ALTER COLUMN id DROP DEFAULT;

--------------------------------------------------------
-- CATEGORIES
--------------------------------------------------------

ALTER TABLE categories
ALTER COLUMN id DROP DEFAULT;

--------------------------------------------------------
-- ITEMS
--------------------------------------------------------

ALTER TABLE items
ALTER COLUMN id DROP DEFAULT;

--------------------------------------------------------
-- ITEM IMAGES
--------------------------------------------------------

ALTER TABLE item_images
ALTER COLUMN id DROP DEFAULT;

--------------------------------------------------------
-- RENTALS
--------------------------------------------------------

ALTER TABLE rentals
ALTER COLUMN id DROP DEFAULT;

--------------------------------------------------------
-- AVAILABILITY SLOTS
--------------------------------------------------------

ALTER TABLE availability_slots
ALTER COLUMN id DROP DEFAULT;

--------------------------------------------------------
-- REVIEWS
--------------------------------------------------------

ALTER TABLE reviews
ALTER COLUMN id DROP DEFAULT;
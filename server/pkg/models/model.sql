CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  thumbnail VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  type VARCHAR(255) NOT NULL,
  value BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS campaigns (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  value BIGINT NOT NULL,
  product_id INTEGER NOT NULL REFERENCES products(id),
  required_contributors INTEGER NOT NULL,
  total_contributors INTEGER NOT NULL,
  started_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  closed_at TIMESTAMP NOT NULL DEFAULT NOW() + INTERVAL '14 days'
);

CREATE TABLE IF NOT EXISTS contributors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    campaign_id INTEGER NOT NULL REFERENCES campaigns(id),
    confirmed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS results (
  id SERIAL PRIMARY KEY,
  campaign_id INTEGER NOT NULL REFERENCES campaigns(id),
  receiver_id INTEGER NOT NULL REFERENCES contributors(id),
  amount BIGINT NOT NULL,
  closed_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- CREATE OR REPLACE FUNCTION create_product(
--     thumbnail VARCHAR(255),
--     name VARCHAR(255),
--     description TEXT,
--     type VARCHAR(255),
--     value NUMERIC(10,2)
-- ) RETURNS products AS $$
-- DECLARE
--     inserted_product products;
-- BEGIN
--     INSERT INTO products (thumbnail, name, description, type, value)
--     VALUES (thumbnail, name, description, type, value)
--     RETURNING * INTO inserted_product;
--
--     RETURN inserted_product;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION update_product(
--   product_id INTEGER,
--   thumbnail VARCHAR(255),
--   name VARCHAR(255),
--   description TEXT,
--   type VARCHAR(255),
--   value NUMERIC(10,2)
-- ) RETURNS products AS $$
-- DECLARE
--   updated_product products;
-- BEGIN
--     UPDATE products SET thumbnail = thumbnail, name = name, description = description, type = type, value = value, updated_at = NOW() WHERE products.id = product_id RETURNING * INTO updated_product;
--
--     RETURN updated_product;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION delete_product(
--   product_id INTEGER
-- ) RETURNS products AS $$
-- DECLARE deleted_product products;
-- BEGIN
--     DELETE FROM products WHERE products.id == product_id RETURNING * INTO deleted_product;
--
--     RETURN deleted_product;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION get_product(
--   product_id INTEGER
-- ) RETURNS products AS $$
-- DECLARE
--     product products;
-- BEGIN
--    SELECT * FROM products WHERE products.id = product_id INTO product;
--
--     RETURN product;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION get_products() RETURNS SETOF products AS $$
-- DECLARE
--     products products;
-- BEGIN
--    SELECT * FROM products INTO products;
--
--     RETURN NEXT products;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION create_campaign(
--     title VARCHAR(255),
--     value NUMERIC(10,2),
--     product_id INTEGER,
--     required_contributors INTEGER,
--     total_contributors INTEGER
-- ) RETURNS campaigns AS $$
-- DECLARE
--     inserted_campaign campaigns;
-- BEGIN
--     INSERT INTO campaigns (title, value, product_id, required_contributors, total_contributors) VALUES (title, value, product_id, required_contributors, total_contributors) RETURNING * INTO inserted_campaign;
--
--     RETURN inserted_campaign;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION update_campaign(
--     campaign_id INTEGER,
--     title VARCHAR(255),
--     value NUMERIC(10,2),
--     product_id INTEGER,
--     required_contributors INTEGER,
--     total_contributors INTEGER
-- ) RETURNS campaigns AS $$
-- DECLARE
--     updated_campaign campaigns;
-- BEGIN
--     UPDATE campaigns SET title = title, value = value, product_id = product_id, required_contributors = required_contributors, total_contributors = total_contributors, updated_at = NOW() WHERE campaigns.id = campaign_id RETURNING * INTO updated_campaign;
--
--     RETURN updated_campaign;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION delete_campaign(
--     campaign_id INTEGER
-- ) RETURNS campaigns AS $$
-- DECLARE
--     deleted_campaign campaigns;
-- BEGIN
--     DELETE FROM campaigns WHERE campaigns.id = campaign_id RETURNING * INTO deleted_campaign;
--
--     RETURN deleted_campaign;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION get_campaign(
--     campaign_id INTEGER
-- ) RETURNS campaigns AS $$
-- DECLARE
--     campaign campaigns;
-- BEGIN
--     SELECT * FROM campaigns WHERE campaigns.id = campaign_id INTO campaign;
--
--     RETURN campaign;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION get_campaigns() RETURNS SETOF campaigns AS $$
-- BEGIN
--     SELECT * FROM campaigns;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION create_contributor(
--   name VARCHAR(255),
--   email VARCHAR(255),
--   phone VARCHAR(255),
--   campaign_id INTEGER
-- ) RETURNS contributors AS $$
-- DECLARE inserted_contributor contributors;
-- BEGIN
--     INSERT INTO contributors (name, email, phone, campaign_id) VALUES (name, email, phone, campaign_id) RETURNING * INTO inserted_contributor;
--
--     RETURN inserted_contributor;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION confirm_contributor(
-- contributor_id INTEGER,
-- foreign_campaign_id INTEGER
-- ) RETURNS contributors AS $$
-- DECLARE
--     confirmed_contributor contributors;
-- BEGIN
--     UPDATE contributors SET confirmed = true, updated_at = NOW() WHERE contributors.id = contributor_id AND contributors.campaign_id = foreign_campaign_id RETURNING * INTO confirmed_contributor;
--
--     RETURN confirmed_contributor;
-- END
-- $$ LANGUAGE plpgsql;
--
-- --increase_total_contributors
-- CREATE OR REPLACE FUNCTION increase_total_contributors(
--   foreign_campaign_id INTEGER
-- ) RETURNS campaigns AS $$
-- DECLARE
--     updated_campaign campaigns;
-- BEGIN
--     UPDATE campaigns SET total_contributors = total_contributors + 1, updated_at = NOW() WHERE campaigns.id = foreign_campaign_id RETURNING * INTO updated_campaign;
--
--     RETURN updated_campaign;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION get_contributor(
--   contributor_id INTEGER,
--   foreign_campaign_id INTEGER
-- ) RETURNS contributors AS $$
-- DECLARE
--     contributor contributors;
-- BEGIN
--     SELECT * FROM contributors WHERE contributors.id = contributor_id AND contributors.campaign_id = foreign_campaign_id INTO contributor;
--
--     RETURN contributor;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION get_contributors(
--   foreign_campaign_id INTEGER
-- ) RETURNS SETOF contributors AS $$
-- BEGIN
--     SELECT * FROM contributors WHERE contributors.campaign_id = foreign_campaign_id;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION create_result(
--   campaign_id INTEGER,
--   amount NUMERIC(10,2),
--   receiver_id INTEGER
-- ) RETURNS results AS $$
-- DECLARE
--     inserted_result results;
-- BEGIN
--     INSERT INTO results (campaign_id, amount, receiver_id) VALUES (campaign_id, amount, receiver_id) RETURNING * INTO inserted_result;
--
--     RETURN inserted_result;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION get_result(
--   result_id INTEGER
-- ) RETURNS results AS $$
-- DECLARE
--     result results;
-- BEGIN
--    SELECT * FROM results WHERE results.id = result_id INTO result;
--
--     RETURN result;
-- END
-- $$ LANGUAGE plpgsql;
--
-- CREATE OR REPLACE FUNCTION get_results() RETURNS SETOF results AS $$
-- BEGIN
--    SELECT * FROM results;
-- END
-- $$ LANGUAGE plpgsql;

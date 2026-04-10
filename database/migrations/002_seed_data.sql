-- Insert a test landlord
INSERT INTO landlords (name, overall_rating) 
VALUES ('Aditi Sharma', 4.5) 
RETURNING id;

-- Insert a property (Using coordinates for Koramangala 4th Block)
INSERT INTO properties (landlord_id, title, price_monthly, area_sqft, coordinates, _type, house_type)
VALUES (
    (SELECT id FROM landlords LIMIT 1), 
    'Luxury 2BHK near Sony World Signal', 
    45000, 
    1200, 
    ST_GeographyFromText('POINT(77.6245 12.9339)'),
    '2 BHK',
    'gated'
) RETURNING id;

-- Insert the "Worth It" metrics for this property
INSERT INTO valuation_metrics (property_id, worth_score, market_avg_price, amenity_score, commute_index)
VALUES (
    (SELECT id FROM properties LIMIT 1), 
    82.50, -- This is a good deal!
    48000.00, 
    8, 
    9
);
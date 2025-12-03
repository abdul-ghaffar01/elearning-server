INSERT INTO auth (
	userId, 
    refreshToken, 
    ip, 
    deviceType, 
    os, 
    browser, 
    country, 
    city
)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
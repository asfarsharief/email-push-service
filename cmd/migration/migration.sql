-- SQLite migration script to create the required tables

-- Users Table
CREATE TABLE IF NOT EXISTS Users (
    userId TEXT PRIMARY KEY,
    tenantId TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

-- Credentials Table
CREATE TABLE IF NOT EXISTS Credentials (
    userId TEXT PRIMARY KEY,
    provider TEXT NOT NULL,
    accessToken TEXT NOT NULL,
    refreshToken TEXT NOT NULL,
    expiresAt TEXT NOT NULL,
    FOREIGN KEY (userId) REFERENCES Users(userId)
);

-- Quota Tracking Table
CREATE TABLE IF NOT EXISTS QuotaTracking (
    tenantId TEXT NOT NULL,
    date TEXT NOT NULL,
    emailsSent INTEGER DEFAULT 0,
    dailyLimit INTEGER NOT NULL,
    quotaMultiplier INTEGER NOT NULL, 
    PRIMARY KEY (tenantId, date),
    FOREIGN KEY (tenantId) REFERENCES Users(tenantId)
);
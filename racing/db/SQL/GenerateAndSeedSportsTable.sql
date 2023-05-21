CREATE TABLE IF NOT EXISTS matches (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    stadium TEXT,
    sport TEXT,
    team1 TEXT,
    team2 TEXT,
 time TIMESTAMP
);

INSERT INTO matches (name, stadium, sport, team1, team2, time)
VALUES
    ('Team A vs Team B at Stadium 1', 'Stadium 1', 'AFL', 'Team A', 'Team B', DATETIME('now', '+1 day')),
    ('Team C vs Team D at Stadium 2', 'Stadium 2', 'Cricket', 'Team C', 'Team D', DATETIME('now', '+2 day')),
    ('Team E vs Team F at Stadium 3', 'Stadium 3', 'Rugby', 'Team E', 'Team F', DATETIME('now', '+3 day')),
    ('Team G vs Team H at Stadium 4', 'Stadium 4', 'Cricket', 'Team G', 'Team H', DATETIME('now', '+4 day'));

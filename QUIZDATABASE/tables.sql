DROP TABLE IF EXISTS 4aq;
CREATE TABLE 4aq (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  question   TEXT NOT NULL,
  answer     TEXT NOT NULL,
  explination TEXT NOT NULL
);

INSERT INTO 4aq
  (question, answer, explination)
VALUES 
    ('A plane travelling 534 mi/hr encounters a headwind of 16 mi/hr. Determine the Resultant velocity(answer must be an integer ex: 500).', '966457169', 'Headwind slows down the plane which means you would subtract the magnitude of each velocity to get the resultant'),
    ('A ball is thrown horizontally with an initial velocity of 20.0 meters per second from the top of a tower 60.0 meters high.\nWhat is the horizontal velocity of the ball just before it reaches the ground(answer must be integer ex: 1,2,3,4)?', '2381486463', 'Since we ignore air resistance in Regents-Physics there are no forces acting on the horizontal velocity meaning the velocity right before the ball hits the ground will be equal to the initial horizontal velocity.'),
    ('A motorboat traveling 5 m/s East encounters a current travelling 2.5 m/s North\nWhat is the resultant velocity of the motorboat? (round to nearest hundredths ex: 2.38)', '4121720140', 'Use the Pythagorean Theorem(a^2+b^2=c^2) to find the resultant velocity');
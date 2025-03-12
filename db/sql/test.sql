-- Create tables
CREATE TABLE `profile` (
    `username` TEXT NOT NULL PRIMARY KEY,
    `created_at` TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `notes`  (
    `id` TEXT NOT NULL PRIMARY KEY,
    `username` TEXT NOT NULL,
    `content` TEXT NOT NULL,
    `created_at` TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(`username`) REFERENCES `profile`(`username`) ON DELETE CASCADE
);



-- Insert a user into the profile table
INSERT INTO `profile` (`username`) VALUES ('test-user1');
INSERT INTO `profile` (`username`) VALUES ('test-user2');
INSERT INTO `profile` (`username`) VALUES ('test-user3');

-- Insert some notes associated with the user
INSERT INTO `notes` (`id`, `username`, `content`) VALUES 
('1', 'test-user1', 'first note of test-user1'),
('2', 'test-user3', 'test-user3 note.'),
('3', 'test-user2', 'test-user2 note again'),
('4', 'test-user1', 'test-user1 note'),
('5', 'test-user2', 'test-user2 note one more.'),
('6', 'test-user3', 'hello world');

SELECT * FROM profile;
SELECT * FROM notes;


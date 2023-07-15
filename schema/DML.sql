INSERT INTO `nationality` (`nationality_id`, `nationality_name`, `nationality_code`) VALUES (NULL, 'Indonesia', 'ID');

INSERT INTO `customer` (`cst_id`, `nationality_id`, `cst_name`, `cst_dob`, `cst_phoneNum`, `cst_email`) VALUES 
(NULL, '1', 'muhammad nur basari', '2023-07-01', '081383058536', 'm.nurbasari@gmail.com');

INSERT INTO `family_list` (`fl_id`, `cst_id`, `fl_relation`, `fl_name`, `fl_dob`) VALUES 
(NULL, '1', 'istri', 'itri_test_1', '2023-07-11'), (NULL, '1', 'anak_ke_1', 'anak_test_1', '2023-07-15');
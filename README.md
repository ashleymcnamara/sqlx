## Getting Started

* Clone this repository to $GOPATH/src/bketelsen/sqlx -- OR 
* go get github.com/bketelsen/sqlx
* Ensure docker is running

## Starting the Database

* cd $GOPATH/src/github.com/bketelsen/sqlx
* ./run.sh

## Stopping the Database

* ./stopDatabase.sh

## Getting a mysql shell in a running container

* ./shell.sh

## mysql u/p
docker/docker

## Schema 

mysql> show tables;
+----------------------+
| Tables_in_employees |
+----------------------+
| current_dept_emp |
| departments |
| dept_emp |
| dept_emp_latest_date |
| dept_manager |
| employees |
| salaries |
| titles |
| v_full_departments |
| v_full_employees |
+----------------------+
10 rows in set (0.00 sec)


mysql> describe current_dept_emp;
+-----------+---------+------+-----+---------+-------+
| Field | Type | Null | Key | Default | Extra |
+-----------+---------+------+-----+---------+-------+
| emp_no | int(11) | NO | | NULL | |
| dept_no | char(4) | NO | | NULL | |
| from_date | date | YES | | NULL | |
| to_date | date | YES | | NULL | |
+-----------+---------+------+-----+---------+-------+
4 rows in set (0.00 sec)



mysql> describe departments;
+-----------+-------------+------+-----+---------+-------+
| Field | Type | Null | Key | Default | Extra |
+-----------+-------------+------+-----+---------+-------+
| dept_no | char(4) | NO | PRI | NULL | |
| dept_name | varchar(40) | NO | UNI | NULL | |
+-----------+-------------+------+-----+---------+-------+
2 rows in set (0.00 sec)


mysql> describe dept_emp;
+-----------+---------+------+-----+---------+-------+
| Field | Type | Null | Key | Default | Extra |
+-----------+---------+------+-----+---------+-------+
| emp_no | int(11) | NO | PRI | NULL | |
| dept_no | char(4) | NO | PRI | NULL | |
| from_date | date | NO | | NULL | |
| to_date | date | NO | | NULL | |
+-----------+---------+------+-----+---------+-------+
4 rows in set (0.00 sec)



mysql> describe dept_emp_latest_date;
+-----------+---------+------+-----+---------+-------+
| Field | Type | Null | Key | Default | Extra |
+-----------+---------+------+-----+---------+-------+
| emp_no | int(11) | NO | | NULL | |
| from_date | date | YES | | NULL | |
| to_date | date | YES | | NULL | |
+-----------+---------+------+-----+---------+-------+
3 rows in set (0.01 sec)


mysql> describe dept_manager;
+-----------+---------+------+-----+---------+-------+
| Field | Type | Null | Key | Default | Extra |
+-----------+---------+------+-----+---------+-------+
| emp_no | int(11) | NO | PRI | NULL | |
| dept_no | char(4) | NO | PRI | NULL | |
| from_date | date | NO | | NULL | |
| to_date | date | NO | | NULL | |
+-----------+---------+------+-----+---------+-------+
4 rows in set (0.00 sec)



mysql> describe employees;
+------------+---------------+------+-----+---------+-------+
| Field | Type | Null | Key | Default | Extra |
+------------+---------------+------+-----+---------+-------+
| emp_no | int(11) | NO | PRI | NULL | |
| birth_date | date | NO | | NULL | |
| first_name | varchar(14) | NO | | NULL | |
| last_name | varchar(16) | NO | | NULL | |
| gender | enum('M','F') | NO | | NULL | |
| hire_date | date | NO | | NULL | |
+------------+---------------+------+-----+---------+-------+
6 rows in set (0.00 sec)

mysql> describe salaries;
+-----------+---------+------+-----+---------+-------+
| Field | Type | Null | Key | Default | Extra |
+-----------+---------+------+-----+---------+-------+
| emp_no | int(11) | NO | PRI | NULL | |
| salary | int(11) | NO | | NULL | |
| from_date | date | NO | PRI | NULL | |
| to_date | date | NO | | NULL | |
+-----------+---------+------+-----+---------+-------+
4 rows in set (0.00 sec)



mysql> describe titles;
+-----------+-------------+------+-----+---------+-------+
| Field | Type | Null | Key | Default | Extra |
+-----------+-------------+------+-----+---------+-------+
| emp_no | int(11) | NO | PRI | NULL | |
| title | varchar(50) | NO | PRI | NULL | |
| from_date | date | NO | PRI | NULL | |
| to_date | date | YES | | NULL | |
+-----------+-------------+------+-----+---------+-------+
4 rows in set (0.00 sec)



create database if not exists cnf;
use cnf;

create table if not exists claim (
  id int not null AUTO_INCREMENT,
  cnf_version varchar(50) not null,
  executed_by  varchar(50) not null,
  upload_time datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  partner_name varchar(50),
  s3_file_key varchar(255),
  primary key (id)
);

create table if not exists claim_result (
  id int not null AUTO_INCREMENT,
  claim_id int not null,
  suite_name varchar(255),
  test_id varchar(255),
  test_status varchar(10) not null,
  primary key (id),
  foreign key (claim_id) references claim(id)
);

set @x := (
  select count(*) from information_schema.statistics
  where table_name = 'claim'
  and index_name = 'claim_upload_datetime'
  and table_schema = database());
set @sql := if( 
  @x > 0, 
  'select ''Index exists.''', 
  'create index claim_upload_datetime on claim (upload_time);');
PREPARE stmt FROM @sql;
EXECUTE stmt;

create table if not exists authenticator (
  partner_name varchar(255) not null,
  encoded_password varchar(255) not null
);

-- CollectorAdminUser's and CollectorAdminPassword's values are in tnf-secrets/collector-secrets.json
insert into authenticator (partner_name, encoded_password)
select 'CollectorAdminUser', 'CollectorAdminPassword'
from dual
where not exists (select 1 from authenticator limit 1);
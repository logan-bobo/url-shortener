CREATE ROLE url_shortener 
LOGIN 
PASSWORD 'password';

CREATE DATABASE url_shortener
WITH
   OWNER =  url_shortener
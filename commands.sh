curl  -x "http://localhost:8818"  "http://www.google.com" 

-v debug output of http
-k ignore certification
-s silent do not show errorrs
-S show error message
-L if 3XX error return make new call on new address
-i Include the HTTP response headers in the output
-I fetch the header only
-f fail silently (no output at all) on server errors

curl -vv -ksL "https://example.com" -x "http://<proxy>:<port>"
curl -I "https://www.google.com" -x 1.1.1.1:8080

sudo curl -x http://10.1.1.50:8080/ -fsSL https://download.docker.com/linux/ubuntu/gpg

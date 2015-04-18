# golang
Go programming language experiments.

Project to read CVS, PRN files and output them as HTML.

File content is served by golang http server.
Several http servers are created to handle requests in parallel on different ports, application url, content folders.

Once http servers launched by 'runServer.go', CSV, PRN files are rendered using the following URLs
from 'csvprndata1', 'csvprndata2', 'csvprndata3' folders accordingly :

http://localhost:8080/app1<br>
http://localhost:8085/app2<br>
http://localhost:8090/app3<br>

const http = require('http');

const server = http.CreateServer((req, res) => {
    if (req.url === '/'){
        res.write('Welcome');
        res.end();
    }

    if (req.url === '/page'){
        res.write(JSON.stringify([32,5423,124]));
        res.end();
    }
});

server.listen(8000);


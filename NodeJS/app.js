const http = require('http');
const fs = require('fs').promises;
const server = http.createServer((req, res) => {
    let context = '';
    console.log(req.url)
    if (req.url === '/'){
        context = readFile('index.js');
        res.writeHead(200, { "Content-Type": "text/plain" });
        res.write(context);
        res.end();
    }
    else{    
    context =  readFile(req.url);
    console.log("this is")
        res.writeHead(200, { "Content-Type": "text/plain" });
        res.write(context);
        res.end();
    }
});

async function readFile(filePath) {
  try {
    const data = await fs.readFile(filePath);
   return data.toString();
  } catch (error) {
    return `Got an error trying to read the file: ${error.message}`;
  }
}

server.listen(8000);
const http = require('http')
const fs = require('fs')
const path = require('path')


const server = http.createServer((req, res) =>{
    if(req.url === '/'){
        fs.readFile(path.join(__dirname, 'index.js'), 'utf-8', (err, data)=>{
            if(err) throw err;
            res.end(data)
        })
    }
    // else{
    //     fs.readFile(path.join(__dirname, req.url), 'utf-8', (err, data)=>{
    //         if(err) throw err;
    //         res.end(data)
    //     })
    // }
})
const PORT = process.env.PORT || 3000

server.listen(PORT, ()=> console.log('Server running on port: ${PORT}'))
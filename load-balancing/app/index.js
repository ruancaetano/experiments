const express = require('express');
const app = express();
const port = process.env.PORT

let requestCounter = 0

app.get('/', async (req, res) => {
    console.log("Requests received: " + requestCounter++)
    await new Promise(resolve => setTimeout(resolve, 100)); 
    res.status(200).json({})
});

app.listen(port, () => {
    console.log(`Server is listening at http://localhost:${port}`);
});
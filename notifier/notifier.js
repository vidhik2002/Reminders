const express = require('express');
// const { notify } = require('node-notifier');
const bodyParser = require("body-parser")
const app = express()
const port = process.env.PORT || 9000;

app.use(bodyParser.json())

app.get('/health', (req,res) => {
    res.status(200).send()
})

app.post('/notify', (req,res) => {
    notifying(req.body, reply => res.send(reply))
})
app.get("/", (req, res) => {
    res.send("homepage");
});

app.listen(port, () => console.log(`Listening to port ${port}`))

const notifying = ({title, msg}, cb) => {
    console.log(title, msg)
    cb("some string")
} 
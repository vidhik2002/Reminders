const express = require('express');
const notifier = require('node-notifier');
const bodyParser = require("body-parser")
const app = express()
const port = process.env.PORT || 9000;

app.use(bodyParser.json())
app.setAppUserModelId("com.notifier.id");

app.get('/health', (req,res) => {
    res.status(200).send()
})

app.post('/notify', (req,res) => {
    notify(req.body, reply => res.send(reply))
})
app.get("/", (req, res) => {
    res.send("homepage");
});

app.listen(port, () => console.log(`Listening to port ${port}`))

const notify = ({title, msg}, cb) => {
    
    notifier.notify(
      {
        //appName: "com.notifier.id",
        title: title || "unknown",
        msg: msg || "message",
        sound: true,
        wait: true,
        reply: true,
        closeLabel: "Done?",
        timeout: 15,
      },
      (err, ressponse, reply) => {
        cb(reply);
      }
    );
} 
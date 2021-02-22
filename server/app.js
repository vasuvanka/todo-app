const express = require('express');
const {
    join
} = require('path')

const app = express()
app.use(express.static(join(__dirname, '..', 'dist', 'todo')))

app.listen(process.env.NODE_PORT || 3000, () => {
    console.log('listening on port : ' + (process.env.NODE_PORT || 3000));
})

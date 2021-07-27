const express = require("express");
// const cors = require("cors");

const app = express();
const port = 3000;

// app.use(
//   cors({
//     origin: "http://localhost:8000",
//   })
// );

app.get("/hello", (_, res) => {
  res.json({ name: "alice" });
});

app.listen(port, "0.0.0.0");

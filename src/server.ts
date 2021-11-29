import express from "express";
import cors from "cors";
import env from "dotenv";
import path from "path";
import fs from "fs";

const app = express();

//configs
env.config();

// utilities
app.use(cors());

app.get("*", (req:any, res:any) => {
	const folderPath = path.join(__dirname, "../images/");
	if (fs.existsSync(folderPath + req.url)) {
		res.status(200).sendFile(req.url, {root: folderPath});
	} else {
		res.status(404).json("404 File not found");
	}
});

// app.get("/files/*", getFile);

export default app; //Export server for use in index.js

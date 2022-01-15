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
	const filename = req.query.filename;
	if (filename) {
		if (fs.existsSync(folderPath + filename)) {
			res.status(200).sendFile(filename, {root: folderPath});
		} else {
			res.status(404).json("404 File not found with filename: " + filename);
		}
	} else {
		res.status(400).json("You must provide a filename query");
	}
});

// app.get("/files/*", getFile);

export default app; //Export server for use in index.js

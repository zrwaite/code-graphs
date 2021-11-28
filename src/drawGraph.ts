import path from "path";
import canvasClass from "canvas";
import {createCanvas, loadImage} from "canvas";
import fs from "fs";

const createSlice = (ctx: any, start: number, end: number) => {
	end = start+end;
	start = (start/100)*2*Math.PI-Math.PI/2;
	end = (end/100)*2*Math.PI-Math.PI/2;
	let median = (start+end)/2
	let offset = 10;
	let ox = Math.cos(median)*offset;
	let oy = Math.sin(median)*offset;

	ctx.fillStyle = "green";
	ctx.beginPath();
	ctx.moveTo(500 + ox, 500 + oy);
	ctx.arc(500+ox, 500+oy, 300, start, end);
	ctx.lineTo(500+ox, 500+oy);
	ctx.lineWidth = 5;
	ctx.stroke();
	ctx.fill();
}

const testImage = () => {
	const jsonFolderPath = path.join(__dirname, "../json/");
	let fileName = "languages.json";
	let filePath = jsonFolderPath+fileName;
	let languages = JSON.parse(fs.readFileSync(filePath).toString()).languages;
	let width = 1000;
	let height = 1000;
	let canvas = createCanvas(width, height);
	let ctx = canvas.getContext('2d');
	ctx.fillStyle = "white";
	ctx.fillRect(0, 0, width, height); 
	let totalAngle = 0;

	createSlice(ctx, 0, 25);
	createSlice(ctx, 25, 30);


	
	
	// createSlice(ctx, 0.5);
	// let text = "hello world";
	// ctx.textBaseline = "top";
	// ctx.font = 'bold 70pt Menlo';
	// ctx.textAlign = 'center';
	// ctx.fillText(text, 600, 170);
	// ctx.fillText("This is zoc", 500, 530);
	const imageFolderPath = path.join(__dirname, "../images/");
	let buffer = canvas.toBuffer("image/png");
	let imageFileName = "test.png";
	let imageFilePath = imageFolderPath+imageFileName;
	fs.writeFileSync(imageFilePath, buffer);
}

export default testImage;
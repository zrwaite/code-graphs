import path from "path";
import canvasClass from "canvas";
import {createCanvas, loadImage} from "canvas";
import fs from "fs";
import parseLanguages from "./parseLanguages";

const createSlice = (width: number, height: number, ctx: any, start: number, end: number, colour: string) => {
	end = start+end;
	start = (start/100)*2*Math.PI-Math.PI/2;
	end = (end/100)*2*Math.PI-Math.PI/2;
	let median = (start+end)/2
	let middle = {
		x: width/2, y: height/2
	}
	let radius = Math.min(width, height)*0.48;
	let offset = 0; //radius*0.04;
	let ox = Math.cos(median)*offset;
	let oy = Math.sin(median)*offset;
	ctx.fillStyle = colour;
	ctx.beginPath();
	ctx.moveTo(middle.x + ox, middle.y + oy);
	ctx.arc(middle.x+ox, middle.y+oy, radius, start+offset*0.0001, end-offset*0.0001);
	ctx.lineTo(middle.x+ox, middle.y+oy);
	ctx.lineWidth = 0; //offset*0.3;
	ctx.stroke();
	ctx.fill();
}

const getLanguages = async () => {
	await parseLanguages();
	const folderPath = path.join(__dirname, "../json/");
	let fileName = "languages.json";
	let filePath = folderPath+fileName;
	let languageData = JSON.parse(fs.readFileSync(filePath).toString());
	return languageData.languages;
}

const createPiGraph = async (ctx: any, languages: any[], width:number, height:number) => {
	let totalAngle = 0;
	languages.forEach((language) => {
		if (language.percent >= 0.2) {
			createSlice(width, height, ctx, totalAngle, language.percent, language.colour);
			totalAngle += language.percent;
		}
	})
}

const createLegend = async (ctx: any, languages: any[], start: number, width:number, height:number) => {
	ctx.fillStyle = "white";
	ctx.fillRect(start, 0, width, height); 
	ctx.fillStyle = "lightgrey";
	ctx.fillRect(start+ width/20, height/40, width-width/10, height-height/8);
	ctx.textBaseline = "top";
	ctx.textAlign = 'center';

	ctx.font = 'bold 30pt Menlo';
	ctx.fillStyle = "black";
	ctx.fillText("Top Languages", start+width/2, height/30);
	ctx.fillText("Since October", start+width/2, 2*height/20);

	ctx.font = 'bold 21pt Menlo';
	ctx.fillStyle = "black";
	ctx.textAlign = 'left';
	for (let i=0; i<languages.length; i++) {
		let language = languages[i];
		ctx.fillStyle = language.colour;
		ctx.fillRect(start+width/10, (i+4)*height/20, height/50, height/50);
		ctx.fillText(language.name, start+2*width/10, (i+3.7)*height/20);
		ctx.fillText(language.percent.toString() + "%", start+7*width/10, (i+3.7)*height/20);
	}
	ctx.font = 'bold 14pt Menlo';
	ctx.fillStyle = "black";
	ctx.textAlign = 'center';
	let date = new Date().toLocaleDateString().toString();
	ctx.fillText("Updated on: "+date, start+width/2, 18.4*height/20);
	ctx.fillText("Charts coded by Zac in nodejs", start+width/2, 19*height/20);
}

const createImage = async () => {
	let width = 1000;
	let height = 650;
	let canvas = createCanvas(width, height);
	let ctx = canvas.getContext('2d');
	let languages = await getLanguages();
	createPiGraph(ctx, languages, width*0.6, height)
	createLegend(ctx, languages, width*0.6, width*0.4, height)
	
	const imageFolderPath = path.join(__dirname, "../images/");
	let buffer = canvas.toBuffer("image/png");
	let imageFileName = "pigraph.png";
	let imageFilePath = imageFolderPath+imageFileName;
	fs.writeFileSync(imageFilePath, buffer);
}

export default createImage;